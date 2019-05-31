package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
)

// receiveAgentMesssage : Check and store an agent's message
func receiveAgentMesssage(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
	msg := models.AgentMessage{}
	body := services.ReadBytesFromStream(context.Request.Body)
	json.Unmarshal(body, &msg)

	err := msg.Config.VerifySignature(issuer.Certificate, msg.Signature)
	if err != nil {
		return http.StatusBadRequest, err
	}
	issuerID := fmt.Sprintf("%s:%s", issuer.Organization, issuer.CommonName)
	if msg.SenderID != issuerID {
		return http.StatusForbidden, fmt.Errorf("sender_id / certificate DN mismatch : [%s] vs [%s]", msg.SenderID, issuerID)
	}

	msg.SenderID = fmt.Sprintf("%s:%s", issuer.Organization, issuer.CommonName)
	services.DbZAdd(msg.GetFullKey(), msg.GetSortedSetScore(), msg)

	agent := &models.Agent{
		Organization: issuer.Organization,
		CommonName:   issuer.CommonName,
		LastMessage:  msg,
	}
	services.DbSet(agent.GetFullKeyFor("info"), agent)

	if msg.Type == "backup_done" || msg.Type == "routine_start" {
		sendHouseKeeperEvent(&msg, issuer)
	}

	return http.StatusOK, "ok"
}

// serveAgentMesssages : Returns the last messages from a given agent
func serveAgentMesssages(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
	messages, err := getLastMessages(fmt.Sprintf("messages:%s", context.Param("id")))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, messages
}

// serveAgentBackups : Returns the backup_done messages from a given agent
func serveAgentBackups(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
	keys, err := services.Db.Keys(fmt.Sprintf("messages:%s*", context.Param("id"))).Result()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	backupMsg := make([]*models.AgentMessage, 0)
	for _, key := range keys {
		messages, err := getLastMessages(key)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		for _, msg := range messages {
			typedMsg := msg.(*models.AgentMessage)
			if typedMsg.Type == "backup_done" {
				backupMsg = append(backupMsg, typedMsg)
			}
		}
	}

	return http.StatusOK, backupMsg
}

func getLastMessages(id string) ([]interface{}, error) {
	return services.DbZRevRange(id, 0, 100, models.AgentMessage{})
}

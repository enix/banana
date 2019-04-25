package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
)

// ReceiveAgentMesssage : Check and store an agent's message
func ReceiveAgentMesssage(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	msg := models.AgentMessage{}
	body := services.ReadBytesFromStream(context.Request.Body)
	json.Unmarshal(body, &msg)

	err := msg.VerifySignature(issuer.Certificate)
	if err != nil {
		return http.StatusForbidden, err
	}
	issuerID := fmt.Sprintf("%s:%s", issuer.Organization, issuer.CommonName)
	if msg.Info.SenderID != issuerID {
		return http.StatusForbidden, fmt.Errorf("sender_id / certificate DN mismatch : [%s] vs [%s]", msg.Info.SenderID, issuerID)
	}

	msg.Info.SenderID = fmt.Sprintf("%s:%s", issuer.Organization, issuer.CommonName)
	services.DbZAdd(msg.Info.GetFullKey(), msg.Info.GetSortedSetScore(), msg)

	agent := models.NewAgent(issuer.Organization, issuer.CommonName, msg)
	services.DbSet(agent.GetFullKeyFor("info"), agent)
	return http.StatusOK, "ok"
}

// ServeAgentMesssages : Returns the last messages from a given agent
func ServeAgentMesssages(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	zkey := fmt.Sprintf("messages:%s", context.Param("id"))
	messages, err := services.DbZRevRange(zkey, 0, 100, models.AgentMessage{})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, messages
}

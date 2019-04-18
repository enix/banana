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
	body := services.ReadBytesFromStream(context.Request.Body)
	msg := models.Message{}
	msg.SenderID = fmt.Sprintf("%s:%s", issuer.Organization, issuer.CommonName)
	json.Unmarshal(body, &msg)
	services.DbZAdd(msg.GetFullKey(), msg.GetSortedSetScore(), msg)
	return http.StatusOK, "ok"
}

// ServeAgentMesssages : Returns the last messages from a given agent
func ServeAgentMesssages(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	zkey := fmt.Sprintf("messages:%s", context.Param("id"))
	messages, err := services.DbZRange(zkey, 0, 10, models.Message{})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, messages
}

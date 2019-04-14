package routes

import (
	"fmt"
	"net/http"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
)

// ReceiveAgentMesssage : Check and store an agent's message
func ReceiveAgentMesssage(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	body := services.ParseJSONFromStream(context.Request.Body)
	msg := models.NewMessage(body)
	services.DbSet(msg.GetFullKey(), msg)
	return http.StatusOK, "ok"
}

// ServeAgentMesssages : Returns the last messages from a given agent
// TODO: maybe handle this better cause .Keys() is O(N)
// and the for loop can probably be avoided by a .GetAll() call or
// someting similar
func ServeAgentMesssages(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	pattern := fmt.Sprintf("message:%s:%s", issuer.Organization, issuer.CommonName)
	messages := services.Db.Keys(pattern).Val()
	response := make([]models.Message, 0, len(messages))

	for _, elem := range messages {
		response = append(response, services.DbGet(elem).(models.Message))
	}

	return http.StatusOK, response
}

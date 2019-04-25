package routes

import (
	"enix.io/banana/src/models"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes : All API endpoints
func InitializeRoutes(router *gin.Engine) {
	router.GET("/ping", handleClientRequest(handlePingRequest))
	router.GET("/agents", handleClientRequest(ServeAgentList))
	router.GET("/agents/:id", handleClientRequest(ServeAgent))
	router.GET("/agents/:id/messages", handleClientRequest(ServeAgentMesssages))
	router.POST("/agents/notify", handleClientRequest(ReceiveAgentMesssage))
	router.GET("/housekeeper/ws", handleHouseKeeperConnection)
	router.GET("/test", func(context *gin.Context) {
		msg := models.HouseKeeperMessage{
			Config: models.Config{
				TTL: 3,
			},
		}

		houseKeeperEvents <- msg
		context.JSON(200, msg)
	})
}

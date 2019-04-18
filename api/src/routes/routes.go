package routes

import "github.com/gin-gonic/gin"

// InitializeRoutes : All API endpoints
func InitializeRoutes(router *gin.Engine) {
	router.GET("/ping", handleClientRequest(handlePingRequest))
	router.GET("/agents", handleClientRequest(ServeAgentList))
	router.POST("/agents/notify", handleClientRequest(ReceiveAgentMesssage))
	router.GET("/agents/:id", handleClientRequest(ServeAgent))
	router.GET("/agents/:id/messages", handleClientRequest(ServeAgentMesssages))
}

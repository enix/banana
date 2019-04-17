package routes

import "github.com/gin-gonic/gin"

// InitializeRoutes : All API endpoints
func InitializeRoutes(router *gin.Engine) {
	router.GET("/ping", handleClientRequest(handlePingRequest))
	router.POST("/ping", handleSignedRequest(handlePingRequest))
	router.GET("/containers", handleClientRequest(ServeBackupContainerList))
	router.GET("/containers/:containerName", handleClientRequest(ServeBackupContainer))
	router.GET("/containers/:containerName/tree/:treeName", handleClientRequest(ServeBackupTree))
	router.GET("/agents", handleClientRequest(ServeAgentList))
	router.POST("/agents", handleClientRequest(RegisterAgent))
	router.POST("/agents/:id/notify", handleClientRequest(ReceiveAgentMesssage))
}

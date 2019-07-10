package routes

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes : All API endpoints
func InitializeRoutes(router *gin.Engine) {
	router.GET("/ping", handleClientRequest(handlePingRequest))
	router.GET("/agents", handleClientRequest(serveAgentList))
	router.GET("/agents/:id", handleClientRequest(serveAgent))
	router.GET("/agents/:id/messages", handleClientRequest(serveAgentMesssages))
	// router.GET("/agents/:id/backups", handleClientRequest(serveAgentBackups))
	router.GET("/agents/:id/backups", handleClientRequest(serveBackupMetadata))
	router.POST("/agents/notify", handleClientRequest(receiveAgentMesssage))
	router.POST("/agents/metadata/:id", handleClientRequest(receiveBackupMetadata))
	router.GET("/housekeeper/ws", handleHouseKeeperConnection)
}

package routes

import (
	"net/http"

	"enix.io/banana/src/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RequestHandler : Shortand for func type that handle client requests
type RequestHandler = func(*gin.Context) (int, interface{})

func handleClientRequest(handler RequestHandler) func(context *gin.Context) {
	return func(context *gin.Context) {
		status, data := handler(context)

		if dataAsError, ok := data.(error); ok {
			context.JSON(status, map[string]string{"error": dataAsError.Error()})
			return
		}

		if dataAsString, ok := data.(string); ok {
			context.JSON(status, map[string]string{"response": dataAsString})
			return
		}

		context.JSON(status, data)
	}
}

func handlePingRequest(context *gin.Context) (int, interface{}) {
	return http.StatusOK, "pong"
}

// InitializeRouter : Initialize all server routes
func InitializeRouter(store *storage.ObjectStorage) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", handleClientRequest(handlePingRequest))
	router.GET("/backups", handleClientRequest(ServeContainerList(store)))
	router.GET("/backups/:bucketName", handleClientRequest(ServeBackupTreeListFromContainer(store)))

	return router, nil
}

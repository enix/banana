package main

import (
	"net/http"

	"enix.io/banana/src/routes"
	"enix.io/banana/src/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func handleClientRequest(handler func(*gin.Context) (int, interface{})) func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(handler(context))
	}
}

func handlePingRequest(context *gin.Context) (int, interface{}) {
	return http.StatusOK, map[string]string{
		"response": "pong",
	}
}

// InitializeRouter : Initialize all server routes
func InitializeRouter(store *storage.ObjectStorage) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", handleClientRequest(handlePingRequest))
	router.GET("/buckets", handleClientRequest(routes.ServeBucketsList(store)))

	return router, nil
}

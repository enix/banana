package routes

import (
	"fmt"
	"net/http"

	"enix.io/banana/src/logger"
	"enix.io/banana/src/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RequestIssuer : Used to authenticate requests
type RequestIssuer struct {
	CommonName   string
	Organization string
	Signature    string
}

// RequestHandler : Shorthand for func type that handle client requests
type RequestHandler = func(*gin.Context, *RequestIssuer) (int, interface{})

func authenticateClientRequest(context *gin.Context) (*RequestIssuer, error) {
	cname, err := services.GetDNFieldValue(context, "CN")
	if err != nil {
		return nil, err
	}

	oname, err := services.GetDNFieldValue(context, "O")
	if err != nil {
		return nil, err
	}

	client := &RequestIssuer{
		CommonName:   cname,
		Organization: oname,
		Signature:    context.GetHeader("X-Signature"),
	}

	return client, nil
}

func handleClientRequest(handler RequestHandler) func(*gin.Context) {
	return func(context *gin.Context) {
		client, err := authenticateClientRequest(context)
		if err != nil {
			fmt.Println(err)
			context.JSON(401, map[string]string{"error": "invalid authentication data"})
			return
		}
		logger.Log("Hello, %s from %s\n", client.CommonName, client.Organization)

		status, data := handler(context, client)
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

func handleSignedRequest(handler RequestHandler) func(*gin.Context) {
	return func(context *gin.Context) {
		rawData, err := context.GetRawData()
		if err != nil {
			context.JSON(400, map[string]string{"error": err.Error()})
			return
		}

		signature := context.GetHeader("X-Signature")
		if len(signature) == 0 {
			context.JSON(400, map[string]string{"error": "missing X-Signature header"})
			return
		}

		err = services.VerifySha256Signature(rawData, signature, context.GetHeader("X-Client-Certificate"))
		if err != nil {
			logger.LogError(err)
			context.JSON(401, map[string]string{"error": "invalid signature"})
			return
		}

		handleClientRequest(handler)(context)
	}
}

func handlePingRequest(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	return http.StatusOK, map[string]string{
		"issuer":       issuer.CommonName,
		"organization": issuer.Organization,
		"signature":    issuer.Signature,
		"data":         "pong",
	}
}

// InitializeRouter : Initialize all server routes
func InitializeRouter() (*gin.Engine, error) {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", handleClientRequest(handlePingRequest))
	router.POST("/ping", handleSignedRequest(handlePingRequest))
	router.GET("/containers", handleClientRequest(ServeBackupContainerList))
	router.GET("/containers/:containerName", handleClientRequest(ServeBackupContainer))
	router.GET("/containers/:containerName/tree/:treeName", handleClientRequest(ServeBackupTree))

	return router, nil
}

package routes

import (
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
	Certificate  string
}

// RequestHandler : Shorthand for func type that handle client requests
type RequestHandler = func(*gin.Context, *RequestIssuer) (int, interface{})

func authenticateClientRequest(context *gin.Context) (*RequestIssuer, error) {
	dn := context.GetHeader("X-Client-Subject-DN")
	cname, err := services.GetDNFieldValue(dn, "CN")
	if err != nil {
		return nil, err
	}

	oname, err := services.GetDNFieldValue(dn, "O")
	if err != nil {
		return nil, err
	}

	client := &RequestIssuer{
		CommonName:   cname,
		Organization: oname,
		Certificate:  context.GetHeader("X-Client-Certificate"),
	}

	return client, nil
}

func handleClientRequest(handler RequestHandler) func(*gin.Context) {
	return func(context *gin.Context) {
		client, err := authenticateClientRequest(context)
		if err != nil {
			context.JSON(401, map[string]string{"error": "invalid authentication data"})
			return
		}
		logger.Log("Hello, %s from %s", client.CommonName, client.Organization)

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

func handlePingRequest(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	return http.StatusOK, map[string]string{
		"issuer":       issuer.CommonName,
		"organization": issuer.Organization,
		"data":         "pong",
	}
}

// InitializeRouter : Initialize router and API endpoints
func InitializeRouter() (*gin.Engine, error) {
	router := gin.Default()
	router.Use(cors.Default())
	InitializeRoutes(router)
	return router, nil
}

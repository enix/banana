package routes

import (
	"fmt"
	"net/http"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
)

// serveAgent : Returns informations about a specific agent
func serveAgent(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
	var agent models.Agent
	err := services.DbGet("agent:info:"+context.Param("id"), &agent)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("agent %s not found", context.Param("id"))
	}

	return http.StatusOK, agent
}

// serveAgentList : Returns the agent list
func serveAgentList(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
	keys, err := services.Db.Keys("agent*").Result()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if len(keys) < 1 {
		return http.StatusOK, make([]interface{}, 0)
	}

	agents, err := services.DbMGet(keys, models.Agent{})
	if err != nil {
		return http.StatusNotFound, err
	}
	return http.StatusOK, agents
}

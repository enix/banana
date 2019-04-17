package routes

import (
	"net/http"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
)

// RegisterAgent : Add a new agent to the agent list
func RegisterAgent(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	agent := models.NewAgent(issuer.Organization, issuer.Organization)
	services.DbSet(agent.GetFullKey(), agent)

	agents := services.DbGet("agents").([]models.Agent)
	agents = append(agents, *agent)
	services.DbSet("agents", agents)
	return http.StatusOK, agent
}

// ServeAgentList : Returns the agent list
func ServeAgentList(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	return http.StatusOK, services.DbGet("agents")
}
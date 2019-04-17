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

	var agents []models.Agent
	services.DbGet("agents", &agents)
	agents = append(agents, *agent)

	services.DbSet("agents", agents)
	return http.StatusOK, agent
}

// ServeAgentList : Returns the agent list
func ServeAgentList(context *gin.Context, issuer *RequestIssuer) (int, interface{}) {
	var agents []models.Agent
	services.DbGet("agents", &agents)
	return http.StatusOK, agents
}

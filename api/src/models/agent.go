package models

import (
	"fmt"
)

// Agent : Representation of an agent
type Agent struct {
	Organization string       `json:"organization"`
	CommonName   string       `json:"cn"`
	LastMessage  AgentMessage `json:"last_message"`
}

// NewAgent : Convenience function for creating an agent
func NewAgent(orga, cn string, lastMsg AgentMessage) *Agent {
	return &Agent{
		Organization: orga,
		CommonName:   cn,
		LastMessage:  lastMsg,
	}
}

// GetFullKeyFor : Generate the key that will be used to store within redis
func (agent *Agent) GetFullKeyFor(field string) string {
	return fmt.Sprintf("agent:%s:%s:%s", field, agent.Organization, agent.CommonName)
}

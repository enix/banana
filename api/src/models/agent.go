package models

import (
	"fmt"

	"github.com/rs/xid"
)

// Agent : Representation of an agent
type Agent struct {
	ID           xid.ID   `json:"id"`
	Organization string   `json:"organization"`
	CommonName   string   `json:"cn"`
	Messages     []xid.ID `json:"messages"`
}

// NewAgent : Convenience function for creating a message
func NewAgent(orga, cn string) *Agent {
	return &Agent{xid.New(), orga, cn, make([]xid.ID, 0)}
}

// GetFullKey : Generate the key that will be used to store within redis
func (agent *Agent) GetFullKey() string {
	return fmt.Sprintf("agent:%s", agent.ID)
}

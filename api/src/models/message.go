package models

import (
	"fmt"

	"github.com/rs/xid"
)

// Message : Representation of an agent's notification
type Message struct {
	ID   xid.ID                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

// NewMessage : Convenience function for creating a message
func NewMessage(data map[string]interface{}) *Message {
	return &Message{xid.New(), data}
}

// GetFullKey : Generate the key that will be used to store within redis
func (msg *Message) GetFullKey() string {
	return fmt.Sprintf("message:%s", msg.ID)
}

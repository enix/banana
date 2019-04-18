package models

import (
	"fmt"
)

// Message : Representation of an agent's notification
type Message struct {
	SenderID  string                 `json:"sender_id"`
	Version   string                 `json:"version"`
	Timestamp int64                  `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}

// NewMessage : Convenience function for creating a message
func NewMessage(timestamp int64, data map[string]interface{}) *Message {
	return &Message{
		Version:   "1",
		Timestamp: timestamp,
		Data:      data,
	}
}

// GetSortedSetScore : Generate the score that will be used to store within redis sorted set
func (msg *Message) GetSortedSetScore() float64 {
	return float64(msg.Timestamp)
}

// GetFullKey : Get the key to be used within redis
func (msg *Message) GetFullKey() string {
	return fmt.Sprintf("messages:%s", msg.SenderID)
}

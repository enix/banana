package models

import (
	"fmt"
)

// Message : Representation of an agent's notification
type Message struct {
	Version   int8   `json:"version"`
	Timestamp int64  `json:"timestamp"`
	SenderID  string `json:"sender_id"`
	Type      string `json:"type"`
}

// AgentMessage : Representation of an agent's notification
type AgentMessage struct {
	Info      Message                `json:"info"`
	Config    Config                 `json:"config"`
	Command   map[string]interface{} `json:"command"`
	Logs      string                 `json:"logs"`
	Signature string                 `json:"signature,omitempty"`
}

// HouseKeeperMessage : Representation of a housekeeper event
type HouseKeeperMessage struct {
	Info      Message                `json:"info"`
	Config    Config                 `json:"config"`
	Command   map[string]interface{} `json:"command"`
	Signature string                 `json:"signature,omitempty"`
}

// GetSortedSetScore : Generate the score that will be used to store within redis sorted set
func (msg *Message) GetSortedSetScore() float64 {
	return float64(msg.Timestamp)
}

// GetFullKey : Get the key to be used within redis
func (msg *Message) GetFullKey() string {
	return fmt.Sprintf("messages:%s", msg.SenderID)
}

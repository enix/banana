package models

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"enix.io/banana/src/logger"
	"enix.io/banana/src/services"
)

// Message : Representation of an agent's notification
type Message struct {
	Version   int8                   `json:"version"`
	Timestamp int64                  `json:"timestamp"`
	SenderID  string                 `json:"sender_id"`
	Type      string                 `json:"type"`
	Config    map[string]interface{} `json:"config"`
	Command   map[string]interface{} `json:"command"`
	Logs      string                 `json:"logs"`
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

// VerifySignature : Verify that the signature match the message's content
func (msg *Message) VerifySignature(cert string) error {
	sig := msg.Signature
	msg.Signature = ""
	rawMessage, _ := json.Marshal(msg)
	err := services.VerifySha256Signature(rawMessage, sig, cert)
	if err != nil {
		logger.LogError(err)
		return errors.New("invalid signature")
	}
	msg.Signature = sig
	return nil
}

// Sign : Marshal the struct and generate signature from the result
func (msg *Message) Sign(privkey *rsa.PrivateKey) error {
	rawMessage, _ := json.Marshal(msg)
	hash := sha256.New()
	hash.Write(rawMessage)
	digest := hash.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, privkey, crypto.SHA256, digest)
	if err != nil {
		return err
	}
	base64sig := make([]byte, base64.StdEncoding.EncodedLen(len(sig)))
	base64.StdEncoding.Encode(base64sig, sig)
	msg.Signature = string(base64sig)
	return err
}

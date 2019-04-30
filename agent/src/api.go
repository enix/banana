package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
)

// SendToMonitor : Sign given message and POST it to the monitor API
func SendToMonitor(config *models.Config, message *models.AgentMessage) error {
	fmt.Print("waiting for monitor... ")

	httpClient := services.GetHTTPClient()
	dn := services.Credentials.Cert.Subject.ToRDNSequence().String()
	oname, _ := services.GetDNFieldValue(dn, "O")
	cname, _ := services.GetDNFieldValue(dn, "CN")
	message.Info.SenderID = fmt.Sprintf("%s:%s", oname, cname)

	url := fmt.Sprintf("%s/agents/notify", config.MonitorURL)
	rawMessage, _ := json.Marshal(message)
	res, err := httpClient.Post(url, "application/json", strings.NewReader(string(rawMessage)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	if res.StatusCode != 200 {
		os.Exit(1)
	}
	return nil
}

// SendMessageToMonitor : Convenience function to create and send a message
func SendMessageToMonitor(typ string, config *models.Config, cmd Command, logs string) {
	msg := &models.AgentMessage{
		Info: models.Message{
			Version:   1,
			Timestamp: time.Now().Unix(),
			Type:      typ,
		},
		Config:  *config,
		Command: cmd.JSONMap(),
		Logs:    logs,
	}

	msg.Signature, _ = msg.Config.Sign(services.Credentials.PrivateKey)
	err := SendToMonitor(config, msg)
	if err != nil {
		log.Fatal(err)
	}
}

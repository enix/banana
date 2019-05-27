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

// sendToMonitor : Sign given message and POST it to the monitor API
func sendToMonitor(config *models.Config, message *models.AgentMessage) error {
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

// sendMessageToMonitor : Convenience function to create and send a message
func sendMessageToMonitor(typ string, config *models.Config, cmd command, logs string) {
	msg := &models.AgentMessage{
		Info: models.Message{
			Version:   1,
			Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
			Type:      typ,
		},
		Config:  *config,
		Command: cmd.jsonMap(),
		Logs:    logs,
	}

	msg.Signature, _ = msg.Config.Sign(services.Credentials.PrivateKey)
	err := sendToMonitor(config, msg)
	if err != nil {
		log.Fatal(err)
	}
}

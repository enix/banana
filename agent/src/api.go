package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"k8s.io/klog"
)

// sendToMonitor : Sign given message and POST it to the monitor API
func sendToMonitor(config *models.Config, message *models.AgentMessage) error {
	fmt.Print("waiting for monitor... ")

	httpClient := services.GetHTTPClient()
	oname := services.Credentials.Cert.Subject.Organization[0]
	cname := services.Credentials.Cert.Subject.CommonName
	message.SenderID = fmt.Sprintf("%s:%s", oname, cname)

	url := fmt.Sprintf("%s/agents/notify", config.MonitorURL)
	rawMessage, _ := json.Marshal(message)
	res, err := httpClient.Post(url, "application/json", strings.NewReader(string(rawMessage)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	if res.StatusCode != 200 {
		err, _ := ioutil.ReadAll(res.Body)
		klog.Error(string(err))
		os.Exit(1)
	}
	return nil
}

// sendMessageToMonitor : Convenience function to create and send a message
func sendMessageToMonitor(typ string, config *models.Config, cmd command, logs string) {
	if config.MonitorURL == "" {
		klog.Fatal("monitor URL not set, set using -m")
	}

	rawCommand := map[string]interface{}{}
	if cmd != nil {
		rawCommand = cmd.jsonMap()
	}

	msg := &models.AgentMessage{
		Message: models.Message{
			Version:   1,
			Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
			Type:      typ,
		},
		Config:  *config,
		Command: rawCommand,
		Logs:    logs,
	}

	msg.Signature, _ = msg.Config.Sign(services.Credentials.PrivateKey)
	err := sendToMonitor(config, msg)
	if err != nil {
		klog.Fatal(err)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"k8s.io/klog"
)

func fireAPIRequest(config *models.Config, url string, data []byte) (string, error) {
	fmt.Print("waiting for monitor... ")
	httpClient := services.GetHTTPClient(config.SkipTLSVerify)
	res, err := httpClient.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	response, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		klog.Error(string(response))
		return "", errors.New(string(response))
	}

	return string(response), nil
}

// sendToMonitor : Sign given message and POST it to the monitor API
func sendToMonitor(config *models.Config, message *models.AgentMessage) (string, error) {
	oname := services.Credentials.Cert.Subject.Organization[0]
	cname := services.Credentials.Cert.Subject.CommonName
	message.SenderID = fmt.Sprintf("%s:%s", oname, cname)
	url := fmt.Sprintf("%s/agents/notify", config.MonitorURL)
	rawMessage, _ := json.Marshal(message)

	return fireAPIRequest(config, url, rawMessage)
}

// sendMessageToMonitor : Convenience function to create and send a message
func sendMessageToMonitor(
	typ string,
	config *models.Config,
	cmd command,
	metadata *models.BackupMetadata,
	logs string) string {

	if config.MonitorURL == "" {
		assert(errors.New("monitor URL not set, set using -m"))
	}

	rawCommand := map[string]interface{}{}
	if cmd != nil {
		rawCommand = cmd.jsonMap()
	}

	if metadata == nil {
		metadata = &models.BackupMetadata{}
	}

	msg := &models.AgentMessage{
		Message: models.Message{
			Version:   1,
			Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
			Type:      typ,
		},
		Config:   *config,
		Command:  rawCommand,
		Metadata: *metadata,
		Logs:     logs,
	}

	msg.Signature, _ = msg.Config.Sign(services.Credentials.PrivateKey)
	res, err := sendToMonitor(config, msg)
	assert(err)
	return res
}

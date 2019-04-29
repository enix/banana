package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"enix.io/banana/src/logger"
	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gorilla/websocket"
)

var pending = make(map[string]*models.HouseKeeperMessage)

func watchPendingBackups() {
	for {
		time.Sleep(time.Millisecond * 1000)
		if len(pending) == 0 {
			continue
		}

		logger.Log("checking for expired TTLs...")
		now := time.Now().UTC().Unix()
		for key, value := range pending {
			if now-value.Info.Timestamp > value.Config.TTL {
				removeFromStorage(value)
				delete(pending, key)
			}
		}
	}
}

func openSocketConnection() *websocket.Conn {
	url := &url.URL{
		Scheme: "ws",
		Host:   "api.banana.enix.io:443",
		Path:   "/housekeeper/ws",
	}
	socket, err := tls.Dial("tcp", url.Host, services.GetTLSConfig())
	assert(err)
	conn, _, err := websocket.NewClient(socket, url, nil, 1024, 1024)
	assert(err)

	return conn
}

func synchroniseBackups() {
	dn := services.Credentials.Cert.Subject.ToRDNSequence().String()
	org, err := services.GetDNFieldValue(dn, "O")
	assert(err)
	url := fmt.Sprintf("%s/agents/%s/backups", "https://api.banana.enix.io", org)
	httpClient := services.GetHTTPClient()
	res, err := httpClient.Get(url)
	assert(err)
	defer res.Body.Close()

	messages := make([]models.HouseKeeperMessage, 0)
	err = json.Unmarshal(services.ReadBytesFromStream(res.Body), &messages)
	assert(err)

	for index := range messages {
		handleMessage(&messages[index])
	}
}

func listenForMessages(conn *websocket.Conn) {
	msg := models.HouseKeeperMessage{}

	for {
		err := conn.ReadJSON(&msg)
		assert(err)
		handleMessage(&msg)
	}
}

func handleMessage(msg *models.HouseKeeperMessage) {
	pending[msg.Signature] = msg
	logger.Log("new backup added to pending, TTL: %d", msg.Config.TTL)
}

func removeFromStorage(msg *models.HouseKeeperMessage) {
	logger.Log("%+v", msg)
	// services.DeleteObject(msg.Config.BucketName, fmt.Sprintf("%s/%s", msg.Command.Name, msg.C))
}

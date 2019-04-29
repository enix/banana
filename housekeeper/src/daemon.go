package main

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"time"

	"enix.io/banana/src/logger"
	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gorilla/websocket"
)

var pending = make(map[string]*models.Config)

func watchPendingBackups() {
	for {
		now := time.Now()
		time.Sleep(time.Millisecond * 1000)
		if len(pending) == 0 {
			continue
		}

		logger.Log("checking for expired TTLs...")
		delta := int64(time.Since(now).Seconds())
		now = time.Now()
		for key, value := range pending {
			value.TTL -= delta

			if value.TTL <= 0 {
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

func listenForMessages(conn *websocket.Conn) {
	msg := models.HouseKeeperMessage{}

	for {
		err := conn.ReadJSON(&msg)
		assert(err)
		handleMessage(&msg)
	}
}

func handleMessage(msg *models.HouseKeeperMessage) {
	pending[msg.Signature] = &msg.Config
	fmt.Printf("new backup added to pending, TTL: %d\n", msg.Config.TTL)
}

func removeFromStorage(config *models.Config) {
	logger.Log("removing %+v\n", config)
}

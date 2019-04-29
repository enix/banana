package routes

import (
	"fmt"

	"enix.io/banana/src/logger"
	"enix.io/banana/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var houseKeeperStreams = make(map[string]chan models.HouseKeeperMessage)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// HandleHouseKeeperConnection : Upgrade HTTP to WS and send events to the housekeeper
func handleHouseKeeperConnection(context *gin.Context) {
	conn, err := wsUpgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		fmt.Printf("failed to set websocket upgrade: %+v\n", err)
		return
	}

	issuer, err := authenticateClientRequest(context)
	houseKeeperStreams[issuer.Organization] = make(chan models.HouseKeeperMessage)

	for {
		conn.WriteJSON(<-houseKeeperStreams[issuer.Organization])
	}
}

func sendHouseKeeperEvent(msg *models.AgentMessage, issuer *RequestIssuer) {
	event := models.HouseKeeperMessage{
		Info:      msg.Info,
		Config:    msg.Config,
		Command:   msg.Command,
		Signature: msg.Signature,
	}

	select {
	case houseKeeperStreams[issuer.Organization] <- event:
	default:
		logger.Log("warning: no housekeeper was listening for last backup")
	}
}

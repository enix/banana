package routes

import (
	"fmt"

	"enix.io/banana/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var houseKeeperEvents = make(chan models.HouseKeeperMessage)

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

	for {
		conn.WriteJSON(<-houseKeeperEvents)
	}
}

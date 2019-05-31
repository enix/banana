package routes

import (
	"fmt"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"k8s.io/klog"
)

var houseKeeperStreams = make(map[string]chan models.HouseKeeperMessage)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// handleHouseKeeperConnection : Upgrade HTTP to WS and send events to the housekeeper
func handleHouseKeeperConnection(context *gin.Context) {
	conn, err := wsUpgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		klog.Infof("failed to set websocket upgrade: %+v\n", err)
		return
	}

	res := make(map[string]string)
	issuer, err := authenticateClientRequest(context)
	houseKeeperStreams[issuer.Organization] = make(chan models.HouseKeeperMessage)

	go func() {
		keys, _ := services.Db.Keys(fmt.Sprintf("messages:%s*", issuer.Organization)).Result()
		for _, key := range keys {
			messages, _ := getLastMessages(key)
			for _, msg := range messages {
				typedMsg := msg.(*models.AgentMessage)
				if typedMsg.Type == "backup_done" {
					sendHouseKeeperEvent(typedMsg, issuer)
				}
			}
		}
	}()

	for {
		msg := <-houseKeeperStreams[issuer.Organization]
		conn.WriteJSON(msg)
		conn.ReadJSON(&res)

		if len(res["error"]) > 0 {
			services.SendAlert("arthur.chaloin@enix.fr", res["error"])
		}
	}
}

func sendHouseKeeperEvent(msg *models.AgentMessage, issuer *requestIssuer) {
	event := models.HouseKeeperMessage{
		Message:   msg.Message,
		Config:    msg.Config,
		Command:   msg.Command,
		Signature: msg.Signature,
	}

	select {
	case houseKeeperStreams[issuer.Organization] <- event:
	default:
		klog.Warning("no housekeeper was listening for last backup")
	}
}

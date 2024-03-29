package main

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"net/url"
	"time"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"k8s.io/klog"
)

var pending = make(map[int64]*models.HouseKeeperMessage)

func watchPendingBackups() {
	for {
		time.Sleep(time.Millisecond * 1000)
		if len(pending) == 0 {
			continue
		}

		klog.Info("checking for expired TTLs...")
		now := time.Now().UTC().Unix()
		for key, value := range pending {
			if now-value.Timestamp > value.Config.TTL {
				removeFromStorage(value)
				delete(pending, key)
				klog.Infof("%d remaining backup(s)", len(pending))
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
	if err != nil {
		klog.Error(err)
		time.Sleep(3 * time.Second)
		return openSocketConnection()
	}
	conn, _, err := websocket.NewClient(socket, url, nil, 1024, 1024)
	assert(err)

	return conn
}

func listenForMessages(conn *websocket.Conn) {
	msg := models.HouseKeeperMessage{}

	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			klog.Error(errors.New("disconnected from monitor, retrying"))
			conn = openSocketConnection()
			defer conn.Close()
			continue
		}

		err = handleMessage(&msg)
		if err == nil {
			conn.WriteJSON(map[string]string{"response": "ok"})
		} else {
			klog.Error(err)
			conn.WriteJSON(map[string]string{"error": err.Error()})
		}
	}
}

func handleMessage(msg *models.HouseKeeperMessage) error {
	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: []byte(services.Credentials.Cert.Raw)})
	err := msg.Config.VerifySignature(string(pemCert), msg.Signature)
	if err != nil {
		klog.Error(err)
	}

	if msg.Type == "backup_done" {
		pending[msg.Timestamp] = msg
		klog.Infof("new backup added to pending, TTL: %d", msg.Config.TTL)
	} else if msg.Type == "routine_start" {
		fmt.Println(msg.Config)
	}

	return nil
}

// this function was hard-coded for duplicity-formatted backups
func removeFromStorage(msg *models.HouseKeeperMessage) {
	manifestFilename := fmt.Sprintf("%s/duplicity-full.%s.manifest.gpg", msg.Command["name"], msg.Command["OpaqueID"])
	diffFilename := fmt.Sprintf("%s/duplicity-full.%s.vol1.difftar.gpg", msg.Command["name"], msg.Command["OpaqueID"])
	sigFilename := fmt.Sprintf("%s/duplicity-full-signatures.%s.sigtar.gpg", msg.Command["name"], msg.Command["OpaqueID"])

	fmt.Println()
	klog.Infof("deleting backup %s from %s in bucket %s", msg.Command["OpaqueID"], msg.Command["name"], msg.Config.BucketName)
	klog.Infof("the following files will be deleted: \n\t* %s\n\t* %s\n\t* %s\n", manifestFilename, diffFilename, sigFilename)

	_, err := services.Storage.DeleteObject(&msg.Config.BucketName, &manifestFilename)
	if err != nil {
		klog.Error(errors.Wrap(err, "backup could not be deleted, this is not normal"))
	}
	_, err = services.Storage.DeleteObject(&msg.Config.BucketName, &diffFilename)
	if err != nil {
		klog.Error(errors.Wrap(err, "backup could not be deleted, this is not normal"))
	}
	_, err = services.Storage.DeleteObject(&msg.Config.BucketName, &sigFilename)
	if err != nil {
		klog.Error(errors.Wrap(err, "backup could not be deleted, this is not normal"))
	}
}

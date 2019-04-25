package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gorilla/websocket"
)

func logFatal(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", fmt.Sprintf("error: %s", err.Error()))
	os.Exit(1)
}

func assert(err error) {
	if err != nil {
		logFatal(err)
	}
}

func loadCredentialsToMem() {
	privkeyBytes, err := ioutil.ReadFile("../../security/out/the.agent.key")
	assert(err)
	privkeyBlock, _ := pem.Decode([]byte(privkeyBytes))
	privkey, err := x509.ParsePKCS1PrivateKey(privkeyBlock.Bytes)
	assert(err)

	certBytes, err := ioutil.ReadFile("../../security/out/the.agent.pem")
	assert(err)
	certBlock, _ := pem.Decode([]byte(certBytes))
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	assert(err)

	cacertBytes, err := ioutil.ReadFile("../../security/ca/ca.pem")
	assert(err)
	cacertBlock, _ := pem.Decode([]byte(cacertBytes))
	cacert, err := x509.ParseCertificate(cacertBlock.Bytes)
	assert(err)

	services.Credentials = &services.APICredentials{
		PrivateKey: privkey,
		Cert:       cert,
		CaCert:     cacert,
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
	msg := models.Message{}

	for {
		err := conn.ReadJSON(&msg)
		assert(err)
		handleMessage(&msg)
	}
}

func handleMessage(msg *models.Message) {
	fmt.Println(msg)
}

func main() {
	loadCredentialsToMem()
	conn := openSocketConnection()
	defer conn.Close()

	listenForMessages(conn)
}

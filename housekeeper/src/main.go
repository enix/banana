package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"k8s.io/klog"

	"enix.io/banana/src/services"
)

func assert(err error) {
	if err != nil {
		klog.Fatal(err)
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

func main() {
	err := services.OpenVaultConnection()
	assert(err)
	err = services.OpenStorageConnection()
	assert(err)

	loadCredentialsToMem()
	conn := openSocketConnection()
	defer conn.Close()

	go watchPendingBackups()
	listenForMessages(conn)
}

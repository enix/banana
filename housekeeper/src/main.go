package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"enix.io/banana/src/services"
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

func main() {
	err := services.OpenVaultConnection()
	assert(err)
	err = services.OpenStorageConnection()
	assert(err)

	bucket := "banana-test2"
	object := "etc/duplicity-full-signatures.20190404T130959Z.sigtar.gpg"
	services.Storage.DeleteObject(&bucket, &object)

	// loadCredentialsToMem()
	// conn := openSocketConnection()
	// defer conn.Close()

	// go watchPendingBackups()
	// listenForMessages(conn)
}

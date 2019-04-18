package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
)

// APICredentials : Wrapper struct that contains the required certs and key
// to authenticate with the monitor API and sign messages
type APICredentials struct {
	PrivateKey *rsa.PrivateKey
	Cert       *x509.Certificate
	CaCert     *x509.Certificate
}

// SendMessageToMonitor : Sign given message and POST it to the monitor API
func SendMessageToMonitor(config *Config, message *models.Message) error {
	fmt.Println("sending summary to monitor...")

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			tls.Certificate{
				Certificate: [][]byte{Credentials.Cert.Raw},
				PrivateKey:  Credentials.PrivateKey,
			},
		},
	}
	tlsConfig.BuildNameToCertificate()

	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(Credentials.CaCert)

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	httpClient := &http.Client{Transport: transport}

	dn := Credentials.Cert.Subject.ToRDNSequence().String()
	oname, _ := services.GetDNFieldValue(dn, "O")
	cname, _ := services.GetDNFieldValue(dn, "CN")
	message.SenderID = fmt.Sprintf("%s:%s", oname, cname)
	message.Sign(Credentials.PrivateKey)

	url := fmt.Sprintf("%s/agents/notify", config.MonitorURL)
	rawMessage, _ := json.Marshal(message)
	res, err := httpClient.Post(url, "application/json", strings.NewReader(string(rawMessage)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	return nil
}

package services

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Credentials : Used to sign messages and authenticate with Vault/Monitor API from client-side
var Credentials *APICredentials

// APICredentials : Wrapper struct that contains the required certs and key
// to authenticate with the monitor API and sign messages
type APICredentials struct {
	PrivateKey *rsa.PrivateKey
	Cert       *x509.Certificate
	CaCert     *x509.Certificate
}

// GetCertificatePublicKey : Extracts the pubkey from a given url-escaped PEM cert
func GetCertificatePublicKey(base64cert string) (*rsa.PublicKey, error) {
	unescapedCert, err := url.QueryUnescape(base64cert)
	if err != nil {
		return nil, fmt.Errorf("failed to unescape certificate: %s", err.Error())
	}

	block, rest := pem.Decode([]byte(unescapedCert))
	if len(rest) != 0 {
		return nil, errors.New("failed to parse PEM certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse x509 data: %s", err.Error())
	}

	pubkey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("certificate key algorithm is not RSA")
	}

	return pubkey, nil
}

// VerifySha256Signature : Check if data's signature was issued using an agent's private key
func VerifySha256Signature(data []byte, base64sig, base64cert string) error {
	h := sha256.New()
	h.Write(data)
	digest := h.Sum(nil)

	pubkey, err := GetCertificatePublicKey(base64cert)
	if err != nil {
		return err
	}

	signature, err := base64.StdEncoding.DecodeString(base64sig)
	if err != nil {
		return fmt.Errorf("failed to decode signature: %s", err)
	}

	err = rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, digest, signature)
	if err != nil {
		return fmt.Errorf("signature does not match: %s", err)
	}

	return nil
}

// GetTLSConfig : Returns the TLS config for sending requests to the Monitor
func GetTLSConfig() *tls.Config {
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(Credentials.CaCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			tls.Certificate{
				Certificate: [][]byte{Credentials.Cert.Raw},
				PrivateKey:  Credentials.PrivateKey,
			},
		},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true, // TODO: find a way to get CA and remove this
	}
	tlsConfig.BuildNameToCertificate()
	return tlsConfig
}

// GetHTTPClient : Returns the TLS-configured http client for sending requests to the Monitor
func GetHTTPClient() *http.Client {
	transport := &http.Transport{TLSClientConfig: GetTLSConfig()}
	return &http.Client{Transport: transport}
}

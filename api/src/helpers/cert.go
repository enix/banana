package helpers

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
)

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

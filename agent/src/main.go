package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"enix.io/banana/src/models"
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

func loadCredentialsToEnv(config *services.VaultConfig) {
	vault, err := services.NewVaultClient(config)
	assert(err)
	accessToken, err := vault.GetStorageAccessToken()
	assert(err)
	secretToken, err := vault.GetStorageSecretToken()
	assert(err)
	passphrase, err := vault.GetStoragePassphrase()
	assert(err)

	os.Setenv("AWS_ACCESS_KEY_ID", accessToken)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretToken)
	os.Setenv("PASSPHRASE", passphrase)
}

func loadCredentialsToMem(config *models.Config) error {
	privkeyBytes, err := ioutil.ReadFile(config.PrivKeyPath)
	assert(err)
	privkeyBlock, _ := pem.Decode([]byte(privkeyBytes))
	privkey, err := x509.ParsePKCS1PrivateKey(privkeyBlock.Bytes)
	assert(err)

	certBytes, err := ioutil.ReadFile(config.CertPath)
	assert(err)
	certBlock, _ := pem.Decode([]byte(certBytes))
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	assert(err)

	cacertBytes, err := ioutil.ReadFile(config.CaCertPath)
	assert(err)
	cacertBlock, _ := pem.Decode([]byte(cacertBytes))
	cacert, err := x509.ParseCertificate(cacertBlock.Bytes)
	assert(err)

	services.Credentials = &services.APICredentials{
		PrivateKey: privkey,
		Cert:       cert,
		CaCert:     cacert,
	}

	return nil
}

func unloadCredentialsFromEnv() {
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")
}

func main() {
	args := LoadArguments()
	if args.DisplayHelp || len(args.Values) < 1 {
		Usage()
	}

	config := &models.Config{}
	config.LoadDefaults()
	config.LoadFromFile(args.ConfigPath)
	err := config.LoadFromEnv()
	assert(err)
	err = config.LoadFromArgs(&args.Flags)
	assert(err)
	cmd, err := NewCommand(args)
	assert(err)

	loadCredentialsToMem(config)
	loadCredentialsToEnv(&config.Vault)
	err = cmd.Execute(config)
	assert(err)
	unloadCredentialsFromEnv()
}

package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"io/ioutil"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"k8s.io/klog"
)

func assert(err error) {
	if err != nil {
		panic(err)
	}
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

	services.Credentials = &services.APICredentials{
		PrivateKey: privkey,
		Cert:       cert,
	}

	return nil
}

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "1")
	// flag.Parse()

	args := loadArguments()
	if args.DisplayHelp || len(args.Values) < 1 {
		Usage()
	}

	config := &models.Config{}
	config.LoadDefaults()
	if args.Values[0] == "init" {
		config.Backend = ""
		config.BucketName = ""
		config.TTL = 0
	}
	config.LoadFromFile(args.ConfigPath)
	err := config.LoadFromEnv()
	assert(err)
	err = config.LoadFromArgs(&args.Flags)
	assert(err)

	cmd, err := newCommand(args)
	assert(err)

	if args.Values[0] != "init" {
		loadCredentialsToMem(config)
	}
	err = services.OpenVaultConnection(&config.Vault, config.SkipTLSVerify)
	assert(err)

	err = cmd.execute(config)
	assert(err)
}

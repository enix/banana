package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"k8s.io/klog"
)

func assert(err error) {
	if err != nil {
		deleteLockfile()
		if version == "develop" {
			panic(err)
		} else {
			klog.Fatal(err)
		}
	}
}

func createLockfile() {
	pid := os.Getpid()
	bytes, err := ioutil.ReadFile("/tmp/banana.lock")

	if err == nil || !os.IsNotExist(err) {
		klog.Fatalf(
			"refusing to start as PID %s is already running. if it is not the case, you can delete /tmp/banana.lock",
			string(bytes),
		)
	}

	ioutil.WriteFile("/tmp/banana.lock", []byte(strconv.Itoa(pid)), 00644)
}

func deleteLockfile() {
	os.Remove("/tmp/banana.lock")
}

func handleSignals(sigs chan os.Signal) {
	sig := <-sigs
	fmt.Printf("received signal %v, exiting...", sig)
	deleteLockfile()
	os.Exit(0)
}

func loadCredentialsToMem(config *models.Config) error {
	privkeyBytes, err := ioutil.ReadFile(config.PrivKeyPath)
	assert(err)
	privkeyBlock, rest := pem.Decode([]byte(privkeyBytes))
	if len(rest) > 0 {
		assert(errors.New("failed to parse private key. is it in pem format?"))
	}
	privkey, err := x509.ParsePKCS1PrivateKey(privkeyBlock.Bytes)
	assert(err)

	certBytes, err := ioutil.ReadFile(config.CertPath)
	assert(err)
	certBlock, rest := pem.Decode([]byte(certBytes))
	if len(rest) > 0 {
		assert(errors.New("failed to parse certificate. is it in pem format?"))
	}
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

	if args.Values[0] == "version" {
		fmt.Println(version)
		os.Exit(0)
	}

	createLockfile()
	defer deleteLockfile()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go handleSignals(sigs)

	config := &models.Config{}
	config.LoadDefaults()
	if args.Values[0] == "init" {
		config.BucketName = ""
		config.TTL = 0
	}
	err := config.LoadFromFile(args.ConfigPath)
	if err != nil && !os.IsNotExist(err) {
		assert(err)
	}
	err = config.LoadFromEnv()
	assert(err)
	err = config.LoadFromArgs(&args.Flags)
	assert(err)

	cmd, err := newCommand(args)
	assert(err)

	if args.Values[0] != "init" {
		loadCredentialsToMem(config)
	}
	err = services.OpenVaultConnection(config.Vault, config.SkipTLSVerify)
	assert(err)

	err = cmd.execute(config)
	assert(err)
}

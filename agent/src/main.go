package main

import (
	"fmt"
	"os"

	"enix.io/banana/src/helpers"
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

func loadCredentialsToEnv(config *helpers.VaultConfig) {
	vault, err := helpers.NewVaultClient(config)
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

func unloadCredentialsFromEnv() {
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")
}

func main() {
	args := LoadArguments()
	if args.DisplayHelp || len(args.Values) < 1 {
		Usage()
	}

	config := &Config{}
	config.LoadDefaults()
	config.LoadFromFile(args.ConfigPath)
	err := config.LoadFromEnv()
	assert(err)
	err = config.LoadFromArgs(&args.Flags)
	assert(err)
	cmd, err := NewCommand(args)
	assert(err)

	loadCredentialsToEnv(&config.Vault)
	err = cmd.Execute(config)
	assert(err)
	unloadCredentialsFromEnv()
}

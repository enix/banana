package main

import (
	"errors"
	"os"

	"enix.io/banana/src/helpers"
	"enix.io/banana/src/logger"
	"enix.io/banana/src/routes"
	"enix.io/banana/src/storage"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// Assert : Ensure that the given error is a nil pointer
// 					otherwise print it and exit process with status code 1
func Assert(err error) {
	if err != nil {
		logger.LogError(err)
		os.Exit(1)
	}
}

func openStorageAPIConnection() (*storage.ObjectStorage, error) {
	vault, err := helpers.NewVaultClient(&helpers.VaultConfig{
		Addr:       os.Getenv("VAULT_ADDR"),
		Token:      os.Getenv("VAULT_TOKEN"),
		SecretPath: "storage_access",
	})

	Assert(err)
	accessToken, err := vault.GetStorageAccessToken()
	Assert(err)
	secretToken, err := vault.GetStorageSecretToken()
	Assert(err)

	var store storage.ObjectStorage
	store.Connect(
		os.Getenv("API_ENDPOINT"),
		credentials.NewStaticCredentials(accessToken, secretToken, ""),
	)

	_, err = store.ListBuckets()
	if err != nil {
		return &store, errors.New("fatal: failed to list buckets from remote. configuration error?")
	}

	return &store, nil
}

func main() {
	store, err := openStorageAPIConnection()
	Assert(err)
	router, err := routes.InitializeRouter(store)
	Assert(err)
	router.Run(":80")
}

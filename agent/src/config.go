package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"enix.io/banana/src/helpers"
	"github.com/imdario/mergo"
)

// Config : Contains data such as credentials that will be used to execute commands
type Config struct {
	BucketName  string              `json:"bucket"`
	StorageHost string              `json:"storage_host"`
	Vault       helpers.VaultConfig `json:"vault"`
}

// CliConfig : Extended config struct for stuff that can be passed from cli only
type CliConfig struct {
	Config
	Backend string
}

// LoadDefaults : Prepare some default values in configuration
func (config *Config) LoadDefaults() {
	*config = Config{
		BucketName:  "backup-bucket",
		StorageHost: "object-storage.example.com",
		Vault: helpers.VaultConfig{
			Addr:       "http://localhost:7777",
			Token:      "myroot",
			SecretPath: "storage_access",
		},
	}
}

// LoadFromFile : Load configuration from given filename
func (config *Config) LoadFromFile(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("warning: can't load config file " + path + ", using config from env and command-line only")
		return err
	}

	json.Unmarshal(bytes, config)
	return nil
}

// LoadFromArgs : Load configuration from parsed command line arguments
func (config *Config) LoadFromArgs(args *CliConfig) error {
	return mergo.Merge(config, args.Config, mergo.WithOverride)
}

// LoadFromEnv : Load configuration from env variables
func (config *Config) LoadFromEnv() error {
	env := Config{
		BucketName:  os.Getenv("BANANA_BUCKET_NAME"),
		StorageHost: os.Getenv("BANANA_STORAGE_HOST"),
		Vault: helpers.VaultConfig{
			Addr:       os.Getenv("VAULT_ADDR"),
			Token:      os.Getenv("VAULT_TOKEN"),
			SecretPath: os.Getenv("BANANA_VAULT_SECRET_PATH"),
		},
	}

	return mergo.Merge(config, env, mergo.WithOverride)
}

// GetEndpoint : Returns the storage endpoint based on host, bucket and backup name
func (config *Config) GetEndpoint(backupName string) string {
	return fmt.Sprintf("s3://%s/%s/%s", config.StorageHost, config.BucketName, backupName)
}

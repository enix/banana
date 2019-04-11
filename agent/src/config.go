package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/imdario/mergo"
)

// Config : Contains data such as credentials that will be used to execute commands
type Config struct {
	BucketName  string      `json:"bucket"`
	StorageHost string      `json:"storage_host"`
	Vault       VaultConfig `json:"vault"`
}

// CliConfig : Extended config struct for stuff that can be passed from cli only
type CliConfig struct {
	Config
	Backend string
}

// VaultConfig : Configuration for vault API access
type VaultConfig struct {
	Addr       string `json:"address"`
	Token      string `json:"token"`
	SecretPath string `json:"secret_path"`
}

// LoadConfigDefaults : Prepare some default values in configuration
func LoadConfigDefaults(config *Config) {
	*config = Config{
		BucketName:  "backup-bucket",
		StorageHost: "object-storage.example.com",
		Vault: VaultConfig{
			Addr:       "http://localhost:7777",
			Token:      "myroot",
			SecretPath: "storage_access",
		},
	}
}

// LoadConfigFromFile : Load configuration from given filename
func LoadConfigFromFile(config *Config, path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("warning: can't load config file " + path + ", using config from env and command-line only")
		return err
	}

	json.Unmarshal(bytes, config)
	return nil
}

// LoadConfigFromArgs : Load configuration from parsed command line arguments
func LoadConfigFromArgs(config *Config, args *CliConfig) error {
	return mergo.Merge(config, args.Config, mergo.WithOverride)
}

// LoadConfigFromEnv : Load configuration from env variables
func LoadConfigFromEnv(config *Config) error {
	env := Config{
		BucketName:  os.Getenv("BANANA_BUCKET_NAME"),
		StorageHost: os.Getenv("BANANA_STORAGE_HOST"),
		Vault: VaultConfig{
			Addr:       os.Getenv("VAULT_ADDR"),
			Token:      os.Getenv("VAULT_TOKEN"),
			SecretPath: os.Getenv("BANANA_VAULT_SECRET_PATH"),
		},
	}

	return mergo.Merge(config, env, mergo.WithOverride)
}

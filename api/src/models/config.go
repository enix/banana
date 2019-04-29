package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"enix.io/banana/src/services"
	"github.com/imdario/mergo"
)

// Config : Contains full confugration will be used to execute commands
type Config struct {
	MonitorURL  string               `json:"monitor_url"`
	Backend     string               `json:"backend"`
	PrivKeyPath string               `json:"private_key_path"`
	CertPath    string               `json:"client_cert_path"`
	CaCertPath  string               `json:"ca_cert_path"`
	BucketName  string               `json:"bucket"`
	StorageHost string               `json:"storage_host"`
	OpaqueID    string               `json:"opaque_id"`
	TTL         int64                `json:"ttl"`
	Vault       services.VaultConfig `json:"vault"`
}

// CliConfig : Extended config struct for stuff that can be passed from cli only
type CliConfig struct {
	Config
}

// LoadDefaults : Prepare some default values in configuration
func (config *Config) LoadDefaults() {
	*config = Config{
		MonitorURL:  "https://api.banana.enix.io",
		Backend:     "duplicity",
		PrivKeyPath: "/etc/banana/privkey.pem",
		CertPath:    "/etc/banana/cert.pem",
		CaCertPath:  "/etc/banana/cacert.pem",
		BucketName:  "backup-bucket",
		StorageHost: "object-storage.r1.nxs.enix.io",
		TTL:         3600,
		Vault: services.VaultConfig{
			Addr:       "http://localhost:7777",
			Token:      "myroot",
			SecretPath: "banana",
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
		MonitorURL:  os.Getenv("BANANA_MONITOR_URL"),
		Backend:     os.Getenv("BANANA_BACKEND"),
		PrivKeyPath: os.Getenv("BANANA_PRIVATE_KEY_PATH"),
		CertPath:    os.Getenv("BANANA_CLIENT_CERT_PATH"),
		CaCertPath:  os.Getenv("BANANA_CA_CERT_PATH"),
		BucketName:  os.Getenv("BANANA_BUCKET_NAME"),
		StorageHost: os.Getenv("BANANA_STORAGE_HOST"),
		Vault: services.VaultConfig{
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

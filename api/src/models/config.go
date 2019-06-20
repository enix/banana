package models

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"enix.io/banana/src/services"
	"github.com/imdario/mergo"
	"k8s.io/klog"
)

// ScheduledBackupConfig : Contains informations of when to run inc/full backups, and config overrides
type ScheduledBackupConfig struct {
	Config
	Target    string  `json:"target,omitempty"`
	Interval  float32 `json:"interval,omitempty"`
	FullEvery int     `json:"full_every,omitempty"`
}

// Config : Contains full confugration will be used to execute commands
type Config struct {
	MonitorURL         string                           `json:"monitor_url,omitempty"`
	Backend            string                           `json:"backend,omitempty"`
	StatePath          string                           `json:"state_path,omitempty"`
	PrivKeyPath        string                           `json:"private_key_path,omitempty"`
	CertPath           string                           `json:"client_cert_path,omitempty"`
	ScheduleConfigPath string                           `json:"schedule_config_path,omitempty"`
	BucketName         string                           `json:"bucket,omitempty"`
	StorageHost        string                           `json:"storage_host,omitempty"`
	SkipTLSVerify      bool                             `json:"skip_tls_verify,omitempty"`
	TTL                int64                            `json:"ttl,omitempty"`
	Vault              services.VaultConfig             `json:"vault,omitempty"`
	ScheduledBackups   map[string]ScheduledBackupConfig `json:"schedule,omitempty"`
}

// CliConfig : Extended config struct for stuff that can be passed from cli only
type CliConfig struct {
	Config
}

// LoadDefaults : Prepare some default values in configuration
func (config *Config) LoadDefaults() {
	*config = Config{
		MonitorURL:         "https://api.banana.enix.io",
		Backend:            "duplicity",
		StatePath:          "/etc/banana/state.json",
		PrivKeyPath:        "/etc/banana/privkey.pem",
		CertPath:           "/etc/banana/cert.pem",
		ScheduleConfigPath: "/etc/banana/schedule.json",
		BucketName:         "backup-bucket",
		StorageHost:        "object-storage.r1.nxs.enix.io",
		TTL:                3600 * 24 * 30 * 6,
		Vault: services.VaultConfig{
			Addr:              "https://vault.banana.enix.io:7777",
			StorageSecretPath: "openstack",
			RootPath:          "banana",
		},
	}
}

// LoadFromFile : Load configuration from given filename
func (config *Config) LoadFromFile(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		klog.V(2).Info("warning: can't load config file " + path + ", using config from env and command-line only")
		klog.V(2).Info(err)
		return err
	}
	json.Unmarshal(bytes, config)

	if len(config.ScheduleConfigPath) > 0 {
		bytes, err := ioutil.ReadFile(config.ScheduleConfigPath)
		if err != nil {
			klog.V(2).Info("warning: can't load schedule config file " + config.ScheduleConfigPath)
			klog.V(2).Info(err)
			return err
		}

		return json.Unmarshal(bytes, &config.ScheduledBackups)
	}

	return nil
}

// LoadFromArgs : Load configuration from parsed command line arguments
func (config *Config) LoadFromArgs(args *CliConfig) error {
	return mergo.Merge(config, args.Config, mergo.WithOverride)
}

// LoadFromEnv : Load configuration from env variables
func (config *Config) LoadFromEnv() error {
	env := Config{
		MonitorURL:         os.Getenv("BANANA_MONITOR_URL"),
		Backend:            os.Getenv("BANANA_BACKEND"),
		StatePath:          os.Getenv("BANANA_STATE_PATH"),
		PrivKeyPath:        os.Getenv("BANANA_PRIVATE_KEY_PATH"),
		CertPath:           os.Getenv("BANANA_CLIENT_CERT_PATH"),
		ScheduleConfigPath: os.Getenv("BANANA_SCHEDULE_CONFIG_PATH"),
		BucketName:         os.Getenv("BANANA_BUCKET_NAME"),
		StorageHost:        os.Getenv("BANANA_STORAGE_HOST"),
		Vault: services.VaultConfig{
			Addr:              os.Getenv("VAULT_ADDR"),
			RootPath:          os.Getenv("BANANA_VAULT_ROOT_PATH"),
			StorageSecretPath: os.Getenv("BANANA_VAULT_STORAGE_SECRET_PATH"),
		},
	}

	return mergo.Merge(config, env, mergo.WithOverride)
}

// GetEndpoint : Returns the storage endpoint based on host, bucket and backup name
func (config *Config) GetEndpoint(backupName string) string {
	return fmt.Sprintf("s3://%s/%s/%s", config.StorageHost, config.BucketName, backupName)
}

// VerifySignature : Verify that the signature match the struct content
func (config *Config) VerifySignature(cert, sig string) error {
	rawConfig, _ := json.Marshal(config)
	return services.VerifySha256Signature(rawConfig, sig, cert)
}

// Sign : Marshal the struct and generate signature from the result
func (config *Config) Sign(privkey *rsa.PrivateKey) (string, error) {
	rawConfig, _ := json.Marshal(config)
	hash := sha256.New()
	hash.Write(rawConfig)
	digest := hash.Sum(nil)

	sig, err := rsa.SignPKCS1v15(rand.Reader, privkey, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	base64sig := make([]byte, base64.StdEncoding.EncodedLen(len(sig)))
	base64.StdEncoding.Encode(base64sig, sig)
	return string(base64sig), err
}

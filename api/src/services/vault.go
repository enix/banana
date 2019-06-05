package services

import (
	"errors"
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

// Vault : Use this API to interact with Vault
var Vault *VaultClient

// VaultClient : Convenience wrapper of vault client
type VaultClient struct {
	Client        *vault.Client
	Path          string
	StorageAccess *map[string]string
}

// VaultConfig : Configuration for vault API access
type VaultConfig struct {
	Addr       string `json:"address"`
	Token      string `json:"token"`
	SecretPath string `json:"secret_path"`
}

// FetchSecret : Retreive a secret map from the current set path
func (vault *VaultClient) FetchSecret(key string) (map[string]string, error) {
	secret, err := vault.Client.Logical().Read("secret/data/" + key)
	if err != nil {
		return nil, err
	}
	if secret == nil {
		return nil, fmt.Errorf("secret %s not found in vault", key)
	}

	kvInterface, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("vault API returned an unexpected data type")
	}

	kv := make(map[string]string)
	for key, value := range kvInterface {
		kv[key] = fmt.Sprintf("%v", value)
	}

	return kv, nil
}

// GetStorageAccess : Put storage credentials in memory cache
//										and returns the given key from the secret map
func (vault *VaultClient) GetStorageAccess(key string) (string, error) {
	if vault.StorageAccess == nil {
		kv, err := vault.FetchSecret(vault.Path)
		if err != nil {
			return "", err
		}

		vault.StorageAccess = &kv
	}

	val := (*vault.StorageAccess)[key]
	if len(val) == 0 {
		return "", fmt.Errorf("key %s not found within storage access secret", key)
	}

	return val, nil
}

// GetStorageAccessToken : Get storage access token, from memory cache if possible
func (vault *VaultClient) GetStorageAccessToken() (string, error) {
	return vault.GetStorageAccess("AWS_ACCESS_KEY_ID")
}

// GetStorageSecretToken : Get storage secret token, from memory cache if possible
func (vault *VaultClient) GetStorageSecretToken() (string, error) {
	return vault.GetStorageAccess("AWS_SECRET_ACCESS_KEY")
}

// GetStoragePassphrase : Get storage secret token, from memory cache if possible
func (vault *VaultClient) GetStoragePassphrase() (string, error) {
	return vault.GetStorageAccess("PASSPHRASE")
}

// newVaultClient : Create and authenticate a vault client
func newVaultClient(config *VaultConfig, skipTLSVerify bool) (*VaultClient, error) {
	if config.Addr == "" || config.Token == "" {
		return nil, errors.New("missing vault address or token")
	}

	client, err := vault.NewClient(&vault.Config{
		Address:    config.Addr,
		HttpClient: GetHTTPClient(skipTLSVerify),
	})
	if err != nil {
		return nil, err
	}

	client.SetToken(config.Token)
	return &VaultClient{
		Client: client,
		Path:   config.SecretPath,
	}, nil
}

// OpenVaultConnection : Etablish connection with vault
// Other calls will crash if used before this
func OpenVaultConnection(config *VaultConfig, skipTLSVerify bool) error {
	var err error
	Vault, err = newVaultClient(config, skipTLSVerify)
	return err
}

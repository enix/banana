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
	Config        *VaultConfig
	StorageAccess *map[string]string
	EntityID      string
}

// VaultConfig : Configuration for vault API access
type VaultConfig struct {
	Addr              string `json:"address"`
	StorageSecretPath string `json:"storage_secret_path"`
	RootPath          string `json:"root_path"`
}

// FetchSecret : Retreive a secret map from the current set path
func (vault *VaultClient) FetchSecret(key string) (map[string]string, error) {
	company := Credentials.Cert.Subject.Organization[0]
	secret, err := vault.Client.Logical().Read(fmt.Sprintf("%s/%s/secrets/data/%s", vault.Config.RootPath, company, key))
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

// WriteSecret : Write secret map to the current set path
func (vault *VaultClient) WriteSecret(key string, data map[string]interface{}) error {
	company := Credentials.Cert.Subject.Organization[0]
	_, err := vault.Client.Logical().Write(
		fmt.Sprintf("%s/%s/secrets/data/%s", vault.Config.RootPath, company, key),
		map[string]interface{}{
			"data": data,
		},
	)

	return err
}

// GetStorageAccess : Put storage credentials in memory cache
//										and returns the given key from the secret map
func (vault *VaultClient) GetStorageAccess(key string) (string, error) {
	if vault.StorageAccess == nil {
		kv, err := vault.FetchSecret("backends/" + vault.Config.StorageSecretPath)
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

// newVaultClient : Create and authenticate a vault client
func newVaultClient(config *VaultConfig, skipTLSVerify bool) (*VaultClient, error) {
	if config.Addr == "" {
		return nil, errors.New("missing vault address or token")
	}

	client, err := vault.NewClient(&vault.Config{
		Address:    config.Addr,
		HttpClient: GetHTTPClient(skipTLSVerify),
	})
	if err != nil {
		return nil, err
	}

	return &VaultClient{
		Client: client,
		Config: config,
	}, nil
}

// OpenVaultConnection : Etablish connection with vault
// Other calls will crash if used before this
func OpenVaultConnection(config *VaultConfig, skipTLSVerify bool) error {
	var err error
	Vault, err = newVaultClient(config, skipTLSVerify)
	if err != nil {
		return err
	}
	if Credentials != nil {
		secret, err := Vault.Client.Logical().Write("auth/banana/cert/login", map[string]interface{}{})
		if err != nil {
			return err
		}
		Vault.Client.SetToken(secret.Auth.ClientToken)
		self, err := Vault.Client.Logical().Read("auth/token/lookup-self")
		if err != nil {
			return err
		}
		entityID := self.Data["entity_id"]
		Vault.EntityID, _ = entityID.(string)
	}
	return nil
}

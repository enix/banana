package main

import (
	"errors"
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

// Vault : Convenience wrapper of vault client
type Vault struct {
	Client        *vault.Client
	Path          string
	StorageAccess *map[string]string
}

// NewVaultClient : Create and authenticate a vault client
func NewVaultClient(config *VaultConfig) (*Vault, error) {
	client, err := vault.NewClient(&vault.Config{Address: config.Addr})
	if err != nil {
		return nil, err
	}

	client.SetToken(config.Token)
	return &Vault{
		Client: client,
		Path:   config.SecretPath,
	}, nil
}

// FetchSecret : Retreive a secret map from the current set path
func (vault *Vault) FetchSecret(key string) (map[string]string, error) {
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
		kv[fmt.Sprintf("%v", key)] = fmt.Sprintf("%v", value)
	}

	return kv, nil
}

// CacheStorageAccess : Put storage credentials in memory cache
func (vault *Vault) CacheStorageAccess() error {
	if vault.StorageAccess == nil {
		kv, err := vault.FetchSecret(vault.Path)
		if err != nil {
			return err
		}

		vault.StorageAccess = &kv
	}

	return nil
}

// GetStorageAccessToken : Get storage access token, from memory cache if possible
func (vault *Vault) GetStorageAccessToken() (string, error) {
	err := vault.CacheStorageAccess()
	if err != nil {
		return "", err
	}
	return (*vault.StorageAccess)["AWS_ACCESS_KEY_ID"], nil
}

// GetStorageSecretToken : Get storage access token, from memory cache if possible
func (vault *Vault) GetStorageSecretToken() (string, error) {
	err := vault.CacheStorageAccess()
	if err != nil {
		return "", err
	}
	return (*vault.StorageAccess)["AWS_SECRET_ACCESS_KEY"], nil
}

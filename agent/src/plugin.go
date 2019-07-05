package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
)

// plugin : Interface for communicatin with plugins such as duplicity, rsync, tar...
type plugin struct {
	name string
}

// newPlugin : Instanciate the corresponding plugin from its name
func newPlugin(name string) (*plugin, error) {
	if len(name) == 0 {
		return nil, errors.New(name + ": please specify plugin using the --plugin (-p) command line argument")
	}

	return &plugin{name}, nil
}

func (p *plugin) spawn(config *models.Config, args ...string) ([]byte, []byte, error) {
	return execute(config.PluginsDir+"/"+p.name, args...)
}

func (p *plugin) version(config *models.Config) (string, error) {
	stdout, _, err := p.spawn(config, "version")
	return string(stdout), err
}

func (p *plugin) backup(config *models.Config, cmd *backupCmd) ([]byte, error) {
	err := loadCredentialsToEnv()
	if err != nil {
		return nil, err
	}

	args := []string{"backup", cmd.Type, config.GetEndpoint(cmd.Name)}
	args = append(args, cmd.PluginArgs...)

	_, stderr, err := p.spawn(config, args...)
	return stderr, err
}

func (p *plugin) restore(config *models.Config, cmd *restoreCmd) ([]byte, error) {
	err := loadCredentialsToEnv()
	if err != nil {
		return nil, err
	}

	args := []string{"restore", cmd.TargetTime, config.GetEndpoint(cmd.Name)}
	args = append(args, cmd.PluginArgs...)

	_, stderr, err := p.spawn(config, args...)
	return stderr, err
}

func generateAndUploadPassphrase() (string, error) {
	rand.Seed(int64(time.Now().Unix()))
	passphrase := fmt.Sprintf("%d", rand.Int())
	return passphrase, services.Vault.WriteSecret("agents/"+services.Vault.EntityID, map[string]interface{}{
		"passphrase": passphrase,
	})
}

func loadCredentialsToEnv() error {
	accessToken, err := services.Vault.GetStorageAccessToken()
	if err != nil {
		return err
	}
	secretToken, err := services.Vault.GetStorageSecretToken()
	if err != nil {
		return err
	}
	agentSecret, err := services.Vault.FetchSecret("agents/" + services.Vault.EntityID)
	passphrase := ""
	if err != nil {
		passphrase, err = generateAndUploadPassphrase()
		if err != nil {
			return err
		}
	} else {
		passphrase = agentSecret["passphrase"]
	}

	os.Setenv("AWS_ACCESS_KEY_ID", accessToken)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretToken)
	os.Setenv("PASSPHRASE", passphrase)
	return nil
}

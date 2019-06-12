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

// duplicityBackend : BackupBackend implementation for duplicity
type duplicityBackend struct{}

func checkForWarnings() bool {
	_, err := execute("grep", "-q", "WARNING", "/tmp/backup.log")
	execute("rm", "-f", "/tmp/backup.log")
	return err == nil
}

func getBackupID() string {
	idRegExp := "'s/\\.[^.]*\\.(?:[^.]*\\.)?(?:[^.]*\\.)?([^.a-z]*)\\.(?:vol1\\.)?[^.]*\\.gpg/\\1/g'"
	id, _ := execute("sh", "-c", "grep Writing /tmp/backup.log | tail -n 1 | perl -pe "+idRegExp)

	if len(id) < 1 {
		return ""
	}
	return string(id)[:len(id)-1]
}

func generateAndUploadPassphrase() (string, error) {
	rand.Seed(int64(time.Now().Unix()))
	passphrase := fmt.Sprintf("%d", rand.Int())
	return passphrase, services.Vault.WriteSecret("agents/"+services.Vault.EntityID, map[string]interface{}{
		"passphrase": passphrase,
	})
}

func loadCredentialsToEnv() {
	accessToken, err := services.Vault.GetStorageAccessToken()
	assert(err)
	secretToken, err := services.Vault.GetStorageSecretToken()
	assert(err)
	agentSecret, err := services.Vault.FetchSecret("agents/" + services.Vault.EntityID)
	passphrase := ""
	if err != nil {
		passphrase, err = generateAndUploadPassphrase()
		assert(err)
	} else {
		passphrase = agentSecret["passphrase"]
	}

	os.Setenv("AWS_ACCESS_KEY_ID", accessToken)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretToken)
	os.Setenv("PASSPHRASE", passphrase)
}

// Backup : BackupBackend's Backup call implementation for duplicity
func (d *duplicityBackend) backup(config *models.Config, cmd *backupCmd) ([]byte, error) {
	output, err := execute("duplicity", cmd.Type, "-v8", "--log-file", "/tmp/backup.log", cmd.Target, config.GetEndpoint(cmd.Name))
	cmd.OpaqueID = getBackupID()

	if checkForWarnings() {
		return output, errors.New("backup finished with warnings")
	}

	return output, err
}

// Restore : BackupBackend's Restore call implementation for duplicity
func (d *duplicityBackend) restore(config *models.Config, cmd *restoreCmd) ([]byte, error) {
	output, err := execute("duplicity", "--log-file", "/tmp/backup.log", "--restore-time", cmd.TargetTime, config.GetEndpoint(cmd.Name), cmd.TargetDirectory)

	if checkForWarnings() {
		return output, errors.New("restore finished with warnings")
	}
	return output, err
}

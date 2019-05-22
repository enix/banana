package main

import (
	"errors"

	"enix.io/banana/src/models"
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
	return string(id)[:len(id)-1]
}

// Backup : BackupBackend's Backup call implementation for duplicity
func (d *duplicityBackend) backup(config *models.Config, cmd *backupCmd) ([]byte, error) {
	output, err := execute("duplicity", "full", "-v8", "--log-file", "/tmp/backup.log", cmd.Target, config.GetEndpoint(cmd.Name))
	config.OpaqueID = getBackupID()

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

package main

import (
	"errors"

	"enix.io/banana/src/models"
)

// DuplicityBackend : BackupBackend implementation for duplicity
type DuplicityBackend struct{}

func checkForWarnings() bool {
	_, err := Execute("grep", "-q", "WARNING", "/tmp/backup.log")
	Execute("rm", "-f", "/tmp/backup.log")
	return err == nil
}

func getBackupID() string {
	idRegExp := "'s/\\.[^.]*\\.(?:[^.]*\\.)?(?:[^.]*\\.)?([^.a-z]*)\\.(?:vol1\\.)?[^.]*\\.gpg/\\1/g'"
	id, _ := Execute("sh", "-c", "grep Writing /tmp/backup.log | tail -n 1 | perl -pe "+idRegExp)
	return string(id)[:len(id)-1]
}

// Backup : BackupBackend's Backup call implementation for duplicity
func (d *DuplicityBackend) Backup(config *models.Config, cmd *BackupCmd) ([]byte, error) {
	output, err := Execute("duplicity", "full", "-v8", "--log-file", "/tmp/backup.log", cmd.Target, config.GetEndpoint(cmd.Name))
	config.OpaqueID = getBackupID()

	if checkForWarnings() {
		return output, errors.New("backup finished with warnings")
	}

	return output, err
}

// Restore : BackupBackend's Restore call implementation for duplicity
func (d *DuplicityBackend) Restore(config *models.Config, cmd *RestoreCmd) ([]byte, error) {
	output, err := Execute("duplicity", "--log-file", "/tmp/backup.log", "--restore-time", cmd.TargetTime, config.GetEndpoint(cmd.Name), cmd.TargetDirectory)

	if checkForWarnings() {
		return output, errors.New("restore finished with warnings")
	}
	return output, err
}

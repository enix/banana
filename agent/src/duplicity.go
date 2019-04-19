package main

import "errors"

// DuplicityBackend : BackupBackend implementation for duplicity
type DuplicityBackend struct{}

func checkForWarnings() bool {
	_, err := Execute("grep", "-q", "WARNING", "/tmp/backup.log")
	Execute("rm", "-f", "/tmp/backup.log")
	return err == nil
}

// Backup : BackupBackend's Backup call implementation for duplicity
func (d *DuplicityBackend) Backup(config *Config, cmd *BackupCmd) ([]byte, error) {
	output, err := Execute("duplicity", "--log-file", "/tmp/backup.log", "--full-if-older-than", "1W", cmd.Target, config.GetEndpoint(cmd.Name))

	if checkForWarnings() {
		return output, errors.New("backup finished with warnings")
	}
	return output, err
}

// Restore : BackupBackend's Restore call implementation for duplicity
func (d *DuplicityBackend) Restore(config *Config, cmd *RestoreCmd) ([]byte, error) {
	output, err := Execute("duplicity", "--log-file", "/tmp/backup.log", "--restore-time", cmd.TargetTime, config.GetEndpoint(cmd.Name), cmd.TargetDirectory)

	if checkForWarnings() {
		return output, errors.New("restore finished with warnings")
	}
	return output, err
}

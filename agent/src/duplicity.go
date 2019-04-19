package main

// DuplicityBackend : BackupBackend implementation for duplicity
type DuplicityBackend struct{}

// Backup : BackupBackend's Backup call implementation for duplicity
func (d *DuplicityBackend) Backup(config *Config, cmd *BackupCmd) ([]byte, error) {
	return Execute("duplicity", "--full-if-older-than", "1W", cmd.Target, config.GetEndpoint(cmd.Name))
}

// Restore : BackupBackend's Restore call implementation for duplicity
func (d *DuplicityBackend) Restore(config *Config, cmd *RestoreCmd) ([]byte, error) {
	return Execute("duplicity", "--restore-time", cmd.TargetTime, config.GetEndpoint(cmd.Name), cmd.TargetDirectory)
}

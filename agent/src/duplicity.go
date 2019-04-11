package main

import "fmt"

// DuplicityBackend : BackupBackend implementation for duplicity
type DuplicityBackend struct{}

// Backup : BackupBackend's Backup call implementation for duplicity
func (d *DuplicityBackend) Backup(config *Config, cmd *BackupCmd) error {
	fmt.Println("running duplicity!")
	return nil
}

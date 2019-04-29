package main

import (
	"errors"

	"enix.io/banana/src/models"
)

// BackupBackend : Interface for communicatin with backends
//								 such as duplicity, rsync, tar...
type BackupBackend interface {
	Backup(*models.Config, *BackupCmd) ([]byte, error)
	Restore(*models.Config, *RestoreCmd) ([]byte, error)
}

// NewBackupBackend : Instanciate the corresponding backend from its name
func NewBackupBackend(name string) (BackupBackend, error) {
	if len(name) == 0 {
		return nil, errors.New(name + ": please specify backend using the --backend (-b) command line argument")
	}

	switch name {
	case "duplicity":
		return &DuplicityBackend{}, nil
	default:
		return nil, errors.New(name + ": unknown or unsupported backend")
	}
}

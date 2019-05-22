package main

import (
	"errors"

	"enix.io/banana/src/models"
)

// backupBackend : Interface for communicatin with backends
//								 such as duplicity, rsync, tar...
type backupBackend interface {
	backup(*models.Config, *backupCmd) ([]byte, error)
	restore(*models.Config, *restoreCmd) ([]byte, error)
}

// newBackupBackend : Instanciate the corresponding backend from its name
func newBackupBackend(name string) (backupBackend, error) {
	if len(name) == 0 {
		return nil, errors.New(name + ": please specify backend using the --backend (-b) command line argument")
	}

	switch name {
	case "duplicity":
		return &duplicityBackend{}, nil
	default:
		return nil, errors.New(name + ": unknown or unsupported backend")
	}
}

package main

import (
	"errors"

	"enix.io/banana/src/models"
)

// plugin : Interface for communicatin with plugins such as duplicity, rsync, tar...
type plugin interface {
	backup(*models.Config, *backupCmd) ([]byte, error)
	restore(*models.Config, *restoreCmd) ([]byte, error)
}

// newPlugin : Instanciate the corresponding plugin from its name
func newPlugin(name string) (plugin, error) {
	if len(name) == 0 {
		return nil, errors.New(name + ": please specify plugin using the --plugin (-b) command line argument")
	}

	switch name {
	case "duplicity":
		return &duplicityBackend{}, nil
	default:
		return nil, errors.New(name + ": unknown or unsupported plugin")
	}
}

package main

import (
	"errors"

	"enix.io/banana/src/models"
)

// command : Generic interface for all possible user commands
//					 Filled from command line arguments
type command interface {
	execute(*models.Config) error
	jsonMap() map[string]interface{}
}

// newCommand : Instanciate the corresponding implementation of Command
//							depending on loaded configuration
func newCommand(args *launchArgs) (command, error) {
	switch args.Values[0] {
	case "b":
		fallthrough
	case "backup":
		return newBackupCmd(args)
	case "r":
		fallthrough
	case "restore":
		return newRestoreCmd(args)
	case "routine":
		return newRoutineCmd(args)
	default:
		return nil, errors.New(args.Values[0] + ": no such command")
	}
}

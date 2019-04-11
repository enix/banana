package main

import "errors"

// Command : Generic interface for all possible user commands
//					 Filled from command line arguments
type Command interface {
	Execute(*Config) error
}

// NewCommand : Instanciate the corresponding implementation of Command
//							depending on loaded configuration
func NewCommand(args *LaunchArgs) (Command, error) {
	switch args.Values[0] {
	case "backup":
		return NewBackupCmd(args)
	default:
		return nil, errors.New(args.Values[0] + ": no such command")
	}
}

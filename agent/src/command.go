package main

// Command : Generic interface for all possible user commands
//					 Filled from command line arguments
type Command interface {
	Execute(*Config)
}

// NewCommand : Instanciate the corresponding implementation of Command
//							depending on loaded configuration
//							If both return values are nil, the usage will be displayed
func NewCommand(args *LaunchArgs) (Command, error) {
	return NewBackupCmd(args)
}

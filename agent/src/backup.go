package main

import "fmt"

// BackupCmd : Command implementation for 'backup'
type BackupCmd struct{}

// NewBackupCmd : Creates backup command from command line args
func NewBackupCmd(args *LaunchArgs) (*BackupCmd, error) {
	return &BackupCmd{}, nil
}

// Execute : Start the backup using the config in 'this'
func (cmd *BackupCmd) Execute(config *Config) {
	fmt.Printf("command: %v\n", cmd)
	fmt.Printf("config: %+v\n", config)
}

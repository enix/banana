package main

import (
	"errors"
	"fmt"
)

// BackupCmd : Command implementation for 'backup'
type BackupCmd struct {
	Name   string
	Target string
}

// NewBackupCmd : Creates backup command from command line args
func NewBackupCmd(args *LaunchArgs) (*BackupCmd, error) {
	if len(args.Values) < 2 {
		return nil, errors.New("no backup name specified")
	}
	if len(args.Values) < 3 {
		return nil, errors.New("no target folder specified")
	}

	return &BackupCmd{
		Name:   args.Values[1],
		Target: args.Values[2],
	}, nil
}

// Execute : Start the backup using specified backend
func (cmd *BackupCmd) Execute(config *Config) error {
	backend, err := NewBackupBackend(config.Backend)
	if err != nil {
		return err
	}

	SendMessageToMonitor("backup_start", config)
	fmt.Printf("running %s, see you on the other side\n", config.Backend)
	err = backend.Backup(config, cmd)
	if err != nil {
		return err
	}
	SendMessageToMonitor("backup_done", config)
	fmt.Println("backup done, everything OK")
	return nil
}

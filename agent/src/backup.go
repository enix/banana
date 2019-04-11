package main

import "errors"

// BackupCmd : Command implementation for 'backup'
type BackupCmd struct {
	Backend string
	Name    string
	Target  string
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
		Backend: args.Flags.Backend,
		Name:    args.Values[1],
		Target:  args.Values[2],
	}, nil
}

// Execute : Start the backup using specified backend
func (cmd *BackupCmd) Execute(config *Config) error {
	backend, err := NewBackupBackend(cmd.Backend)
	if err != nil {
		return err
	}

	return backend.Backup(config, cmd)
}

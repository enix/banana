package main

import "errors"

// RestoreCmd : Command implementation for 'backup'
type RestoreCmd struct {
	Name            string
	TargetTime      string
	TargetDirectory string
}

// NewRestoreCmd : Creates backup command from command line args
func NewRestoreCmd(args *LaunchArgs) (*RestoreCmd, error) {
	if len(args.Values) < 2 {
		return nil, errors.New("no backup name specified")
	}
	if len(args.Values) < 3 {
		return nil, errors.New("no target date specified")
	}
	if len(args.Values) < 4 {
		return nil, errors.New("no target folder specified")
	}

	return &RestoreCmd{
		Name:            args.Values[1],
		TargetTime:      args.Values[2],
		TargetDirectory: args.Values[3],
	}, nil
}

// Execute : Start the backup using specified backend
func (cmd *RestoreCmd) Execute(config *Config) error {
	backend, err := NewBackupBackend(config.Backend)
	if err != nil {
		return err
	}

	return backend.Restore(config, cmd)
}

package main

// BackupCmd : Command implementation for 'backup'
type BackupCmd struct {
	Backend string
}

// NewBackupCmd : Creates backup command from command line args
func NewBackupCmd(args *LaunchArgs) (*BackupCmd, error) {
	return &BackupCmd{
		Backend: args.Flags.Backend,
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

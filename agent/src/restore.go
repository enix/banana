package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// RestoreCmd : Command implementation for 'backup'
type RestoreCmd struct {
	Name            string `json:"name"`
	TargetTime      string `json:"target_timestamp"`
	TargetDirectory string `json:"target_directory"`
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

	SendMessageToMonitor("restore_start", config, cmd, "")
	fmt.Printf("running %s, see you on the other side\n", config.Backend)
	logs, err := backend.Restore(config, cmd)
	if logs == nil {
		SendMessageToMonitor("agent_crashed", config, cmd, err.Error())
		return err
	}
	if err != nil {
		SendMessageToMonitor("restore_failed", config, cmd, string(logs))
		return err
	}
	SendMessageToMonitor("restore_done", config, cmd, string(logs))

	return err
}

// JSONMap : Convert struct to an anonymous map with given JSON keys
func (cmd *RestoreCmd) JSONMap() (out map[string]interface{}) {
	raw, _ := json.Marshal(cmd)
	json.Unmarshal(raw, &out)
	return
}

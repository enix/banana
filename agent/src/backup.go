package main

import (
	"encoding/json"
	"errors"

	"enix.io/banana/src/models"
	"k8s.io/klog"
)

// BackupCmd : Command implementation for 'backup'
type BackupCmd struct {
	Name   string `json:"name"`
	Target string `json:"target"`
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
func (cmd *BackupCmd) Execute(config *models.Config) error {
	backend, err := NewBackupBackend(config.Backend)
	if err != nil {
		return err
	}

	SendMessageToMonitor("backup_start", config, cmd, "")
	klog.Infof("running %s, see you on the other side\n", config.Backend)
	logs, err := backend.Backup(config, cmd)
	if logs == nil {
		SendMessageToMonitor("agent_crashed", config, cmd, err.Error())
		return err
	}
	if err != nil {
		SendMessageToMonitor("backup_failed", config, cmd, string(logs))
		return err
	}

	SendMessageToMonitor("backup_done", config, cmd, string(logs))
	klog.Info("backup done, everything OK")
	return nil
}

// JSONMap : Convert struct to an anonymous map with given JSON keys
func (cmd *BackupCmd) JSONMap() (out map[string]interface{}) {
	raw, _ := json.Marshal(cmd)
	json.Unmarshal(raw, &out)
	return
}

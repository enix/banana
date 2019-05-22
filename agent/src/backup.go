package main

import (
	"encoding/json"
	"errors"

	"enix.io/banana/src/models"
	"k8s.io/klog"
)

// backupCmd : Command implementation for 'backup'
type backupCmd struct {
	Name   string `json:"name"`
	Target string `json:"target"`
}

// newBackupCmd : Creates backup command from command line args
func newBackupCmd(args *launchArgs) (*backupCmd, error) {
	if len(args.Values) < 2 {
		return nil, errors.New("no backup name specified")
	}
	if len(args.Values) < 3 {
		return nil, errors.New("no target folder specified")
	}

	return &backupCmd{
		Name:   args.Values[1],
		Target: args.Values[2],
	}, nil
}

// execute : Start the backup using specified backend
func (cmd *backupCmd) execute(config *models.Config) error {
	backend, err := newBackupBackend(config.Backend)
	if err != nil {
		return err
	}

	sendMessageToMonitor("backup_start", config, cmd, "")
	klog.Infof("running %s, see you on the other side\n", config.Backend)
	logs, err := backend.backup(config, cmd)
	if logs == nil {
		sendMessageToMonitor("agent_crashed", config, cmd, err.Error())
		return err
	}
	if err != nil {
		sendMessageToMonitor("backup_failed", config, cmd, string(logs))
		return err
	}

	sendMessageToMonitor("backup_done", config, cmd, string(logs))
	klog.Info("backup done, everything OK")
	return nil
}

// jsonMap : Convert struct to an anonymous map with given JSON keys
func (cmd *backupCmd) jsonMap() (out map[string]interface{}) {
	raw, _ := json.Marshal(cmd)
	json.Unmarshal(raw, &out)
	return
}

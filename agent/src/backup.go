package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"enix.io/banana/src/models"
	"k8s.io/klog"
)

// backupCmd : Command implementation for 'backup'
type backupCmd struct {
	Name     string `json:"name"`
	Target   string `json:"target"`
	Type     string `json:"type"`
	OpaqueID string `json:"opaque_id"`
}

// newBackupCmd : Creates backup command from command line args
func newBackupCmd(args *launchArgs) (*backupCmd, error) {
	if len(args.Values) < 2 {
		return nil, errors.New("no backup type")
	}
	if args.Values[1] != "full" && args.Values[1] != "incremental" {
		return nil, fmt.Errorf("unknown backup type: %s", args.Values[1])
	}
	if len(args.Values) < 3 {
		return nil, errors.New("no backup name specified")
	}
	if len(args.Values) < 4 {
		return nil, errors.New("no target folder specified")
	}

	return &backupCmd{
		Type:   args.Values[1],
		Name:   args.Values[2],
		Target: args.Values[3],
	}, nil
}

// execute : Start the backup using specified plugin
func (cmd *backupCmd) execute(config *models.Config) error {
	plugin, err := newPlugin(config.Plugin)
	if err != nil {
		return err
	}

	sendMessageToMonitor("backup_start", config, cmd, "")
	loadCredentialsToEnv()
	klog.Infof("running %s, see you on the other side\n", config.Plugin)
	logs, err := plugin.backup(config, cmd)
	if logs == nil {
		if err != nil {
			sendMessageToMonitor("agent_crashed", config, cmd, err.Error())
		} else {
			sendMessageToMonitor("agent_crashed", config, cmd, "")
		}
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

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
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	PluginArgs []string `json:"plugin_args"`
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
		return nil, errors.New("no targets specified")
	}

	return &backupCmd{
		Type:       args.Values[1],
		Name:       args.Values[2],
		PluginArgs: args.Values[3:],
	}, nil
}

// execute : Start the backup using specified plugin
func (cmd *backupCmd) execute(config *models.Config) error {
	sendMessageToMonitor("backup_start", config, cmd, nil, "")

	plugin, err := newPlugin(config.Plugin)
	if err != nil {
		return err
	}

	klog.Infof("running %s, see you on the other side\n", config.Plugin)
	rawMetadata, logs, artifactsReader, err := plugin.backup(config, cmd)

	if logs == nil {
		if err != nil {
			sendMessageToMonitor("agent_crashed", config, cmd, nil, err.Error())
		} else {
			sendMessageToMonitor("agent_crashed", config, cmd, nil, "")
		}
		return err
	}
	if err != nil {
		sendMessageToMonitor("backup_failed", config, cmd, nil, string(logs))
		return err
	}

	metadata := &models.BackupMetadata{}
	err = json.Unmarshal(rawMetadata, metadata)
	if err != nil {
		return err
	}

	backupID := sendMessageToMonitor("backup_done", config, cmd, metadata, string(logs))
	klog.Info("backup done, everything OK. uploading artifacts...")

	artifactsPostURL := fmt.Sprintf("%s/agents/artifacts/%s", config.MonitorURL, backupID)
	_, err = fireAPIRequest(config, artifactsPostURL, artifactsReader)
	return err
}

// jsonMap : Convert struct to an anonymous map with given JSON keys
func (cmd *backupCmd) jsonMap() (out map[string]interface{}) {
	raw, _ := json.Marshal(cmd)
	json.Unmarshal(raw, &out)
	return
}

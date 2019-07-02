package main

import (
	"encoding/json"
	"errors"

	"enix.io/banana/src/models"
	"k8s.io/klog"
)

// restoreCmd : Command implementation for 'backup'
type restoreCmd struct {
	Name            string `json:"name"`
	TargetTime      string `json:"target_timestamp"`
	TargetDirectory string `json:"target_directory"`
}

// newRestoreCmd : Creates restore command from command line args
func newRestoreCmd(args *launchArgs) (*restoreCmd, error) {
	if len(args.Values) < 2 {
		return nil, errors.New("no backup name specified")
	}
	if len(args.Values) < 3 {
		return nil, errors.New("no target date specified")
	}
	if len(args.Values) < 4 {
		return nil, errors.New("no target folder specified")
	}

	return &restoreCmd{
		Name:            args.Values[1],
		TargetTime:      args.Values[2],
		TargetDirectory: args.Values[3],
	}, nil
}

// execute : Start the restore using specified plugin
func (cmd *restoreCmd) execute(config *models.Config) error {
	plugin, err := newPlugin(config.Plugin)
	if err != nil {
		return err
	}

	sendMessageToMonitor("restore_start", config, cmd, "")
	loadCredentialsToEnv()
	klog.Infof("running %s, see you on the other side\n", config.Plugin)
	logs, err := plugin.restore(config, cmd)
	if logs == nil {
		sendMessageToMonitor("agent_crashed", config, cmd, err.Error())
		return err
	}
	if err != nil {
		sendMessageToMonitor("restore_failed", config, cmd, string(logs))
		return err
	}

	sendMessageToMonitor("restore_done", config, cmd, string(logs))
	return err
}

// jsonMap : Convert struct to an anonymous map with given JSON keys
func (cmd *restoreCmd) jsonMap() (out map[string]interface{}) {
	raw, _ := json.Marshal(cmd)
	json.Unmarshal(raw, &out)
	return
}

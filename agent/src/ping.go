package main

import (
	"encoding/json"

	"enix.io/banana/src/models"
)

// pingCmd : Command implementation for 'ping'
type pingCmd struct{}

// newPingCmd : Creates ping command from command line args
func newPingCmd(args *launchArgs) (*pingCmd, error) {
	return &pingCmd{}, nil
}

// execute : Start the ping using specified plugin
func (cmd *pingCmd) execute(config *models.Config) error {
	sendMessageToMonitor("ping", config, cmd, "")
	return nil
}

// jsonMap : Convert struct to an anonymous map with given JSON keys
func (cmd *pingCmd) jsonMap() (out map[string]interface{}) {
	raw, _ := json.Marshal(cmd)
	json.Unmarshal(raw, &out)
	return
}

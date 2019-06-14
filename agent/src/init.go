package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
)

// initCmd : Command implementation for 'init'
type initCmd struct {
	Organization string
	Name         string
	Token        string
}

// newInitCmd : Creates init command from command line args
func newInitCmd(args *launchArgs) (*initCmd, error) {
	if len(args.Values) < 4 {
		return nil, errors.New("usage: bananactl init <company name> <agent name> <token>")
	}

	return &initCmd{
		Token:        args.Values[1],
		Organization: args.Values[2],
		Name:         args.Values[3],
	}, nil
}

// execute : Start the init using specified backend
func (cmd *initCmd) execute(config *models.Config) error {
	err := os.Mkdir("/etc/banana", 00755)
	if err != nil && os.IsPermission(err) {
		return err
	}

	services.Vault.Client.SetToken(cmd.Token)
	out, err := services.Vault.Client.Logical().Write(
		fmt.Sprintf("%s/%s/agents-pki/issue/default", config.Vault.RootPath, cmd.Organization),
		map[string]interface{}{
			"common_name": cmd.Name,
		},
	)
	if err != nil {
		return err
	}

	cert, _ := out.Data["certificate"].(string)
	privkey, _ := out.Data["private_key"].(string)
	configRaw, _ := json.MarshalIndent(config, "", "  ")

	err = ioutil.WriteFile(config.CertPath, []byte(cert), 00644)
	assert(err)
	err = ioutil.WriteFile(config.PrivKeyPath, []byte(privkey), 00644)
	assert(err)
	err = ioutil.WriteFile("/etc/banana/banana.json", configRaw, 00644)
	assert(err)
	err = ioutil.WriteFile("/etc/banana/schedule.json", []byte("{}"), 00644)
	assert(err)

	loadCredentialsToMem(config)
	sendMessageToMonitor("initialized", config, cmd, "")
	return nil
}

// jsonMap : Convert struct to an anonymous map with given JSON keys
func (cmd *initCmd) jsonMap() (out map[string]interface{}) {
	raw, _ := json.Marshal(cmd)
	json.Unmarshal(raw, &out)
	return
}

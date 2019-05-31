package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
)

// initCmd : Command implementation for 'init'
type initCmd struct {
	Organization string
	Name         string
}

// newInitCmd : Creates init command from command line args
func newInitCmd(args *launchArgs) (*initCmd, error) {
	if len(args.Values) < 2 {
		return nil, errors.New("usage: bananactl init <company name> <agent name>")
	}

	return &initCmd{
		Organization: args.Values[1],
		Name:         args.Values[2],
	}, nil
}

// execute : Start the init using specified backend
func (cmd *initCmd) execute(config *models.Config) error {
	out, err := services.Vault.Client.Logical().Write("agents-pki/issue/"+cmd.Organization, map[string]interface{}{
		"common_name": cmd.Name,
	})
	if err != nil {
		return err
	}

	cert, _ := out.Data["certificate"].(string)
	privkey, _ := out.Data["private_key"].(string)
	cacert, _ := out.Data["issuing_ca"].(string)
	configRaw, _ := json.Marshal(config)

	os.Mkdir("/etc/banana", 00755)
	err = ioutil.WriteFile(config.CertPath, []byte(cert), 00644)
	assert(err)
	err = ioutil.WriteFile(config.PrivKeyPath, []byte(privkey), 00644)
	assert(err)
	err = ioutil.WriteFile(config.CaCertPath, []byte(cacert), 00644)
	assert(err)
	err = ioutil.WriteFile("/etc/banana/banana.json", configRaw, 00644)
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

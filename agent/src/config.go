package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/imdario/mergo"
)

// Config : Contains data such as credentials that will be used to execute commands
type Config struct {
	BucketName string `json:"bucket"`
	VaultAddr  string `json:"vault_addr"`
	VaultToken string `json:"vault_token"`
}

// LoadConfigDefaults : Prepare some default values in configuration
func LoadConfigDefaults(config *Config) {
	*config = Config{
		BucketName: "backup-bucket",
		VaultAddr:  "http://localhost:7777",
		VaultToken: "myroot",
	}
}

// LoadConfigFromFile : Load configuration from given filename
func LoadConfigFromFile(config *Config, path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("warning: failed to load config file " + path)
		return err
	}

	json.Unmarshal(bytes, config)
	return nil
}

// LoadConfigFromArgs : Load configuration from parsed command line arguments
func LoadConfigFromArgs(config *Config, args *Config) error {
	return mergo.Merge(config, *args, mergo.WithOverride)
}

// LoadConfigFromEnv : Load configuration from env variables
func LoadConfigFromEnv(config *Config) error {
	env := Config{
		BucketName: os.Getenv("BANANA_BUCKET_NAME"),
		VaultAddr:  os.Getenv("VAULT_ADDR"),
		VaultToken: os.Getenv("VAULT_TOKEN"),
	}

	return mergo.Merge(config, env, mergo.WithOverride)
}

// func setIfNotEmpty(dst reflect.Value, value reflect.Value) {
// 	dstPtrType := dst.Type()
// 	dstType := dstPtrType.Elem()
// 	dstValue := reflect.Indirect(dst)
// 	zeroValue := reflect.Zero(dstType)
// 	newValue := reflect.ValueOf(value)

// 	if newValue.Interface() != zeroValue.Interface() {
// 		dstValue.Set(reflect.ValueOf(value))
// 	}
// }

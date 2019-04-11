package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func logFatal(err error) {
	log.Fatal(fmt.Errorf("error: %s", err.Error()))
}

func assert(err error) {
	if err != nil {
		logFatal(err)
	}
}

func main() {
	var config Config

	args := LoadArguments()
	if args.DisplayHelp || len(args.Values) < 1 {
		Usage()
		return
	}

	LoadConfigDefaults(&config)
	LoadConfigFromFile(&config, args.ConfigPath)
	err := LoadConfigFromEnv(&config)
	assert(err)
	err = LoadConfigFromArgs(&config, &args.Flags)
	assert(err)
	cmd, err := NewCommand(args)
	assert(err)

	if cmd == nil || reflect.ValueOf(cmd).IsNil() {
		Usage()
		os.Exit(1)
	}

	cmd.Execute(&config)
}

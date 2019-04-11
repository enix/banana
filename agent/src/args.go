package main

import (
	"os"

	"github.com/pborman/getopt/v2"
)

// LaunchArgs : Filled on launch with parsed command line args
type LaunchArgs struct {
	Flags       CliConfig
	Values      []string
	DisplayHelp bool
	ConfigPath  string
}

// LoadArguments : Load args from os.Args to a LaunchArgs struct
func LoadArguments() *LaunchArgs {
	args := LaunchArgs{
		ConfigPath: "./banana.json",
	}

	getopt.FlagLong(&args.DisplayHelp, "help", 'h', "display this message")
	getopt.FlagLong(&args.ConfigPath, "config", 'c', "path to config file", "banana.json")
	getopt.FlagLong(&args.Flags.Backend, "backend", 'b', "backup engine to be used", "duplicity")
	getopt.FlagLong(&args.Flags.BucketName, "bucket", 'k', "target bucket name", "my-bucket-name")
	getopt.FlagLong(&args.Flags.StorageHost, "storage-host", 's', "storage API host", "object-storage.example.com")
	getopt.FlagLong(&args.Flags.Vault.Addr, "vault-addr", 0, "vault API URL", "http://localhost:7777")
	getopt.FlagLong(&args.Flags.Vault.Token, "vault-token", 0, "vault auth token", "myroot")
	getopt.FlagLong(&args.Flags.Vault.SecretPath, "vault-secret", 0, "vault secret path where credentials are stored", "storage_access")

	opts := getopt.CommandLine
	opts.Parse(os.Args)
	for opts.NArgs() > 0 {
		args.Values = append(args.Values, opts.Arg(0))
		opts.Parse(opts.Args())
	}

	return &args
}

// Usage : Outputs getopt's generated help message
func Usage() {
	getopt.Usage()
}

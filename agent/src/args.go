package main

import (
	"os"

	"github.com/pborman/getopt/v2"
)

// LaunchArgs : Filled on launch with parsed command line args
type LaunchArgs struct {
	Flags       Config
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
	getopt.FlagLong(&args.Flags.BucketName, "bucket", 'b', "target bucket name", "my-bucket-name")
	getopt.FlagLong(&args.Flags.VaultAddr, "vault-addr", 0, "vault api URL", "http://localhost:7777")
	getopt.FlagLong(&args.Flags.VaultToken, "vault-token", 0, "vault auth token", "myroot")

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

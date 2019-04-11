package main

import (
	"os"

	"github.com/pborman/getopt/v2"
)

// LaunchFlags : Filled on launch with command line flags
type LaunchFlags struct {
	DisplayHelp bool
}

// LaunchArgs : Filled on launch with parsed command line args
type LaunchArgs struct {
	Flags  LaunchFlags
	Values []string
}

// LoadArguments : Load args from os.Args to a LaunchArgs struct
func LoadArguments() *LaunchArgs {
	var args LaunchArgs
	getopt.FlagLong(&args.Flags.DisplayHelp, "help", 'h', "display this message")

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

package main

import (
	"os"

	"enix.io/banana/src/models"
	"github.com/pborman/getopt/v2"
)

// launchArgs : Filled on launch with parsed command line args
type launchArgs struct {
	Flags       models.CliConfig
	Values      []string
	DisplayHelp bool
	ConfigPath  string
}

// loadArguments : Load args from os.Args to a LaunchArgs struct
func loadArguments() *launchArgs {
	args := launchArgs{
		ConfigPath: "/etc/banana/banana.json",
	}

	getopt.FlagLong(&args.DisplayHelp, "help", 'h', "display this message")
	getopt.FlagLong(&args.ConfigPath, "config", 'c', "path to config file", "banana.json")
	getopt.FlagLong(&args.Flags.Backend, "backend", 'b', "backup engine to be used", "duplicity")
	getopt.FlagLong(&args.Flags.MonitorURL, "monitor-url", 'm', "monitor API endpoint", "https://api.banana.enix.io")
	getopt.FlagLong(&args.Flags.BucketName, "bucket", 'k', "target bucket name", "my-bucket-name")
	getopt.FlagLong(&args.Flags.TTL, "ttl", 't', "time to live", "3600")
	getopt.FlagLong(&args.Flags.StorageHost, "storage-host", 's', "storage API host", "object-storage.r1.nxs.enix.io")
	getopt.FlagLong(&args.Flags.StatePath, "state", 0, "state location", "/etc/banana/state.json")
	getopt.FlagLong(&args.Flags.PrivKeyPath, "privkey", 0, "private key location", "/etc/banana/privkey.pem")
	getopt.FlagLong(&args.Flags.CertPath, "cert", 0, "client certificate location", "/etc/banana/cert.pem")
	getopt.FlagLong(&args.Flags.CaCertPath, "cacert", 0, "CA certificate location", "/etc/banana/cacert.pem")
	getopt.FlagLong(&args.Flags.Vault.Addr, "vault-addr", 0, "vault API URL", "http://localhost:7777")
	getopt.FlagLong(&args.Flags.Vault.Token, "vault-token", 0, "vault auth token", "myroot")
	getopt.FlagLong(&args.Flags.Vault.SecretPath, "vault-secret", 0, "vault secret path where credentials are stored", "banana")

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
	os.Exit(1)
}

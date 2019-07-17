package main

import (
	"os"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
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
		Flags: models.CliConfig{
			Config: models.Config{
				Vault: &services.VaultConfig{},
			},
		},
	}

	getopt.FlagLong(&args.DisplayHelp, "help", 'h', "display this message")
	getopt.FlagLong(&args.ConfigPath, "config", 'c', "path to config file", "banana.json")
	getopt.FlagLong(&args.Flags.Plugin, "plugin", 'p', "backup plugin to be used", "duplicity")
	getopt.FlagLong(&args.Flags.MonitorURL, "monitor-url", 'm', "monitor API endpoint", "https://api.banana.enix.io")
	getopt.FlagLong(&args.Flags.BucketName, "bucket", 'k', "target bucket name", "my-bucket-name")
	getopt.FlagLong(&args.Flags.TTL, "ttl", 't', "time to live", "3600")
	getopt.FlagLong(&args.Flags.StorageHost, "storage-host", 's', "storage API host", "object-storage.r1.nxs.enix.io")
	getopt.FlagLong(&args.Flags.StatePath, "state", 0, "state location", "/etc/banana/state.json")
	getopt.FlagLong(&args.Flags.PrivKeyPath, "privkey", 0, "private key location", "/etc/banana/privkey.pem")
	getopt.FlagLong(&args.Flags.CertPath, "cert", 0, "client certificate location", "/etc/banana/cert.pem")
	getopt.FlagLong(&args.Flags.PluginsDir, "plugins-dir", 0, "directory to search for plugins", "/etc/banana/plugins.d")
	getopt.FlagLong(&args.Flags.Vault.Addr, "vault-addr", 0, "vault API URL", "http://localhost:7777")
	getopt.FlagLong(&args.Flags.Vault.StorageSecretPath, "vault-storage-secret", 0, "vault secret path where credentials are stored", "storage")
	getopt.FlagLong(&args.Flags.Vault.RootPath, "vault-root-path", 0, "vault root path where all PKIs are mounted", "banana")
	getopt.FlagLong(&args.Flags.SkipTLSVerify, "tls-skip-verify", 0, "ignore tls errors")

	opts := getopt.CommandLine
	opts.Parse(os.Args)
	for opts.NArgs() > 0 {
		nextArg := opts.Arg(0)

		if nextArg == "-" {
			args.Values = append(args.Values, opts.Args()[1:]...)
			break
		}

		args.Values = append(args.Values, nextArg)
		opts.Parse(opts.Args())
	}

	return &args
}

// Usage : Outputs getopt's generated help message
func Usage() {
	getopt.Usage()
	os.Exit(1)
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
Provides the command-line interface for the chainbridge application.

For configuration and CLI commands see the README: https://github.com/ChainSafe/ChainBridge.
*/
package main

import (
	"fmt"
	cfg "github.com/ChainSafe/ChainBridge/cfgBuilder"
	"github.com/ChainSafe/ChainBridge/config"
	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
)

var app = cli.NewApp()

var cliFlags = []cli.Flag{
	config.pathFlag,
	config.outPutFlag,
}

var cfgBuilderCommands = cli.Command{
	Name:        "cfgBuilder",
	Usage:       "build configs",
	Category:    "CONFIG",
	Description: "",
	Subcommands: []cli.Command{
		{
			Action:   cfg.ParseDeployConfig,
			Name:     "generate",
			Usage:    "generate bridge keystore, key type determined by flag",
			Flags:    cliFlags,
			Category: "CONFIG",
			Description: "The generate subcommand is used to generate the bridge keystore.\n" +
				"\tIf no options are specified, a secp256k1 key will be made.",
		},
	},
}

func init() {
	app.Action = run
	app.Copyright = "Copyright 2019 ChainSafe Systems Authors"
	app.Name = "chainbridge"
	app.Usage = "ChainBridge V2"
	app.Author = "ChainSafe Systems 2019"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cfgBuilderCommand,
	}

	app.Flags = append(app.Flags, cliFlags...)
}

func main() {

	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	// Pares first argument for path
	if len(os.Args) < 2 {
		log.Error("Please specify path to config json")
		os.Exit(1)
	}
	path := os.Args[1]
	if path == "" {
		log.Error("must provide path")
		os.Exit(1)
	}

	// Read in the config
	cfg, err := parseDeployConfig(path)
	if err != nil {
		log.Error("failed to parse config", "err", err)
		os.Exit(1)
	}

	// Construct the individual relayer configs
	relayerCfgs, err := CreateRelayerConfigs(cfg)
	if err != nil {
		log.Error("failed to construct relayer configs", "err", err)
		os.Exit(1)
	}

	// Check for output path
	var outPath string
	if len(os.Args) == 3 {
		outPath = os.Args[2]
	}

	// Write all the configs to files
	for i, cfg := range relayerCfgs {
		cfg.ToTOML(filepath.Join(outPath, fmt.Sprintf("config%d.toml", i)))
	}
}

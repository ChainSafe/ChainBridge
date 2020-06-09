// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
Builds the configuration files for the chainbridge application.
*/
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ChainSafe/ChainBridge/cfgBuilder"
	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {
	app.Action = run
	app.Copyright = "Copyright 2019 ChainSafe Systems Authors"
	app.Name = "chainbridge"
	app.Usage = "ChainBridge"
	app.Author = "ChainSafe Systems 2019"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

}

func run(ctx *cli.Context) error {
	// Pares first argument for path
	if ctx.NArg() < 1 {
		log.Error("Please specify path to config json")
		os.Exit(1)
	}
	path := ctx.Args().Get(0)
	if path == "" {
		return errors.New("must provide path")
	}

	// Read in the config
	cfg, err := cfgBuilder.ParseDeployConfig(path)
	if err != nil {
		return fmt.Errorf("failed to parse config, err %s", err)
	}

	// Construct the individual relayer configs
	relayerCfgs, err := cfgBuilder.CreateRelayerConfigs(cfg)
	if err != nil {
		return fmt.Errorf("failed to construct relayer configs, err %s", err)
	}

	// Check for output path
	var outPath string
	if ctx.NArg() == 2 {
		outPath = ctx.Args().Get(1)
	}

	// Write all the configs to files
	for i, cfg := range relayerCfgs {
		cfg.ToJSON(filepath.Join(outPath, fmt.Sprintf("config%d.json", i)))
	}
	return nil
}

func main() {

	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

}

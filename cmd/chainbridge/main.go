// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/ChainSafe/ChainBridgeV2/chains/ethereum"
	"github.com/ChainSafe/ChainBridgeV2/core"
	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

var cliFlags = []cli.Flag{
	ConfigFileFlag,
	VerbosityFlag,
	KeystorePathFlag,
}

var generateFlags = []cli.Flag{
	PrivateKeyFlag,
	PasswordFlag,
	Sr25519Flag,
	Secp256k1Flag,
}

var devFlags = []cli.Flag{
	TestKeyFlag,
}

var accountCommand = cli.Command{
	Name:     "accounts",
	Usage:    "manage bridge keystore",
	Category: "KEYSTORE",
	Description: "The account command is used to manage the bridge keystore.\n" +
		"\tTo generate a new account, key type generated is based on the flag passed: chainbridge account generate\n" +
		"\tTo import a keystore file: chainbridge account import path/to/file\n" +
		"\tTo list keys: chainbridge account list",
	Subcommands: []cli.Command{
		{
			Action:   wrapHandler(handleGenerateCmd),
			Name:     "generate",
			Usage:    "generate bridge keystore, key type determined by flag",
			Flags:    generateFlags,
			Category: "KEYSTORE",
			Description: "The generate subcommand is used to generate the bridge keystore.\n" +
				"\tIf no options are specified, a secp256k1 key will be made.",
		},
		{
			Action:      wrapHandler(handleImportCmd),
			Name:        "import",
			Usage:       "import bridge keystore",
			Category:    "KEYSTORE",
			Description: "The import subcommand is used to import a keystore for the bridge.\n",
		},
		{
			Action:      wrapHandler(handleListCmd),
			Name:        "list",
			Usage:       "list bridge keystore",
			Category:    "KEYSTORE",
			Description: "The list subcommand is used to list all of the bridge keystores.\n",
		},
	},
}

// init initializes CLI
func init() {
	app.Action = run
	app.Copyright = "Copyright 2019 ChainSafe Systems Authors"
	app.Name = "chainbridge"
	app.Usage = "ChainBridge V2"
	app.Author = "ChainSafe Systems 2019"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		accountCommand,
	}

	app.Flags = append(app.Flags, cliFlags...)
	app.Flags = append(app.Flags, devFlags...)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func startLogger(ctx *cli.Context) error {
	logger := log.Root()
	handler := logger.GetHandler()
	var lvl log.Lvl

	if lvlToInt, err := strconv.Atoi(ctx.GlobalString(VerbosityFlag.Name)); err == nil {
		lvl = log.Lvl(lvlToInt)
	} else if lvl, err = log.LvlFromString(ctx.GlobalString(VerbosityFlag.Name)); err != nil {
		return err
	}
	log.Root().SetHandler(log.LvlFilterHandler(lvl, handler))

	return nil
}

func run(ctx *cli.Context) error {
	err := startLogger(ctx)
	if err != nil {
		return err
	}

	log.Info("Starting ChainBridge...")

	cfg, err := getConfig(ctx)
	if err != nil {
		return err
	}

	var ks string
	var insecure bool
	if key := ctx.GlobalString(TestKeyFlag.Name); key != "" {
		ks = key
		insecure = true
	} else {
		ks = cfg.keystorePath
	}

	c := core.NewCore()

	for _, chain := range cfg.Chains {
		var chainconfig core.Chain
		if chain.Type == "ethereum" {
			chainconfig, err = ethereum.InitializeChain(&core.ChainConfig{
				Id:           chain.Id,
				Endpoint:     chain.Endpoint,
				From:         chain.From,
				KeystorePath: ks,
				Insecure:     insecure,
				Opts:         chain.Opts,
			})
		} else if chain.Type == "substrate" {
			chainconfig, err = ethereum.InitializeChain(&core.ChainConfig{
				Id:           chain.Id,
				Endpoint:     chain.Endpoint,
				From:         chain.From,
				KeystorePath: ks,
				Insecure:     insecure,
				Opts:         chain.Opts,
			})
		} else {
			return errors.New("Unrecognized Chain Type")
		}
		if err != nil {
			return err
		}
		c.AddChain(chainconfig)
	}

	c.Start()

	return nil
}

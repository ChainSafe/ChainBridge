// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
Provides the command-line interface for the chainbridge application.

For configuration and CLI commands see the README: https://github.com/ChainSafe/ChainBridge.
*/
package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	"github.com/ChainSafe/ChainBridge/chains/substrate"
	"github.com/ChainSafe/ChainBridge/core"
	msg "github.com/ChainSafe/ChainBridge/message"
	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

var cliFlags = []cli.Flag{
	ConfigFileFlag,
	VerbosityFlag,
	KeystorePathFlag,
	BlockstorePathFlag,
	FreshStartFlag,
}

var generateFlags = []cli.Flag{
	PasswordFlag,
	Sr25519Flag,
	Secp256k1Flag,
}

var devFlags = []cli.Flag{
	TestKeyFlag,
}

var importFlags = []cli.Flag{
	EthereumImportFlag,
	PrivateKeyFlag,
	Sr25519Flag,
	Secp256k1Flag,
	PasswordFlag,
}

var accountCommand = cli.Command{
	Name:     "accounts",
	Usage:    "manage bridge keystore",
	Category: "KEYSTORE",
	Description: "The account command is used to manage the bridge keystore.\n" +
		"\tTo generate a new account, key type generated is based on the flag passed: chainbridge account generate\n" +
		"\tTo import a keystore file: chainbridge account import path/to/file\n" +
		"\tTo import a geth keystore file: chainbridge account import --ethereum path/to/file\n" +
		"\tTo import a keystore file: chainbridge account import --privateKey private_key\n" +
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
			Action:   wrapHandler(handleImportCmd),
			Name:     "import",
			Usage:    "import bridge keystore",
			Flags:    importFlags,
			Category: "KEYSTORE",
			Description: "The import subcommand is used to import a keystore for the bridge.\n" +
				"\tA path to the keystore must be provided\n" +
				"\tUse --ethereum to import an ethereum keystore from external sources such as geth\n" +
				"\tUse --privateKey to create a keystore from a provided private key.",
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

	// Check for test key flag
	var ks string
	var insecure bool
	if key := ctx.GlobalString(TestKeyFlag.Name); key != "" {
		ks = key
		insecure = true
	} else {
		ks = cfg.keystorePath
	}

	// Used to signal core shutdown due to fatal error
	sysErr := make(chan error)
	c := core.NewCore(sysErr)

	for _, chain := range cfg.Chains {
		chainId, err := strconv.Atoi(chain.Id)
		if err != nil {
			return err
		}
		chainConfig := &core.ChainConfig{
			Name:           chain.Name,
			Id:             msg.ChainId(chainId),
			Endpoint:       chain.Endpoint,
			From:           chain.From,
			KeystorePath:   ks,
			Insecure:       insecure,
			BlockstorePath: ctx.GlobalString(BlockstorePathFlag.Name),
			FreshStart:     ctx.GlobalBool(FreshStartFlag.Name),
			Opts:           chain.Opts,
		}
		var newChain core.Chain
		logger := log.Root().New("chain", chainConfig.Name)
		if chain.Type == "ethereum" {
			newChain, err = ethereum.InitializeChain(chainConfig, logger, sysErr)
		} else if chain.Type == "substrate" {
			newChain, err = substrate.InitializeChain(chainConfig, logger, sysErr)
		} else {
			return errors.New("unrecognized Chain Type")
		}

		if err != nil {
			return err
		}
		c.AddChain(newChain)
	}

	c.Start()

	return nil
}

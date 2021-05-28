// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
Provides the command-line interface for the chainbridge application.

For configuration and CLI commands see the README: https://github.com/ChainSafe/ChainBridge.
*/
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"strconv"

	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	"github.com/ChainSafe/ChainBridge/chains/substrate"
	"github.com/ChainSafe/ChainBridge/config"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/metrics/health"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

var cliFlags = []cli.Flag{
	config.ConfigFileFlag,
	config.VerbosityFlag,
	config.KeystorePathFlag,
	config.BlockstorePathFlag,
	config.FreshStartFlag,
	config.LatestBlockFlag,
	config.MetricsFlag,
	config.MetricsPort,
}

var generateFlags = []cli.Flag{
	config.PasswordFlag,
	config.Sr25519Flag,
	config.Secp256k1Flag,
	config.SubkeyNetworkFlag,
}

var devFlags = []cli.Flag{
	config.TestKeyFlag,
}

var importFlags = []cli.Flag{
	config.EthereumImportFlag,
	config.PrivateKeyFlag,
	config.Sr25519Flag,
	config.Secp256k1Flag,
	config.PasswordFlag,
	config.SubkeyNetworkFlag,
}

var accountCommand = cli.Command{
	Name:  "accounts",
	Usage: "manage bridge keystore",
	Description: "The accounts command is used to manage the bridge keystore.\n" +
		"\tTo generate a new account (key type generated is determined on the flag passed in): chainbridge accounts generate\n" +
		"\tTo import a keystore file: chainbridge accounts import path/to/file\n" +
		"\tTo import a geth keystore file: chainbridge accounts import --ethereum path/to/file\n" +
		"\tTo import a private key file: chainbridge accounts import --privateKey private_key\n" +
		"\tTo list keys: chainbridge accounts list",
	Subcommands: []*cli.Command{
		{
			Action: wrapHandler(handleGenerateCmd),
			Name:   "generate",
			Usage:  "generate bridge keystore, key type determined by flag",
			Flags:  generateFlags,
			Description: "The generate subcommand is used to generate the bridge keystore.\n" +
				"\tIf no options are specified, a secp256k1 key will be made.",
		},
		{
			Action: wrapHandler(handleImportCmd),
			Name:   "import",
			Usage:  "import bridge keystore",
			Flags:  importFlags,
			Description: "The import subcommand is used to import a keystore for the bridge.\n" +
				"\tA path to the keystore must be provided\n" +
				"\tUse --ethereum to import an ethereum keystore from external sources such as geth\n" +
				"\tUse --privateKey to create a keystore from a provided private key.",
		},
		{
			Action:      wrapHandler(handleListCmd),
			Name:        "list",
			Usage:       "list bridge keystore",
			Description: "The list subcommand is used to list all of the bridge keystores.\n",
		},
	},
}

var (
	Version = "0.0.1"
)

// init initializes CLI
func init() {
	app.Action = run
	app.Copyright = "Copyright 2019 ChainSafe Systems Authors"
	app.Name = "chainbridge"
	app.Usage = "ChainBridge"
	app.Authors = []*cli.Author{{Name: "ChainSafe Systems 2019"}}
	app.Version = Version
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		&accountCommand,
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

	if lvlToInt, err := strconv.Atoi(ctx.String(config.VerbosityFlag.Name)); err == nil {
		lvl = log.Lvl(lvlToInt)
	} else if lvl, err = log.LvlFromString(ctx.String(config.VerbosityFlag.Name)); err != nil {
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

	cfg, err := config.GetConfig(ctx)
	if err != nil {
		return err
	}

	log.Debug("Config on initialization...", "config", *cfg)

	// Check for test key flag
	var ks string
	var insecure bool
	if key := ctx.String(config.TestKeyFlag.Name); key != "" {
		ks = key
		insecure = true
	} else {
		ks = cfg.KeystorePath
	}

	// Used to signal core shutdown due to fatal error
	sysErr := make(chan error)
	c := core.NewCore(sysErr)

	for _, chain := range cfg.Chains {
		chainId, errr := strconv.Atoi(chain.Id)
		if errr != nil {
			return errr
		}
		chainConfig := &core.ChainConfig{
			Name:           chain.Name,
			Id:             msg.ChainId(chainId),
			Endpoint:       chain.Endpoint,
			From:           chain.From,
			KeystorePath:   ks,
			Insecure:       insecure,
			BlockstorePath: ctx.String(config.BlockstorePathFlag.Name),
			FreshStart:     ctx.Bool(config.FreshStartFlag.Name),
			LatestBlock:    ctx.Bool(config.LatestBlockFlag.Name),
			Opts:           chain.Opts,
		}
		var newChain core.Chain
		var m *metrics.ChainMetrics

		logger := log.Root().New("chain", chainConfig.Name)

		if ctx.Bool(config.MetricsFlag.Name) {
			m = metrics.NewChainMetrics(chain.Name)
		}

		if chain.Type == "ethereum" {
			newChain, err = ethereum.InitializeChain(chainConfig, logger, sysErr, m)
		} else if chain.Type == "substrate" {
			newChain, err = substrate.InitializeChain(chainConfig, logger, sysErr, m)
		} else {
			return errors.New("unrecognized Chain Type")
		}

		if err != nil {
			return err
		}
		c.AddChain(newChain)

	}

	// Start prometheus and health server
	if ctx.Bool(config.MetricsFlag.Name) {
		port := ctx.Int(config.MetricsPort.Name)
		blockTimeoutStr := os.Getenv(config.HealthBlockTimeout)
		blockTimeout := config.DefaultBlockTimeout
		if blockTimeoutStr != "" {
			blockTimeout, err = strconv.ParseInt(blockTimeoutStr, 10, 0)
			if err != nil {
				return err
			}
		}
		h := health.NewHealthServer(port, c.Registry, int(blockTimeout))

		go func() {
			http.Handle("/metrics", promhttp.Handler())
			http.HandleFunc("/health", h.HealthStatus)
			err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
			if errors.Is(err, http.ErrServerClosed) {
				log.Info("Health status server is shutting down", err)
			} else {
				log.Error("Error serving metrics", "err", err)
			}
		}()
	}

	c.Start()

	return nil
}

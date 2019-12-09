package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ChainSafe/ChainBridgeV2/chains/centrifuge"
	"github.com/ChainSafe/ChainBridgeV2/chains/ethereum"
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

var cliFlags = []cli.Flag{
	ConfigFileFlag,
	VerbosityFlag,
	KeystorePathFlag,
}

// init initializes CLI
func init() {
	app.Action = run
	app.Copyright = "Copyright 2019 ChainSafe Systems Authors"
	app.Name = "chainbridge"
	app.Usage = "ChainBridge V2"
	app.Author = "ChainSafe Systems 2019"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{}

	app.Flags = append(app.Flags, cliFlags...)
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

	if lvlToInt, err := strconv.Atoi(ctx.String(VerbosityFlag.Name)); err == nil {
		lvl = log.Lvl(lvlToInt)
	} else if lvl, err = log.LvlFromString(ctx.String(VerbosityFlag.Name)); err != nil {
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
	log.Debug("Loaded config", "config", fmt.Sprintf("%+v", cfg))
	// TODO: parse config for endpoints
	ethEndpoint := ""
	ctfgEndpoint := ""

	eth := ethereum.InitializeChain(msg.EthereumId, &core.ChainConfig{
		Endpoint: ethEndpoint,
		Home:     "",
		Away:     "",
	})

	ctfg := centrifuge.InitializeChain(msg.CentrifugeId, &core.ChainConfig{
		Endpoint: ctfgEndpoint,
		Home:     "",
		Away:     "",
	})

	c := core.NewCore(nil)
	c.AddChain(eth)
	c.AddChain(ctfg)
	c.Start()

	return nil
}

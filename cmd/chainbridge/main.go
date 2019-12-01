package main

import (
	"ChainBridgeV2/chains/centrifuge"
	"ChainBridgeV2/chains/ethereum"
	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"
	"os"

	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

var cliFlags = []cli.Flag{
	VerbosityFlag,
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
		os.Exit(1)
	}
}

func run(ctx *cli.Context) error {
	log.Info("Starting ChainBridge...")

	eth := ethereum.InitializeChain(msg.EthereumId, []byte{}, []byte{})

	ctfg := centrifuge.InitializeChain(msg.CentrifugeId, []byte{}, []byte{})

	c := core.NewCore()
	c.AddChain(eth)
	c.AddChain(ctfg)
	c.Start()

	return nil
}

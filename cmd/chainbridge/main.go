package main

import (
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
	log.Info("Hello, world!")
	return nil
}

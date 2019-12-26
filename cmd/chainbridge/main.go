package main

import (
	"fmt"
	"os"
	"strconv"

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

var accountFlags = []cli.Flag{
	GenerateFlag,
	PasswordFlag,
	ImportFlag,
	ListFlag,
	Secp256k1,
}

var accountCommand = cli.Command{
	Action:   handleAccounts,
	Name:     "accounts",
	Usage:    "manage bridge keystore",
	Flags:    accountFlags,
	Category: "KEYSTORE",
	Description: "The account command is used to manage the bridge keystore.\n" +
		"\tTo generate a new secp256k1 (Ethereum) account: bridge account --generate\n" +
		"\tTo import a keystore file: bridge account --import=path/to/file\n" +
		"\tTo list keys: bridge account --list",
}

// init initializes CLI
func init() {
	app.Action = run
	app.Copyright = "Copyright 2019 ChainSafe Systems Authors"
	app.Name = "chainbridge"
	app.Usage = "ChainBridge V2"
	app.Author = "ChainSafe Systems 2019"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		accountCommand,
	}

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
	log.Debug("Loaded config", "config", fmt.Sprintf("%+v", cfg))
	// TODO: add which key we want to use for each chain to config

	ethA := ethereum.InitializeChain(&core.ChainConfig{
		Id:            msg.EthereumId,
		Endpoint:      cfg.EthereumA.Endpoint,
		Receiver:      cfg.EthereumA.Receiver,
		Emitter:       cfg.EthereumA.Emitter,
		Subscriptions: []string{"Transfer(address,bytes32)"},
	})

	ethB := ethereum.InitializeChain(&core.ChainConfig{
		Id:            msg.CentrifugeId,
		Endpoint:      cfg.EthereumB.Endpoint,
		Receiver:      cfg.EthereumB.Receiver,
		Emitter:       cfg.EthereumB.Emitter,
		Subscriptions: []string{"Transfer(address,bytes32)"},
	})

	//ctfg := centrifuge.InitializeChain(&core.ChainConfig{
	//	Id:       msg.CentrifugeId,
	//	Endpoint: ctfgEndpoint,
	//	Receiver: "",
	//	Emitter:  "",
	//})

	c := core.NewCore(nil)
	c.AddChain(ethA)
	c.AddChain(ethB)
	//c.AddChain(ctfg)
	c.Start()

	return nil
}

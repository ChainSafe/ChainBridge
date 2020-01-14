package main

import (
	"os"
	"strconv"

	"github.com/ChainSafe/ChainBridgeV2/chains/centrifuge"
	"github.com/ChainSafe/ChainBridgeV2/chains/ethereum"
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
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
	Ed25519Flag,
	Sr25519Flag,
	Secp256k1Flag,
}

var accountCommand = cli.Command{
	Action:   handleAccountsCmd,
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

	ks := keystore.NewKeystore(cfg.keystorePath)

	// TODO: Load chains iteratively
	eth := ethereum.InitializeChain(&core.ChainConfig{
		Id:            cfg.Chains[0].Id,
		Endpoint:      cfg.Chains[0].Endpoint,
		Receiver:      cfg.Chains[0].Receiver,
		Emitter:       cfg.Chains[0].Emitter,
		From:          cfg.Chains[0].From,
		Subscriptions: []string{"DepositAsset(address,bytes32)"},
		Keystore:      ks,
	})

	ctfg := centrifuge.InitializeChain(&core.ChainConfig{
		Id:            msg.CentrifugeId,
		Endpoint:      cfg.Chains[1].Endpoint,
		Receiver:      cfg.Chains[1].Receiver,
		Emitter:       cfg.Chains[1].Emitter,
		From:          cfg.Chains[1].From,
		Subscriptions: []string{"nfts"},
		Keystore:      ks,
	})

	c := core.NewCore(nil)
	c.AddChain(eth)
	c.AddChain(ctfg)
	c.Start()

	return nil
}

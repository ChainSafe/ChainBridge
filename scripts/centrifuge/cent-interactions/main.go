// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	app = cli.NewApp()

	setEmitterAddressCmd = cli.Command{
		Name:        "set-emitter",
		Description: "Sets emitter address",
		ArgsUsage:   "<address>",
		Action:      setEmitterAddress,
	}

	getEmitterAddressCmd = cli.Command{
		Name:        "get-emitter",
		Description: "Gets emitter address",
		Action:      getEmitterAddress,
	}

	whitelistChainCmd = cli.Command{
		Name:        "whitelist-chain",
		Description: "Gets emitter address",
		ArgsUsage:   "<chainID>",
		Action:      whitelistChain,
	}

	asssetTxCmd = cli.Command{
		Name:        "asset-tx",
		Description: "Gets emitter address",
		ArgsUsage:   "<destinationID> <to> <tokenID> <metadata>",
		Action:      submitAssetTx,
	}
)

func init() {
	app.Commands = []cli.Command{
		setEmitterAddressCmd,
		getEmitterAddressCmd,
		whitelistChainCmd,
		asssetTxCmd,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"strings"

	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var ganacheCommand = cli.Command{
	Action:   handleGanacheCmd,
	Name:     "ganache",
	Usage:    "deploys test ganache",
	Category: "TESTING",
	Flags:    DeployFlags,
	Description: "\tThe deploy command is used to deploy a ganache instance for testing\n" +
		"\tTo add additional accounts, -testaccount\n" +
		"\tTo only use custom accounts, -reset\n" +
		"\tTo use a mnemonic, -mnemonic",
}

func handleGanacheCmd(ctx *cli.Context) error {
	log.Info("Starting Ganache")

	args := ethereum.NewDeployArgs()

	args.Port = ctx.String(PortFlag.Name)
	args.Amount = ctx.String(AmountFlag.Name)
	args.Mnemonic = ctx.String(MnemFlag.Name)

	if newAccounts := ctx.String(AccountFlag.Name); newAccounts != "" {
		accountList := strings.Split(newAccounts, ",")
		if ctx.Bool(ResetFlag.Name) {
			args.Accounts = accountList
		} else {
			args.Accounts = append(args.Accounts, accountList...)
		}
	}

	err := ethereum.RunGanache(args)
	if err != nil {
		return err
	}

	return nil
}

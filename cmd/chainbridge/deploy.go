// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"encoding/hex"
	"os"
	"os/exec"
	"strings"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
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

// The Private Keys of Alice, Bob, Etc
var BaseAccounts = []string{
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].Encode()),
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.BobKey].Encode()),
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.CharlieKey].Encode()),
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.DaveKey].Encode()),
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.EveKey].Encode()),
}

// deployArgd holds the args to run tests with ganache
type deployArgs struct {
	port     string
	amount   string
	mnemonic string
	accounts []string
}

func newDeployArgs() *deployArgs {
	return &deployArgs{
		accounts: BaseAccounts,
		port:     "8545",
		amount:   "100000000000000000000",
		mnemonic: "",
	}
}

func handleGanacheCmd(ctx *cli.Context) error {
	log.Info("Starting Ganache")

	args := newDeployArgs()

	args.port = ctx.String(PortFlag.Name)
	args.amount = ctx.String(AmountFlag.Name)
	args.mnemonic = ctx.String(MnemFlag.Name)

	if newAccounts := ctx.String(AccountFlag.Name); newAccounts != "" {
		accountList := strings.Split(newAccounts, ",")
		if ctx.Bool(ResetFlag.Name) {
			args.accounts = accountList
		} else {
			args.accounts = append(args.accounts, accountList...)
		}
	}

	err := runGanache(args)
	if err != nil {
		return err
	}

	return nil
}

// Takes the Args struct and converts it to a string array to be passed into RunGanache
func (a *deployArgs) ConvertToStringArray() []string {
	args := []string{}

	args = append(args, "-p", a.port)
	for _, val := range a.accounts {
		input := val + "," + a.amount
		args = append(args, "--account", input)
	}
	if a.mnemonic != "" {
		args = append(args, "-m", a.mnemonic)
	}
	return args
}

// RunGanache takes an input string and the gopath and run ganache-cli with the given inputs
func runGanache(args *deployArgs) error {

	log.Info("Running npm install")
	command := exec.Command("npm", "install")
	command.Dir = "./on-chain/evm-contracts"
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}

	log.Info("Executing Ganache")
	command = exec.Command("./node_modules/.bin/ganache-cli", args.ConvertToStringArray()...) //nolint:gosec
	command.Dir = "./on-chain/evm-contracts"
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()
	if err != nil {
		return err
	}

	return nil

}

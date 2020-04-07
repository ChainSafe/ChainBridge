// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"encoding/hex"
	"os"
	"os/exec"

	"github.com/ChainSafe/ChainBridge/keystore"
	"github.com/ethereum/go-ethereum/log"
)

// deployArgs holds the args to run tests with ganache
type deployArgs struct {
	port     string
	amount   string
	mnemonic string
	accounts []string
}

// The Private Keys of Alice, Bob, Etc
var BaseAccounts = []string{
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].Encode()),
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.BobKey].Encode()),
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.CharlieKey].Encode()),
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.DaveKey].Encode()),
	"0x" + hex.EncodeToString(keystore.TestKeyRing.EthereumKeys[keystore.EveKey].Encode()),
}

func NewDeployArgs() *deployArgs {
	return &deployArgs{
		accounts: BaseAccounts,
		port:     "8545",
		amount:   "100000000000000000000",
		mnemonic: "",
	}
}

// RunGanache takes an input string and the gopath and run ganache-cli with the given inputs
func RunGanache(args *deployArgs) error {

	log.Info("Running npm install")
	command := exec.Command("npm", "install")
	command.Dir = "./solidity"
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}

	log.Info("Executing Ganache")
	command = exec.Command("ganache-cli", args.ConvertToStringArray()...) //nolint:gosec
	command.Dir = "./solidity"
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()
	if err != nil {
		return err
	}

	return nil

}

// ConverToStringArray Takes the Args struct and converts it to a string array to be passed into RunGanache
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

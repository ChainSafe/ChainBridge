package main

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
)

var deployCommand = cli.Command{
	Action:   handleDeployCmd,
	Name:     "testdeploy",
	Usage:    "deploys test ganache",
	Category: "TESTING",
	Flags:    TestFlags,
	Description: "\tThe testdeploy command is used to deploy a ganache instance for testing\n" +
		"\tTo add additional accounts, -testaccount\n" +
		"\tTo only use custom accounts, -reset\n" +
		"\tTo use a mnemonic, -mnemonic",
}

func handleDeployCmd(ctx *cli.Context) error {
	log.Info("Starting Ganache")
	gopath := os.Getenv("GOPATH")
	accounts := []string{"0x000000000000000000000000000000000000000000000000000000416c696365", "0x0000000000000000000000000000000000000000000000000000000000426f62", "0x00000000000000000000000000000000000000000000000000436861726c6965", "0x0000000000000000000000000000000000000000000000000000000044617665", "0x0000000000000000000000000000000000000000000000000000000000457665"}

	args := []string{}

	port := "8545"
	if newPort := ctx.String(PortFlag.Name); newPort != "" {
		port = newPort
	}
	args = append(args, "-p", string(port))

	if newAccounts := ctx.String(AccountFlag.Name); newAccounts != "" {
		accountList := strings.Split(newAccounts, ",")
		if reset := ctx.Bool(ResetFlag.Name); reset {
			accounts = accountList
		} else {
			accounts = append(accounts, accountList...)
		}
	}
	amount := "100000000000000000000"
	if newAmount := ctx.String(AmountFlag.Name); newAmount != "" {
		amount = newAmount
	}

	for _, account := range accounts {
		input := account + "," + amount
		args = append(args, "--account", input)
	}

	if mnem := ctx.String(MnemFlag.Name); mnem != "" {
		args = append(args, "-m", mnem)
	}

	err := runGanache(args, gopath)
	if err != nil {
		return err
	}

	return nil
}

func runGanache(args []string, gopath string) error {

	log.Info("Running npm install")
	command := exec.Command("npm", "install")
	command.Dir = gopath + "/src/github.com/ChainSafe/ChainBridgeV2/on-chain/evm-contracts"
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}
	for _, entry := range args {
		log.Info(entry)
	}
	command = exec.Command("./node_modules/.bin/ganache-cli", args...)
	command.Dir = gopath + "/src/github.com/ChainSafe/ChainBridgeV2/on-chain/evm-contracts"
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()
	if err != nil {
		return err
	}

	return nil

}

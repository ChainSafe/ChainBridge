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

// The Private Keys of Alice, Bob, Etc
var BaseAccounts = []string{"0x000000000000000000000000000000000000000000000000000000416c696365", "0x0000000000000000000000000000000000000000000000000000000000426f62", "0x00000000000000000000000000000000000000000000000000436861726c6965", "0x0000000000000000000000000000000000000000000000000000000044617665", "0x0000000000000000000000000000000000000000000000000000000000457665"}

// Args holds the args to run tests with ganache
type Args struct {
	port     string
	amount   string
	mnemonic string
	accounts []string
}

func newArgs() *Args {
	return &Args{
		accounts: BaseAccounts,
		port:     "8545",
		amount:   "100000000000000000000",
	}
}

func handleDeployCmd(ctx *cli.Context) error {
	log.Info("Starting Ganache")
	gopath := os.Getenv("GOPATH")

	args := newArgs()

	if newPort := ctx.String(PortFlag.Name); newPort != "" {
		args.port = newPort
	}

	if newAccounts := ctx.String(AccountFlag.Name); newAccounts != "" {
		accountList := strings.Split(newAccounts, ",")
		if reset := ctx.Bool(ResetFlag.Name); reset {
			args.accounts = accountList
		} else {
			args.accounts = append(args.accounts, accountList...)
		}
	}

	if newAmount := ctx.String(AmountFlag.Name); newAmount != "" {
		args.amount = newAmount
	}

	if mnem := ctx.String(MnemFlag.Name); mnem != "" {
		args.mnemonic = mnem
	}

	err := RunGanache(args.ConvertToStringArray(), gopath)
	if err != nil {
		return err
	}

	return nil
}

// Takes the Args struct and converts it to a string array to be passed into RunGanache
func (a *Args) ConvertToStringArray() []string {
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
func RunGanache(args []string, gopath string) error {

	log.Info("Running npm install")
	command := exec.Command("npm", "install")
	command.Dir = gopath + "/src/github.com/ChainSafe/ChainBridgeV2/on-chain/evm-contracts"
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}

	command = exec.Command("./node_modules/.bin/ganache-cli", args...) //nolint:gosec
	command.Dir = gopath + "/src/github.com/ChainSafe/ChainBridgeV2/on-chain/evm-contracts"
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Start()
	if err != nil {
		return err
	}

	return nil

}

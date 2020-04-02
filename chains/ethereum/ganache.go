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
	Port     string
	Amount   string
	Mnemonic string
	Accounts []string
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
		Accounts: BaseAccounts,
		Port:     "8545",
		Amount:   "100000000000000000000",
		Mnemonic: "",
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

	args = append(args, "-p", a.Port)
	for _, val := range a.Accounts {
		input := val + "," + a.Amount
		args = append(args, "--account", input)
	}
	if a.Mnemonic != "" {
		args = append(args, "-m", a.Mnemonic)
	}
	return args
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ethereum/go-ethereum/common"
)

//createChainConfig helps create a chain config
func createChainConfig(t *testing.T, bridge string, erc20Handler string, erc721Handler string, genericHandler string) core.ChainConfig {
	return core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"bridge":         bridge,
			"erc20Handler":   erc20Handler,
			"erc721Handler":  erc721Handler,
			"genericHandler": genericHandler,
			"gasLimit":       "10",
			"gasPrice":       "20",
			"http":           "true",
			"startBlock":     "10",
		},
	}
}

//TestParseChainConfig tests parseChainConfig with all handlerContracts provided
func TestParseChainConfig(t *testing.T) {
	input := createChainConfig(t, "0x1234", "0x1234", "0x1234", "0x1234")

	out, err := parseChainConfig(&input)

	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		name:                   "chain",
		id:                     1,
		endpoint:               "endpoint",
		from:                   "0x0",
		keystorePath:           "./keys",
		bridgeContract:         common.HexToAddress("0x1234"),
		erc20HandlerContract:   common.HexToAddress("0x1234"),
		erc721HandlerContract:  common.HexToAddress("0x1234"),
		genericHandlerContract: common.HexToAddress("0x1234"),
		gasLimit:               big.NewInt(10),
		gasPrice:               big.NewInt(20),
		http:                   true,
		startBlock:             big.NewInt(10),
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

//TestChainConfigOneContract Tests chain config providing only one contract
func TestChainConfigOneContract(t *testing.T) {

	input := createChainConfig(t, "0x1234", "0x1234", "", "")

	out, err := parseChainConfig(&input)

	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		name:                 "chain",
		id:                   1,
		endpoint:             "endpoint",
		from:                 "0x0",
		keystorePath:         "./keys",
		bridgeContract:       common.HexToAddress("0x1234"),
		erc20HandlerContract: common.HexToAddress("0x1234"),
		gasLimit:             big.NewInt(10),
		gasPrice:             big.NewInt(20),
		http:                 true,
		startBlock:           big.NewInt(10),
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

func TestRequiredOpts(t *testing.T) {
	// No opts provided
	input := core.ChainConfig{
		Id:           0,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts:         map[string]string{},
	}

	_, err := parseChainConfig(&input)

	if err == nil {
		t.Error("config missing chainId field but no error reported")
	}

	// Empty bridgeContract provided
	input = core.ChainConfig{
		Id:           0,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts:         map[string]string{"bridge": ""},
	}

	_, err2 := parseChainConfig(&input)

	if err2 == nil {
		t.Error("config missing chainId field but no error reported")
	}

}

func TestExtraOpts(t *testing.T) {
	input := core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"bridge":        "0x1234",
			"gasLimit":      "10",
			"gasPrice":      "20",
			"http":          "true",
			"incorrect_opt": "error",
		},
	}

	_, err := parseChainConfig(&input)

	if err == nil {
		t.Error("Config should not accept incorrect opts.")
	}
}

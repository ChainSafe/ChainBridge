// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ethereum/go-ethereum/common"
)

//TestParseChainConfig tests parseChainConfig with all handlerContracts provided
func TestParseChainConfig(t *testing.T) {

	input := core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"bridge":             "0x1234",
			"erc20Handler":       "0x1234",
			"erc721Handler":      "0x1234",
			"genericHandler":     "0x1234",
			"gasLimit":           "10",
			"gasMultiplier":      "1",
			"maxGasPrice":        "20",
			"http":               "true",
			"startBlock":         "10",
			"blockConfirmations": "50",
		},
	}

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
		maxGasPrice:            big.NewInt(20),
		gasMultiplier:          big.NewFloat(1),
		http:                   true,
		startBlock:             big.NewInt(10),
		blockConfirmations:     big.NewInt(50),
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

//TestParseChainConfig tests parseChainConfig with all handlerContracts provided
func TestParseChainConfigWithNoBlockConfirmations(t *testing.T) {

	input := core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"bridge":         "0x1234",
			"erc20Handler":   "0x1234",
			"erc721Handler":  "0x1234",
			"genericHandler": "0x1234",
			"gasLimit":       "10",
			"gasMultiplier":  "1",
			"maxGasPrice":    "20",
			"http":           "true",
			"startBlock":     "10",
		},
	}

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
		maxGasPrice:            big.NewInt(20),
		gasMultiplier:          big.NewFloat(1),
		http:                   true,
		startBlock:             big.NewInt(10),
		blockConfirmations:     big.NewInt(DefaultBlockConfirmations),
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

//TestChainConfigOneContract Tests chain config providing only one contract
func TestChainConfigOneContract(t *testing.T) {

	input := core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"bridge":        "0x1234",
			"erc20Handler":  "0x1234",
			"gasLimit":      "10",
			"maxGasPrice":   "20",
			"gasMultiplier": "1",
			"http":          "true",
			"startBlock":    "10",
		},
	}

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
		maxGasPrice:          big.NewInt(20),
		gasMultiplier:        big.NewFloat(1),
		http:                 true,
		startBlock:           big.NewInt(10),
		blockConfirmations:   big.NewInt(DefaultBlockConfirmations),
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
			"maxGasPrice":   "20",
			"gasMultiplier": "1",
			"http":          "true",
			"incorrect_opt": "error",
		},
	}

	_, err := parseChainConfig(&input)

	if err == nil {
		t.Error("Config should not accept incorrect opts.")
	}
}

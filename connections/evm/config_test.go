// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package evm

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ethereum/go-ethereum/common"
)

func TestParseChainConfig(t *testing.T) {
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
			"gasPrice":       "20",
			"http":           "true",
			"startBlock":     "10",
		},
	}

	out, err := ParseChainConfig(&input)

	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		Name:                   "chain",
		Id:                     1,
		Endpoint:               "endpoint",
		From:                   "0x0",
		KeystorePath:           "./keys",
		BridgeContract:         common.HexToAddress("0x1234"),
		Erc20HandlerContract:   common.HexToAddress("0x1234"),
		Erc721HandlerContract:  common.HexToAddress("0x1234"),
		GenericHandlerContract: common.HexToAddress("0x1234"),
		GasLimit:               big.NewInt(10),
		GasPrice:               big.NewInt(20),
		http:                   true,
		StartBlock:             big.NewInt(10),
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

	_, err := ParseChainConfig(&input)

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

	_, err2 := ParseChainConfig(&input)

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

	_, err := ParseChainConfig(&input)

	if err == nil {
		t.Error("Config should not accept incorrect opts.")
	}
}

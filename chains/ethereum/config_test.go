// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"reflect"
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
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
			"minGasPrice":        "0",
			"http":               "true",
			"startBlock":         "10",
			"blockConfirmations": "50",
			"egsApiKey":          "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // fake key
			"egsSpeed":           "fast",
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
		minGasPrice:            big.NewInt(0),
		gasMultiplier:          big.NewFloat(1),
		http:                   true,
		startBlock:             big.NewInt(10),
		blockConfirmations:     big.NewInt(50),
		egsApiKey:              "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		egsSpeed:               "fast",
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

// TestParseChainConfigHex tests the parseChainConfig with hex values
func TestParseChainConfigHex(t *testing.T) {
	commonHexValue := "0x1234"
	expectedHex, parseErr := utils.ParseUint256OrHex(&commonHexValue)
	if parseErr != nil {
		t.Fatalf("unable to parse value %s, %v", commonHexValue, parseErr)
	}

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
			"gasLimit":           commonHexValue,
			"gasMultiplier":      "1",
			"maxGasPrice":        commonHexValue,
			"minGasPrice":        commonHexValue,
			"http":               "true",
			"startBlock":         "10",
			"blockConfirmations": "50",
			"egsApiKey":          "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // fake key
			"egsSpeed":           "fast",
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
		gasLimit:               expectedHex,
		maxGasPrice:            expectedHex,
		minGasPrice:            expectedHex,
		gasMultiplier:          big.NewFloat(1),
		http:                   true,
		startBlock:             big.NewInt(10),
		blockConfirmations:     big.NewInt(50),
		egsApiKey:              "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		egsSpeed:               "fast",
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
			"minGasPrice":    "0",
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
		minGasPrice:            big.NewInt(0),
		gasMultiplier:          big.NewFloat(1),
		http:                   true,
		startBlock:             big.NewInt(10),
		blockConfirmations:     big.NewInt(DefaultBlockConfirmations),
		egsApiKey:              "",
		egsSpeed:               "fast",
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
			"minGasPrice":   "0",
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
		minGasPrice:          big.NewInt(0),
		gasMultiplier:        big.NewFloat(1),
		http:                 true,
		startBlock:           big.NewInt(10),
		blockConfirmations:   big.NewInt(DefaultBlockConfirmations),
		egsApiKey:            "",
		egsSpeed:             "fast",
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

func TestEthGasStationDefaultSpeed(t *testing.T) {
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
			"minGasPrice":    "0",
			"http":           "true",
			"startBlock":     "10",
			"egsApiKey":      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"egsSpeed":       "",
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
		minGasPrice:            big.NewInt(0),
		gasMultiplier:          big.NewFloat(1),
		http:                   true,
		startBlock:             big.NewInt(10),
		blockConfirmations:     big.NewInt(DefaultBlockConfirmations),
		egsApiKey:              "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		egsSpeed:               "fast",
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

func TestEthGasStationCustomSpeed(t *testing.T) {
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
			"minGasPrice":    "0",
			"http":           "true",
			"startBlock":     "10",
			"egsApiKey":      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"egsSpeed":       "average",
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
		minGasPrice:            big.NewInt(0),
		gasMultiplier:          big.NewFloat(1),
		http:                   true,
		startBlock:             big.NewInt(10),
		blockConfirmations:     big.NewInt(DefaultBlockConfirmations),
		egsApiKey:              "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		egsSpeed:               "average",
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

func TestEthGasStationInvalidSpeed(t *testing.T) {
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
			"minGasPrice":    "0",
			"http":           "true",
			"startBlock":     "10",
			"egsApiKey":      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"egsSpeed":       "standard",
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
		minGasPrice:            big.NewInt(0),
		gasMultiplier:          big.NewFloat(1),
		http:                   true,
		startBlock:             big.NewInt(10),
		blockConfirmations:     big.NewInt(DefaultBlockConfirmations),
		egsApiKey:              "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		egsSpeed:               "fast",
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

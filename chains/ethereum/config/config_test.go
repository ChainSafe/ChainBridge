// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"flag"
	"math/big"
	"testing"

	"github.com/ChainSafe/chainbridge-utils/keystore"

	"github.com/ChainSafe/ChainBridge/config"
	"github.com/ChainSafe/ChainBridge/config/flags"
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/common"
)

var bridgeContract = common.HexToAddress("0x62877dDCd49aD22f5eDfc6ac108e9a4b5D2bD88B")
var erc20HandlerContract = common.HexToAddress("0x3167776db165D8eA0f51790CA2bbf44Db5105ADF")
var erc721HandlerContract = common.HexToAddress("0x3f709398808af36ADBA86ACC617FeB7F5B7B193E")
var genericHandlerContract = common.HexToAddress("0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07")
var aliceAddr = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].CommonAddress().String()

func defaultFlagSet() *cli.Context {
	set := flag.NewFlagSet("test", 0)

	set.String(flags.ConfigFileFlag.Name, flags.ConfigFileFlag.Value, "")
	set.String(flags.KeystorePathFlag.Name, flags.KeystorePathFlag.Value, "")
	set.String(flags.BlockstorePathFlag.Name, flags.BlockstorePathFlag.Value, "")
	set.Bool(flags.FreshStartFlag.Name, false, "")
	set.Bool(flags.LatestBlockFlag.Name, false, "")

	return cli.NewContext(nil, set, nil)
}

func TestParseChainConfig_DefaultOpts(t *testing.T) {
	input := config.RawChainConfig{
		Name:     "chain",
		Id:       "1",
		Endpoint: "endpoint",
		From:     aliceAddr,
		Opts: []byte(
			`{` +
				`"bridge": "0x62877ddcd49ad22f5edfc6ac108e9a4b5d2bd88b", ` +
				`"erc20Handler": "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF", ` +
				`"erc721Handler": "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E", ` +
				`"genericHandler": "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07" ` +
				`}`),
	}

	out, err := ParseChainConfig(&input, defaultFlagSet())

	if err != nil {
		t.Fatal(err)
	}

	expected := &Config{
		Name:                   "chain",
		Id:                     1,
		Endpoint:               "endpoint",
		From:                   aliceAddr,
		KeystorePath:           "./keys",
		BridgeContract:         bridgeContract,
		ERC20HandlerContract:   erc20HandlerContract,
		ERC721HandlerContract:  erc721HandlerContract,
		GenericHandlerContract: genericHandlerContract,
		GasLimit:               big.NewInt(DefaultGasLimit),
		MaxGasPrice:            big.NewInt(DefaultMaxGasPrice),
		GasMultiplier:          big.NewFloat(DefaultGasMultiplier),
		Http:                   false,
		StartBlock:             big.NewInt(0),
		LatestBlock:            false,
		BlockConfirmations:     big.NewInt(DefaultBlockConfirmations),
	}

	compareConfigs(t, expected, out)
}

func TestParseChainConfig_NonDefaultOpts(t *testing.T) {
	input := config.RawChainConfig{
		Name:     "chain",
		Id:       "1",
		Endpoint: "endpoint",
		From:     aliceAddr,
		Opts: []byte("{" +
			`"bridge": "0x62877ddcd49ad22f5edfc6ac108e9a4b5d2bd88b", ` +
			`"erc20Handler": "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF", ` +
			`"erc721Handler": "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E", ` +
			`"genericHandler": "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07", ` +
			`"gasLimit": 1, ` +
			`"maxGasPrice": 2, ` +
			`"gasMultiplier": 3.0, ` +
			`"http": true, ` +
			`"startBlock": 4, ` +
			`"blockConfirmations": 5 ` +
			`}`),
	}

	out, err := ParseChainConfig(&input, defaultFlagSet())
	if err != nil {
		t.Fatal(err)
	}

	expected := &Config{
		Name:                   "chain",
		Id:                     1,
		Endpoint:               "endpoint",
		From:                   aliceAddr,
		KeystorePath:           "./keys",
		BridgeContract:         bridgeContract,
		ERC20HandlerContract:   erc20HandlerContract,
		ERC721HandlerContract:  erc721HandlerContract,
		GenericHandlerContract: genericHandlerContract,
		GasLimit:               big.NewInt(1),
		MaxGasPrice:            big.NewInt(2),
		GasMultiplier:          big.NewFloat(3.0),
		Http:                   true,
		StartBlock:             big.NewInt(4),
		BlockConfirmations:     big.NewInt(5),
	}

	compareConfigs(t, expected, out)
}

func TestParseChainConfig_InvalidAddress(t *testing.T) {
	input := config.RawChainConfig{
		Name:     "chain",
		Id:       "1",
		Endpoint: "endpoint",
		From:     aliceAddr,
		Opts: []byte("{" +
			`"bridge": "0x877ddcd49ad22f5edfc6ac108e9a4b5d2bd88b", ` +
			`"erc20Handler": "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF", ` +
			`"erc721Handler": "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E", ` +
			`"genericHandler": "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07" ` +
			`}`),
	}

	_, err := ParseChainConfig(&input, defaultFlagSet())

	if err == nil {
		t.Fatal("invalid address supplied, no error returned")
	}

	if err.Error() != "hex string has length 38, want 40 for common.Address" {
		t.Fatal("unexpected error")
	}
}

func TestParseChainConfig_MissingRequiredFields(t *testing.T) {
	tests := []struct {
		input    *config.RawChainConfig
		expected string
	}{
		{
			&config.RawChainConfig{
				Name:     "chain",
				Id:       "1",
				Endpoint: "endpoint",
				From:     aliceAddr,
				Opts: []byte("{" +
					`"erc20Handler": "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF", ` +
					`"erc721Handler": "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E", ` +
					`"genericHandler": "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07" ` +
					`}`),
			},
			requiredOptsError("bridge").Error(),
		}, {
			&config.RawChainConfig{
				Name:     "chain",
				Id:       "1",
				Endpoint: "endpoint",
				From:     aliceAddr,
				Opts: []byte("{" +
					`"bridge": "0x62877dDCd49aD22f5eDfc6ac108e9a4b5D2bD88B", ` +
					`"erc721Handler": "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E", ` +
					`"genericHandler": "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07" ` +
					`}`),
			},
			requiredOptsError("erc20Handler").Error(),
		}, {
			&config.RawChainConfig{
				Name:     "chain",
				Id:       "1",
				Endpoint: "endpoint",
				From:     aliceAddr,
				Opts: []byte("{" +
					`"bridge": "0x62877dDCd49aD22f5eDfc6ac108e9a4b5D2bD88B", ` +
					`"erc20Handler": "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF", ` +
					`"genericHandler": "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07" ` +
					`}`),
			},
			requiredOptsError("erc721Handler").Error(),
		}, {
			&config.RawChainConfig{
				Name:     "chain",
				Id:       "1",
				Endpoint: "endpoint",
				From:     aliceAddr,
				Opts: []byte("{" +
					`"bridge": "0x62877dDCd49aD22f5eDfc6ac108e9a4b5D2bD88B", ` +
					`"erc20Handler": "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF", ` +
					`"erc721Handler": "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07" ` +
					`}`),
			},
			requiredOptsError("genericHandler").Error(),
		},
	}

	for _, test := range tests {
		_, err := ParseChainConfig(test.input, defaultFlagSet())

		if err == nil {
			t.Fatal("expected error, got none")
		} else if err.Error() != test.expected {
			t.Fatal("unexpected error")
		}
	}
}

func compareConfigs(t *testing.T, expected, actual *Config) {
	if expected.Name != actual.Name {
		t.Fatalf("name doesn't match")
	}

	if expected.Id != actual.Id {
		t.Fatalf("id doesn't match")
	}

	if expected.Endpoint != actual.Endpoint {
		t.Fatalf("endpoint doesn't match")
	}

	if expected.From != actual.From {
		t.Fatalf("from doesn't match")
	}

	if expected.KeystorePath != actual.KeystorePath {
		t.Fatalf("keystorePath doesn't match")
	}

	if expected.BridgeContract != actual.BridgeContract {
		t.Fatalf("bridgeContract doesn't match")
	}

	if expected.ERC20HandlerContract != actual.ERC20HandlerContract {
		t.Fatalf("erc20HandlerContract doesn't match")
	}

	if expected.ERC721HandlerContract != actual.ERC721HandlerContract {
		t.Fatalf("erc721HandlerContract doesn't match")
	}

	if expected.GenericHandlerContract != actual.GenericHandlerContract {
		t.Fatalf("genericHandlerContract doesn't match")
	}

	if expected.GasLimit.Cmp(actual.GasLimit) != 0 {
		t.Fatalf("gasLimit doesn't match")
	}

	if expected.MaxGasPrice.Cmp(actual.MaxGasPrice) != 0 {
		t.Fatalf("maxGasPrice doesn't match. expected: %s actual: %s", expected.MaxGasPrice, actual.MaxGasPrice)
	}

	if expected.GasMultiplier.Cmp(actual.GasMultiplier) != 0 {
		t.Fatalf("gasMultiplier doesn't match. expected: %s actual: %s", expected.GasMultiplier.String(), actual.GasMultiplier.String())
	}

	if expected.Http != actual.Http {
		t.Fatalf("http doesn't match")
	}

	if expected.StartBlock.Cmp(actual.StartBlock) != 0 {
		t.Fatalf("startBlock doesn't match. expected: %s actual: %s", expected.StartBlock, actual.StartBlock)
	}

	if expected.BlockConfirmations.Cmp(actual.BlockConfirmations) != 0 {
		t.Fatalf("blockConfirmations doesn't match")
	}
}

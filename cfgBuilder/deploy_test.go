// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridge/keystore"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/keycard-go/hexutils"
)

const goerliEndpoint = "http://goerli.com"
const kottiEndpoint = "http://kotti.com"

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var deployerKey = "0x" + hexutils.BytesToHex(AliceKp.Encode())
var exampleContracts = &utils.DeployedContracts{
	BridgeAddress:         common.HexToAddress("0x76F5c0Da89421dC43fA000bFC3a9a7841aA3a5F3"),
	ERC20HandlerAddress:   common.HexToAddress("0x9202584Ac2A5081C6d1F27d637d1DD1Fb2AEc6B7"),
	ERC721HandlerAddress:  common.HexToAddress("0xF171e935472148298A4Fe76628193B6C2020A08a"),
	GenericHandlerAddress: common.HexToAddress("0x59105441977ecD9d805A4f5b060E34676F50F806"),
}

var exampleOptsStruct = EthOpts{
	BridgeAddress:  exampleContracts.BridgeAddress.Hex(),
	Erc20Handler:   exampleContracts.ERC20HandlerAddress.Hex(),
	Erc721Handler:  exampleContracts.ERC721HandlerAddress.Hex(),
	GenericHandler: exampleContracts.GenericHandlerAddress.Hex(),
	GasLimit:       "100",
	GasPrice:       "1000",
}

var exampleRawConfig = &RawConfig{
	DeployerPrivateKey: deployerKey,
	RelayerThreshold:   "3",
	Relayers:           []string{AliceKp.Address(), AliceKp.Address(), AliceKp.Address()},
	Chains: []EthChainConfig{
		{
			Name:     "goerli",
			ChainId:  "1",
			Endpoint: goerliEndpoint,
			GasLimit: "100",
			GasPrice: "1000",
		},
		{
			Name:     "kotti",
			ChainId:  "2",
			Endpoint: kottiEndpoint,
			GasLimit: "100",
			GasPrice: "1000",
		},
	},
}

func createTestConfig(contracts *utils.DeployedContracts) *Config {
	return &Config{
		Deployer:         AliceKp,
		RelayerThreshold: big.NewInt(3),
		Relayers:         []string{AliceKp.Address(), AliceKp.Address(), AliceKp.Address()},
		Chains: []EthChainConfig{
			{
				Name:      "goerli",
				ChainId:   "1",
				Endpoint:  goerliEndpoint,
				GasLimit:  "100",
				GasPrice:  "1000",
				contracts: contracts,
			},
			{
				Name:      "kotti",
				ChainId:   "2",
				Endpoint:  kottiEndpoint,
				GasLimit:  "100",
				GasPrice:  "1000",
				contracts: contracts,
			},
		},
	}
}

func TestCreateEthRelayerConfig(t *testing.T) {
	input := &Config{
		Deployer:         AliceKp,
		RelayerThreshold: big.NewInt(3),
		Chains: []EthChainConfig{
			{Name: "goerli", ChainId: "1", Endpoint: goerliEndpoint, contracts: exampleContracts, GasLimit: "100", GasPrice: "1000"},
		},
	}

	expected := RootConfig{Chains: []RawChainConfig{
		{
			Name:     "goerli",
			Type:     "ethereum",
			Id:       "1",
			From:     AliceKp.CommonAddress().Hex(),
			Endpoint: goerliEndpoint,
			Opts:     exampleOptsStruct,
		},
	}}

	result := constructRelayerConfig(input, AliceKp.Address())

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Mismatch.\n\tExpected: %#v\n\tGot:%#v", expected, result)
	}
}

func TestParseConfig(t *testing.T) {
	// Write to temp file
	f, err := ioutil.TempFile(os.TempDir(), "chainbridge")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	err = json.NewEncoder(f).Encode(exampleRawConfig)
	if err != nil {
		t.Fatal(err)
	}

	// Parse
	result, err := parseDeployConfig(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Verify
	expected := createTestConfig(nil)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Mismatch.\n\tExpected: %#v\n\tGot:%#v", expected, result)
	}
}

func TestCreateRelayerConfigs(t *testing.T) {
	expected := []RootConfig{
		{
			Chains: []RawChainConfig{
				{
					Name:     "goerli",
					Type:     "ethereum",
					Id:       "1",
					Endpoint: goerliEndpoint,
					From:     AliceKp.Address(),
					Opts:     exampleOptsStruct,
				},
				{
					Name:     "kotti",
					Type:     "ethereum",
					Id:       "2",
					Endpoint: kottiEndpoint,
					From:     AliceKp.Address(),
					Opts:     exampleOptsStruct,
				},
			},
		},
		{
			Chains: []RawChainConfig{
				{
					Name:     "goerli",
					Type:     "ethereum",
					Id:       "1",
					Endpoint: goerliEndpoint,
					From:     AliceKp.Address(),
					Opts:     exampleOptsStruct,
				},
				{
					Name:     "kotti",
					Type:     "ethereum",
					Id:       "2",
					Endpoint: kottiEndpoint,
					From:     AliceKp.Address(),
					Opts:     exampleOptsStruct,
				},
			},
		},
		{
			Chains: []RawChainConfig{
				{
					Name:     "goerli",
					Type:     "ethereum",
					Id:       "1",
					Endpoint: goerliEndpoint,
					From:     AliceKp.Address(),
					Opts:     exampleOptsStruct,
				},
				{
					Name:     "kotti",
					Type:     "ethereum",
					Id:       "2",
					Endpoint: kottiEndpoint,
					From:     AliceKp.Address(),
					Opts:     exampleOptsStruct,
				},
			},
		},
	}

	actual, err := CreateRelayerConfigs(createTestConfig(exampleContracts))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Mismatch.\n\tExpected: %#v\n\tGot:%#v", expected, actual)
	}
}

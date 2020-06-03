// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package evm

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ChainSafe/ChainBridge/core"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000

var BlockRetryInterval = time.Second * 5

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	Name                   string      // Human-readable chain name
	Id                     msg.ChainId // ChainID
	Endpoint               string      // url for rpc endpoint
	From                   string      // address of key to use
	KeystorePath           string      // Location of keyfiles
	BlockstorePath         string
	FreshStart             bool // Disables loading from blockstore at start
	BridgeContract         common.Address
	Erc20HandlerContract   common.Address
	Erc721HandlerContract  common.Address
	GenericHandlerContract common.Address
	GasLimit               *big.Int
	GasPrice               *big.Int
	http                   bool // Config for type of connection
	StartBlock             *big.Int
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func ParseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		Name:                   chainCfg.Name,
		Id:                     chainCfg.Id,
		Endpoint:               chainCfg.Endpoint,
		From:                   chainCfg.From,
		KeystorePath:           chainCfg.KeystorePath,
		BlockstorePath:         chainCfg.BlockstorePath,
		FreshStart:             chainCfg.FreshStart,
		BridgeContract:         utils.ZeroAddress,
		Erc20HandlerContract:   utils.ZeroAddress,
		Erc721HandlerContract:  utils.ZeroAddress,
		GenericHandlerContract: utils.ZeroAddress,
		GasLimit:               big.NewInt(DefaultGasLimit),
		GasPrice:               big.NewInt(DefaultGasPrice),
		http:                   false,
		StartBlock:             big.NewInt(0),
	}

	if contract, ok := chainCfg.Opts["bridge"]; ok && contract != "" {
		config.BridgeContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, "bridge")
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for ethereum config")
	}

	if handler, ok := chainCfg.Opts["erc20Handler"]; ok && handler != "" {
		config.Erc20HandlerContract = common.HexToAddress(handler)
		delete(chainCfg.Opts, "erc20Handler")
	} else {
		return nil, fmt.Errorf("must provide opts.erc20Handler field for ethereum config")
	}

	if handler, ok := chainCfg.Opts["erc721Handler"]; ok && handler != "" {
		config.Erc721HandlerContract = common.HexToAddress(handler)
		delete(chainCfg.Opts, "erc721Handler")
	} else {
		return nil, fmt.Errorf("must provide opts.erc721Handler field for ethereum config")
	}

	if handler, ok := chainCfg.Opts["genericHandler"]; ok && handler != "" {
		config.GenericHandlerContract = common.HexToAddress(handler)
		delete(chainCfg.Opts, "genericHandler")
	} else {
		return nil, fmt.Errorf("must provide opts.genericHandler field for ethereum config")
	}

	if gasPrice, ok := chainCfg.Opts["gasPrice"]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.GasPrice = price
			delete(chainCfg.Opts, "gasPrice")
		} else {
			return nil, errors.New("unable to parse gas price")
		}
	}

	if gasLimit, ok := chainCfg.Opts["gasLimit"]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.GasLimit = limit
			delete(chainCfg.Opts, "gasLimit")
		} else {
			return nil, errors.New("unable to parse gas limit")
		}
	}

	if HTTP, ok := chainCfg.Opts["http"]; ok && HTTP == "true" {
		config.http = true
		delete(chainCfg.Opts, "http")
	} else if HTTP, ok := chainCfg.Opts["http"]; ok && HTTP == "false" {
		config.http = false
		delete(chainCfg.Opts, "http")
	}

	if startBlock, ok := chainCfg.Opts["startBlock"]; ok && startBlock != "" {
		block := big.NewInt(0)
		_, pass := block.SetString(startBlock, 10)
		if pass {
			config.StartBlock = block
			delete(chainCfg.Opts, "startBlock")
		} else {
			return nil, errors.New("unable to parse start block")
		}
	}

	if len(chainCfg.Opts) != 0 {
		return nil, fmt.Errorf("unknown Opts Encountered: %#v", chainCfg.Opts)
	}

	return config, nil
}

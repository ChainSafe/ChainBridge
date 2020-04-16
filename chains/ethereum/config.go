// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ChainSafe/ChainBridge/core"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	name                   string      // Human-readable chain name
	id                     msg.ChainId // ChainID
	endpoint               string      // url for rpc endpoint
	from                   string      // address of key to use
	keystorePath           string      // Location of keyfiles
	blockstorePath         string
	freshStart             bool // Disables loading from blockstore at start
	bridgeContract         common.Address
	erc20HandlerContract   common.Address
	genericHandlerContract common.Address
	gasLimit               *big.Int
	gasPrice               *big.Int
	http                   bool // Config for type of connection
	startBlock             *big.Int
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:                   chainCfg.Name,
		id:                     chainCfg.Id,
		endpoint:               chainCfg.Endpoint,
		from:                   chainCfg.From,
		keystorePath:           chainCfg.KeystorePath,
		blockstorePath:         chainCfg.BlockstorePath,
		freshStart:             chainCfg.FreshStart,
		bridgeContract:         utils.ZeroAddress,
		erc20HandlerContract:   utils.ZeroAddress,
		genericHandlerContract: utils.ZeroAddress,
		gasLimit:               big.NewInt(DefaultGasLimit),
		gasPrice:               big.NewInt(DefaultGasPrice),
		http:                   false,
		startBlock:             big.NewInt(0),
	}

	if contract, ok := chainCfg.Opts["bridge"]; ok && contract != "" {
		config.bridgeContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, "bridge")
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for ethereum config")
	}

	if handler, ok := chainCfg.Opts["erc20Handler"]; ok && handler != "" {
		config.erc20HandlerContract = common.HexToAddress(handler)
		delete(chainCfg.Opts, "erc20Handler")
	} else {
		return nil, fmt.Errorf("must provide opts.erc20Handler field for ethereum config")
	}

	if handler, ok := chainCfg.Opts["genericHandler"]; ok && handler != "" {
		config.genericHandlerContract = common.HexToAddress(handler)
		delete(chainCfg.Opts, "genericHandler")
	} else {
		return nil, fmt.Errorf("must provide opts.genericHandler field for ethereum config")
	}

	if gasPrice, ok := chainCfg.Opts["gasPrice"]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.gasPrice = price
			delete(chainCfg.Opts, "gasPrice")
		} else {
			return nil, errors.New("unable to parse gas price")
		}
	}

	if gasLimit, ok := chainCfg.Opts["gasLimit"]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.gasLimit = limit
			delete(chainCfg.Opts, "gasLimit")
		} else {
			return nil, errors.New("unable to parse gas limit")
		}
	}

	if HTTP, ok := chainCfg.Opts["http"]; ok && HTTP == "true" {
		config.http = true
		delete(chainCfg.Opts, "http")
	}

	if startBlock, ok := chainCfg.Opts["startBlock"]; ok && startBlock != "" {
		block := big.NewInt(0)
		_, pass := block.SetString(startBlock, 10)
		if pass {
			config.startBlock = block
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

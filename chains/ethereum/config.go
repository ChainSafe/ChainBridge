// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ChainSafe/ChainBridge/core"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	name                 string      // Human-readable chain name
	id                   msg.ChainId // ChainID
	endpoint             string      // url for rpc endpoint
	from                 string      // address of key to use
	keystorePath         string      // Location of keyfiles
	contract             common.Address
	erc20HandlerContract common.Address
	gasLimit             *big.Int
	gasPrice             *big.Int
	http                 bool // Config for type of connection
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:                 chainCfg.Name,
		id:                   chainCfg.Id,
		endpoint:             chainCfg.Endpoint,
		from:                 chainCfg.From,
		keystorePath:         chainCfg.KeystorePath,
		contract:             ZERO_ADDRESS,
		erc20HandlerContract: ZERO_ADDRESS,
		gasLimit:             big.NewInt(DefaultGasLimit),
		gasPrice:             big.NewInt(DefaultGasPrice),
		http:                 false,
	}

	if contract, ok := chainCfg.Opts["contract"]; ok && contract != "" {
		config.contract = common.HexToAddress(contract)
		delete(chainCfg.Opts, "contract")
	} else {
		return nil, fmt.Errorf("must provide opts.contract field for ethereum config")
	}

	if gasPrice, ok := chainCfg.Opts["gasPrice"]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.gasPrice = price
			delete(chainCfg.Opts, "gasPrice")
		} else {
			return nil, errors.New("Unable to parse gas price.")
		}
	}

	if gasLimit, ok := chainCfg.Opts["gasLimit"]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.gasLimit = limit
			delete(chainCfg.Opts, "gasLimit")
		} else {
			return nil, errors.New("Unable to parse gas limit.")
		}
	}

	if HTTP, ok := chainCfg.Opts["http"]; ok && HTTP == "true" {
		config.http = true
		delete(chainCfg.Opts, "http")
	}

	if len(chainCfg.Opts) != 0 {
		return nil, errors.New("Unknown Opts Encountered")
	}

	return config, nil
}

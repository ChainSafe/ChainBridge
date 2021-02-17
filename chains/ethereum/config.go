// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"fmt"
	"math/big"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000
const DefaultBlockConfirmations = 10
const DefaultGasMultiplier = 1

// Chain specific options
var (
	BridgeOpt             = "bridge"
	Erc20HandlerOpt       = "erc20Handler"
	Erc721HandlerOpt      = "erc721Handler"
	GenericHandlerOpt     = "genericHandler"
	MaxGasPriceOpt        = "maxGasPrice"
	GasLimitOpt           = "gasLimit"
	GasMultiplier         = "gasMultiplier"
	HttpOpt               = "http"
	StartBlockOpt         = "startBlock"
	BlockConfirmationsOpt = "blockConfirmations"
)

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
	erc721HandlerContract  common.Address
	genericHandlerContract common.Address
	gasLimit               *big.Int
	maxGasPrice            *big.Int
	gasMultiplier          *big.Float
	http                   bool // Config for type of connection
	startBlock             *big.Int
	blockConfirmations     *big.Int
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
		erc721HandlerContract:  utils.ZeroAddress,
		genericHandlerContract: utils.ZeroAddress,
		gasLimit:               big.NewInt(DefaultGasLimit),
		maxGasPrice:            big.NewInt(DefaultGasPrice),
		gasMultiplier:          big.NewFloat(DefaultGasMultiplier),
		http:                   false,
		startBlock:             big.NewInt(0),
		blockConfirmations:     big.NewInt(0),
	}

	if contract, ok := chainCfg.Opts[BridgeOpt]; ok && contract != "" {
		config.bridgeContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, BridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for ethereum config")
	}

	config.erc20HandlerContract = common.HexToAddress(chainCfg.Opts[Erc20HandlerOpt])
	delete(chainCfg.Opts, Erc20HandlerOpt)

	config.erc721HandlerContract = common.HexToAddress(chainCfg.Opts[Erc721HandlerOpt])
	delete(chainCfg.Opts, Erc721HandlerOpt)

	config.genericHandlerContract = common.HexToAddress(chainCfg.Opts[GenericHandlerOpt])
	delete(chainCfg.Opts, GenericHandlerOpt)

	if gasPrice, ok := chainCfg.Opts[MaxGasPriceOpt]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.maxGasPrice = price
			delete(chainCfg.Opts, MaxGasPriceOpt)
		} else {
			return nil, errors.New("unable to parse max gas price")
		}
	}

	if gasLimit, ok := chainCfg.Opts[GasLimitOpt]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.gasLimit = limit
			delete(chainCfg.Opts, GasLimitOpt)
		} else {
			return nil, errors.New("unable to parse gas limit")
		}
	}

	if gasMultiplier, ok := chainCfg.Opts[GasMultiplier]; ok {
		multilier := big.NewFloat(1)
		_, pass := multilier.SetString(gasMultiplier)
		if pass {
			config.gasMultiplier = multilier
			delete(chainCfg.Opts, GasMultiplier)
		} else {
			return nil, errors.New("unable to parse gasMultiplier to float")
		}
	}

	if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "true" {
		config.http = true
		delete(chainCfg.Opts, HttpOpt)
	} else if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "false" {
		config.http = false
		delete(chainCfg.Opts, HttpOpt)
	}

	if startBlock, ok := chainCfg.Opts[StartBlockOpt]; ok && startBlock != "" {
		block := big.NewInt(0)
		_, pass := block.SetString(startBlock, 10)
		if pass {
			config.startBlock = block
			delete(chainCfg.Opts, StartBlockOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", StartBlockOpt)
		}
	}

	if blockConfirmations, ok := chainCfg.Opts[BlockConfirmationsOpt]; ok && blockConfirmations != "" {
		val := big.NewInt(DefaultBlockConfirmations)
		_, pass := val.SetString(blockConfirmations, 10)
		if pass {
			config.blockConfirmations = val
			delete(chainCfg.Opts, BlockConfirmationsOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", BlockConfirmationsOpt)
		}
	} else {
		config.blockConfirmations = big.NewInt(DefaultBlockConfirmations)
		delete(chainCfg.Opts, BlockConfirmationsOpt)
	}

	if len(chainCfg.Opts) != 0 {
		return nil, fmt.Errorf("unknown Opts Encountered: %#v", chainCfg.Opts)
	}

	return config, nil
}

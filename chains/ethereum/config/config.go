// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"encoding/json"
	"fmt"
	"github.com/ChainSafe/ChainBridge/chains/ethereum/client"
	"github.com/ChainSafe/ChainBridge/config"
	"github.com/ChainSafe/ChainBridge/config/flags"
	"github.com/urfave/cli/v2"
	"math/big"
	"strconv"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultMaxGasPrice = 20000000000
const DefaultBlockConfirmations = 10
const DefaultGasMultiplier = 1

func requiredOptsError(field string) error {
	return fmt.Errorf("required field %s missing in ethereum config", field)
}

func invalidOptsError(field, value string) error {
	return fmt.Errorf("invalid value for field %s (%s) in ethereum config", field, value)
}

type RawOpts struct {
	Bridge             common.Address `json:"bridge"`
	Erc20Handler       common.Address `json:"erc20Handler"`
	Erc721Handler      common.Address `json:"erc721Handler"`
	GenericHandler     common.Address `json:"genericHandler"`
	MaxGasPrice        *big.Int       `json:"maxGasPrice"`
	GasLimit           *big.Int       `json:"gasLimit"`
	GasMultiplier      float64        `json:"gasMultiplier"` // Prefer not to use big.Float which requires string for unmarshal
	Http               bool           `json:"http"`
	StartBlock         *big.Int       `json:"startBlock"`
	BlockConfirmations *big.Int       `json:"blockConfirmations"`
}

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	Name                   string      // Human-readable chain name
	Id                     msg.ChainId // ChainID
	Endpoint               string      // url for rpc endpoint
	From                   string      // address of key to use
	KeystorePath           string      // Location of keyfiles
	Insecure               bool
	BlockstorePath         string
	FreshStart             bool // Disables loading from blockstore at start
	BridgeContract         common.Address
	ERC20HandlerContract   common.Address
	ERC721HandlerContract  common.Address
	GenericHandlerContract common.Address
	GasLimit               *big.Int
	MaxGasPrice            *big.Int
	GasMultiplier          *big.Float
	Http                   bool // Config for type of connection
	StartBlock             *big.Int
	BlockConfirmations     *big.Int
	LatestBlock            bool
}

func (cfg *Config) EnsureContractsHaveBytecode(conn *client.Client) error {
	err := conn.EnsureHasBytecode(cfg.BridgeContract)
	if err != nil {
		return err
	}
	err = conn.EnsureHasBytecode(cfg.ERC20HandlerContract)
	if err != nil {
		return err
	}
	err = conn.EnsureHasBytecode(cfg.ERC721HandlerContract)
	if err != nil {
		return err
	}
	err = conn.EnsureHasBytecode(cfg.GenericHandlerContract)
	if err != nil {
		return err
	}
	return nil
}

// TODO: docs
func ParseChainConfig(chainCfg *config.RawChainConfig, ctx *cli.Context) (*Config, error) {
	opts := new(RawOpts)
	err := json.Unmarshal(chainCfg.Opts, opts)
	if err != nil {
		return nil, err
	}

	// Check for testkey flag, use the test key as path if used
	var ks string
	var insecure bool
	if key := ctx.String(flags.TestKeyFlag.Name); key != "" {
		ks = key
		insecure = true
	} else {
		ks = ctx.String(flags.KeystorePathFlag.Name)
	}

	chainId, err := strconv.Atoi(chainCfg.Id)
	if err != nil {
		return nil, err
	}

	if opts.Bridge == utils.ZeroAddress {
		return nil, requiredOptsError("bridge")
	}

	if opts.Erc20Handler == utils.ZeroAddress {
		return nil, requiredOptsError("erc20Handler")
	}

	if opts.Erc721Handler == utils.ZeroAddress {
		return nil, requiredOptsError("erc721Handler")
	}

	if opts.GenericHandler == utils.ZeroAddress {
		return nil, requiredOptsError("genericHandler")
	}

	if opts.GasLimit == nil {
		opts.GasLimit = big.NewInt(DefaultGasLimit)
	} else if opts.GasLimit.Cmp(big.NewInt(0)) < 1 { // Must be > 0
		return nil, invalidOptsError("gasLimit", opts.GasLimit.String())
	}

	if opts.MaxGasPrice == nil {
		opts.MaxGasPrice = big.NewInt(DefaultMaxGasPrice)
	} else if opts.MaxGasPrice.Cmp(big.NewInt(0)) < 1 { // Must be > 0
		return nil, invalidOptsError("maxGasPrice", opts.MaxGasPrice.String())
	}

	var gasMultiplier *big.Float
	if opts.GasMultiplier < 0 {
		return nil, invalidOptsError("gasMultiplier", fmt.Sprintf("%f64", opts.GasMultiplier))
	} else if opts.GasMultiplier == 0 {
		gasMultiplier = big.NewFloat(1)
	} else {
		gasMultiplier = big.NewFloat(opts.GasMultiplier)
	}

	if opts.StartBlock == nil {
		opts.StartBlock = big.NewInt(0)
	} else if opts.StartBlock.Cmp(big.NewInt(0)) < 1 { // Must be > 0
		return nil, invalidOptsError("startBlock", opts.StartBlock.String())
	}

	if opts.BlockConfirmations == nil {
		opts.BlockConfirmations = big.NewInt(DefaultBlockConfirmations)
	} else if opts.BlockConfirmations.Cmp(big.NewInt(0)) < 1 { // Must be > 0
		return nil, invalidOptsError("blockConfirmations", opts.BlockConfirmations.String())
	}

	return &Config{
		Name:                   chainCfg.Name,
		Id:                     msg.ChainId(chainId),
		Endpoint:               chainCfg.Endpoint,
		From:                   chainCfg.From,
		KeystorePath:           ks,
		Insecure:               insecure,
		BlockstorePath:         ctx.String(flags.BlockstorePathFlag.Name),
		FreshStart:             ctx.Bool(flags.FreshStartFlag.Name),
		LatestBlock:            ctx.Bool(flags.LatestBlockFlag.Name),
		BridgeContract:         opts.Bridge,
		ERC20HandlerContract:   opts.Erc20Handler,
		ERC721HandlerContract:  opts.Erc721Handler,
		GenericHandlerContract: opts.GenericHandler,
		GasLimit:               opts.GasLimit,
		MaxGasPrice:            opts.MaxGasPrice,
		GasMultiplier:          gasMultiplier,
		Http:                   opts.Http,
		StartBlock:             opts.StartBlock,
		BlockConfirmations:     opts.BlockConfirmations,
	}, nil
}

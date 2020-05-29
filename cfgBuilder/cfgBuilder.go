// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package cfgBuilder

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	log "github.com/ChainSafe/log15"
	"github.com/naoina/toml"
)

type RawConfig struct {
	RelayerThreshold string           `json:"relayerThreshold"`
	Relayers         []string         `json:"relayers"`
	Chains           []EthChainConfig `json:"chains"`
}

type Config struct {
	RelayerThreshold *big.Int
	Relayers         []string
	Chains           []EthChainConfig
}

// Identical to config.RawChainConfig, but uses struct for opts to get desired output formatting
type RawChainConfig struct {
	Name     string  `toml:"name"`
	Type     string  `toml:"type"`
	Id       string  `toml:"id"`       // ChainID
	Endpoint string  `toml:"endpoint"` // url for rpc endpoint
	From     string  `toml:"from"`     // address of key to use
	Opts     EthOpts `toml:"opts"`
}

// Replicates config.Config
type RootConfig struct {
	Chains []RawChainConfig
}

type EthOpts struct {
	BridgeAddress  string `toml:"bridge"`
	Erc20Handler   string `toml:"erc20Handler"`
	Erc721Handler  string `toml:"erc721Handler"`
	GenericHandler string `toml:"genericHandler"`
	GasLimit       string `toml:"gasLimit"`
	GasPrice       string `toml:"gasPrice"`
	StartBlock     string `toml:"startBlock"`
	Http           string `toml:"http"`
}

type EthChainConfig struct {
	Name           string `json:"name"`
	ChainId        string `json:"chainId"`
	Endpoint       string `json:"endpoint"`
	BridgeAddress  string `json:"bridge"`
	Erc20Handler   string `json:"erc20Handler"`
	Erc721Handler  string `json:"erc721Handler"`
	GenericHandler string `json:"genericHandler"`
	GasLimit       string `json:"gasLimit"`
	GasPrice       string `json:"gasPrice"`
	StartBlock     string `json:"startBlock"`
	Http           string `json:"http"`
}

// ToToml writes the config to a file
func (c *RootConfig) ToTOML(file string) *os.File {
	var (
		newFile *os.File
		err     error
	)

	var raw []byte
	if raw, err = toml.Marshal(*c); err != nil {
		log.Warn("error marshalling toml", "err", err)
		os.Exit(1)
	}

	newFile, err = os.Create(file)
	if err != nil {
		log.Warn("error creating config file", "err", err)
	}
	_, err = newFile.Write(raw)
	if err != nil {
		log.Warn("error writing to config file", "err", err)
	}

	if err := newFile.Close(); err != nil {
		log.Warn("error closing file", "err", err)
	}
	return newFile
}

func constructEthChainConfig(cfg EthChainConfig, relayer string) RawChainConfig {
	return RawChainConfig{
		Name:     cfg.Name,
		Type:     "ethereum",
		From:     relayer,
		Id:       cfg.ChainId,
		Endpoint: cfg.Endpoint,
		Opts: EthOpts{
			BridgeAddress:  cfg.BridgeAddress,
			Erc20Handler:   cfg.Erc20Handler,
			Erc721Handler:  cfg.Erc721Handler,
			GenericHandler: cfg.GenericHandler,
			GasLimit:       cfg.GasLimit,
			GasPrice:       cfg.GasPrice,
			StartBlock:     cfg.StartBlock,
			Http:           cfg.Http,
		},
	}
}

func constructRelayerConfig(cfg *Config, relayer string) RootConfig {
	// Create RawConfig structs from the provided Chains
	var rawCfgs []RawChainConfig
	for _, chain := range cfg.Chains {
		raw := constructEthChainConfig(chain, relayer)
		rawCfgs = append(rawCfgs, raw)
	}

	return RootConfig{Chains: rawCfgs}
}

func parseRawConfig(raw *RawConfig) (*Config, error) {
	var res Config

	threshold, ok := big.NewInt(0).SetString(raw.RelayerThreshold, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse relayer threshold")
	}
	res.RelayerThreshold = threshold
	res.Relayers = raw.Relayers
	res.Chains = raw.Chains

	return &res, nil
}

func ParseDeployConfig(path string) (*Config, error) {
	fp, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath.Clean(fp))
	if err != nil {
		return nil, err
	}

	var rawCfg RawConfig
	err = json.NewDecoder(f).Decode(&rawCfg)
	if err != nil {
		return nil, err
	}

	return parseRawConfig(&rawCfg)
}

// CreateRelayerConfigs takes a prepared config and constructs the configs for each relayer
func CreateRelayerConfigs(cfg *Config) ([]RootConfig, error) {
	var configs []RootConfig
	for _, relayer := range cfg.Relayers {
		rCfg := constructRelayerConfig(cfg, relayer)
		configs = append(configs, rCfg)
	}

	return configs, nil
}

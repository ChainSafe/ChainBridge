// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	log "github.com/ChainSafe/log15"
	"github.com/naoina/toml"
	"github.com/status-im/keycard-go/hexutils"
)

type RawConfig struct {
	DeployerPrivateKey string           `json:"deployerKey"`
	RelayerThreshold   string           `json:"relayerThreshold"`
	Relayers           []string         `json:"relayers"`
	Chains             []EthChainConfig `json:"chains"`
}

type Config struct {
	Deployer         *secp256k1.Keypair
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
}

type EthChainConfig struct {
	Name      string `json:"name"`
	ChainId   string `json:"chainId"`
	Endpoint  string `json:"endpoint"`
	contracts *utils.DeployedContracts
	GasLimit  string `json:"gasLimit"`
	GasPrice  string `json:"gasPrice"`
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

// DeployContracts deploys the default set of contracts to each chain and adds the addresses
// to the config.
func DeployContracts(cfg *Config) error {
	for i, chain := range cfg.Chains {
		log.Info("Deploying contracts", "endpoint", chain.Endpoint)
		// TODO: Replace when utils.Client updates land
		// TODO: Pass in gaslimit and price
		pk := hexutils.BytesToHex(cfg.Deployer.Encode())
		chainId, err := strconv.Atoi(chain.ChainId)
		if err != nil {
			return err
		}
		c, err := utils.DeployContracts(pk, uint8(chainId), chain.Endpoint, cfg.RelayerThreshold)
		if err != nil {
			return err
		}
		cfg.Chains[i].contracts = c
	}
	return nil
}

func construcEthChainConfig(cfg EthChainConfig, relayer string) RawChainConfig {
	return RawChainConfig{
		Name:     cfg.Name,
		Type:     "ethereum",
		From:     relayer,
		Id:       cfg.ChainId,
		Endpoint: cfg.Endpoint,
		Opts: EthOpts{
			BridgeAddress:  cfg.contracts.BridgeAddress.Hex(),
			Erc20Handler:   cfg.contracts.ERC20HandlerAddress.Hex(),
			Erc721Handler:  cfg.contracts.ERC721HandlerAddress.Hex(),
			GenericHandler: cfg.contracts.GenericHandlerAddress.Hex(),
			GasLimit:       cfg.GasLimit,
			GasPrice:       cfg.GasPrice,
			// TODO: Add start block
		},
	}
}

func constructRelayerConfig(cfg *Config, relayer string) RootConfig {
	// Create RawConfig structs from the provided Chains
	var rawCfgs []RawChainConfig
	for _, chain := range cfg.Chains {
		raw := construcEthChainConfig(chain, relayer)
		rawCfgs = append(rawCfgs, raw)
	}

	return RootConfig{Chains: rawCfgs}
}

func parseRawConfig(raw *RawConfig) (*Config, error) {
	var res Config
	deployer, err := secp256k1.NewKeypairFromString(raw.DeployerPrivateKey[2:])
	if err != nil {
		return nil, err
	}
	res.Deployer = deployer

	threshold, ok := big.NewInt(0).SetString(raw.RelayerThreshold, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse relayer threshold")
	}
	res.RelayerThreshold = threshold

	res.Relayers = raw.Relayers

	res.Chains = raw.Chains

	return &res, nil
}

func parseDeployConfig(path string) (*Config, error) {
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

func main() {
	// Pares first argument for path
	if len(os.Args) < 2 {
		log.Error("Please specify path to config json")
		os.Exit(1)
	}
	path := os.Args[1]
	if path == "" {
		log.Error("failed to parse path argument")
		os.Exit(1)
	}

	// Read in the config
	cfg, err := parseDeployConfig(path)
	if err != nil {
		log.Error("failed to parse config", "err", err)
		os.Exit(1)
	}

	// Deploy the contracts
	err = DeployContracts(cfg)
	if err != nil {
		log.Error("failed to deploy contracts", "err", err)
		os.Exit(1)
	}

	// Construct the individual relayer configs
	relayerCfgs, err := CreateRelayerConfigs(cfg)
	if err != nil {
		log.Error("failed to construct relayer configs", "err", err)
		os.Exit(1)
	}

	// Write all the configs to files
	for i, cfg := range relayerCfgs {
		cfg.ToTOML(fmt.Sprintf("config%d.toml", i))
	}
}

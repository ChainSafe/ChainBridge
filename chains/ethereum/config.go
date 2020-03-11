// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ethereum/go-ethereum/common"
)

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	id       msg.ChainId        // ChainID
	chainID  *big.Int           // Ethereum chain ID
	endpoint string             // url for rpc endpoint
	from     string             // address of key to use
	keystore *keystore.Keystore // Location of keyfiles
	contract common.Address
	gasLimit *big.Int
	gasPrice *big.Int
}

// ParseChainConfig uses a core.ChainConfig to construct a corresponding Config
func ParseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		id:       chainCfg.Id,
		endpoint: chainCfg.Endpoint,
		from:     chainCfg.From,
		keystore: chainCfg.Keystore,
		chainID:  big.NewInt(1),
		contract: common.HexToAddress("0x0"),
		gasLimit: big.NewInt(6721975),
		gasPrice: big.NewInt(20000000000),
	}

	if chainID, ok := chainCfg.Opts["chainID"]; ok {
		config.chainID.SetString(chainID, 10)
	}

	if bridge, ok := chainCfg.Opts["contract"]; ok {
		config.contract = common.HexToAddress(bridge)
	}

	if gasPrice, ok := chainCfg.Opts["gasPrice"]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.gasPrice = price
		} else {
			return nil, errors.New("Unable to parse gas price.")
		}
	}

	if gasLimit, ok := chainCfg.Opts["gasLimit"]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.gasLimit = limit
		} else {
			return nil, errors.New("Unable to parse gas limit.")
		}
	}

	return config, nil
}

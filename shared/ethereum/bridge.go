// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetDepositNonce(client *ethclient.Client, bridge common.Address, chain msg.ChainId) (uint64, error) {
	instance, err := Bridge.NewBridge(bridge, client)
	if err != nil {
		return 0, err
	}

	count, err := instance.DepositCounts(&bind.CallOpts{}, uint8(chain))
	if err != nil {
		return 0, err
	}

	return count.Uint64(), nil
}

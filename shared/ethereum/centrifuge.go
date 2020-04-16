// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	centrifugeHandler "github.com/ChainSafe/ChainBridge/bindings/CentrifugeAssetHandler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func HashExists(client *ethclient.Client, hash [32]byte, contract common.Address) (bool, error) {
	instance, err := centrifugeHandler.NewCentrifugeAssetHandler(contract, client)
	if err != nil {
		return false, err
	}

	exists, err := instance.GetHash(&bind.CallOpts{}, hash)
	if err != nil {
		return false, err
	}
	return exists, nil
}

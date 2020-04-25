// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ChainSafe/ChainBridge/bindings/CentrifugeAsset"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployAssetStore(client *ethclient.Client, opts *bind.TransactOpts) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	addr, _, _, err := CentrifugeAsset.DeployCentrifugeAsset(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	return addr, nil
}

func HashExists(client *ethclient.Client, hash [32]byte, contract common.Address) (bool, error) {
	instance, err := CentrifugeAsset.NewCentrifugeAsset(contract, client)
	if err != nil {
		return false, err
	}

	exists, err := instance.AssetsStored(&bind.CallOpts{}, hash)
	if err != nil {
		return false, err
	}
	return exists, nil
}

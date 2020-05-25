// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ChainSafe/ChainBridge/bindings/CentrifugeAsset"
	"github.com/ethereum/go-ethereum/common"
)

func DeployAssetStore(client *Client) (common.Address, error) {
	err := UpdateNonce(client)
	if err != nil {
		return ZeroAddress, err
	}

	addr, tx, _, err := CentrifugeAsset.DeployCentrifugeAsset(client.Opts, client.Client)
	if err != nil {
		return ZeroAddress, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return addr, nil
}

func HashExists(client *Client, hash [32]byte, contract common.Address) (bool, error) {
	instance, err := CentrifugeAsset.NewCentrifugeAsset(contract, client.Client)
	if err != nil {
		return false, err
	}

	exists, err := instance.AssetsStored(client.CallOpts, hash)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterResource(client *Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address) error {
	instance, err := Bridge.NewBridge(bridge, client.Client)
	if err != nil {
		return err
	}

	err = UpdateNonce(client)
	if err != nil {
		return err
	}
	tx, err := instance.AdminSetResource(client.Opts, handler, rId, addr)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	return nil
}

func RegisterGenericResource(client *Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address, depositSig, executeSig [4]byte) error {
	instance, err := Bridge.NewBridge(bridge, client.Client)
	if err != nil {
		return err
	}

	err = UpdateNonce(client)
	if err != nil {
		return err
	}
	tx, err := instance.AdminSetGenericResource(client.Opts, handler, rId, addr, depositSig, executeSig)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	return nil
}

func SetBurnable(client *Client, bridge, handler, contract common.Address) error {
	instance, err := Bridge.NewBridge(bridge, client.Client)
	if err != nil {
		return err
	}

	err = UpdateNonce(client)
	if err != nil {
		return err
	}

	tx, err := instance.AdminSetBurnable(client.Opts, handler, contract)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	return nil
}

func GetDepositNonce(client *Client, bridge common.Address, chain msg.ChainId) (uint64, error) {
	instance, err := Bridge.NewBridge(bridge, client.Client)
	if err != nil {
		return 0, err
	}

	count, err := instance.DepositCounts(client.CallOpts, uint8(chain))
	if err != nil {
		return 0, err
	}

	return count.Uint64(), nil
}

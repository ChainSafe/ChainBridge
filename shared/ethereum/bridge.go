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

func RegisterResource(client *utils.Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address) error {
	instance, err := Bridge.NewBridge(bridge, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(client)
	if err != nil {
		return err
	}
	_, err = instance.AdminSetResource(handler, rId, addr)
	if err != nil {
		return err
	}
	return nil
}

func RegisterGenericResource(client *utils.Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address, depositSig, executeSig [4]byte) error {
	instance, err := Bridge.NewBridge(bridge, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(client)
	if err != nil {
		return err
	}
	_, err = instance.AdminSetGenericResource(handler, rId, addr, depositSig, executeSig)
	if err != nil {
		return err
	}
	return nil
}

func SetBurnable(client *utils.Client, bridge, handler, contract common.Address) error {
	instance, err := Bridge.NewBridge(bridge, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(client)
	if err != nil {
		return err
	}

	_, err = instance.AdminSetBurnable(handler, contract)
	if err != nil {
		return err
	}
	return nil
}

func GetDepositNonce(client *utils.Client, bridge common.Address, chain msg.ChainId) (uint64, error) {
	instance, err := Bridge.NewBridge(bridge, client)
	if err != nil {
		return 0, err
	}

	count, err := instance.DepositCounts(client.CallOpts{}, uint8(chain))
	if err != nil {
		return 0, err
	}

	return count.Uint64(), nil
}

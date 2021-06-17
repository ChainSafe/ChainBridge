// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"fmt"
	"math/big"

	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

func QueryStorage(client *Client, prefix, method string, arg1, arg2 []byte, result interface{}) (bool, error) {
	key, err := types.CreateStorageKey(client.Meta, prefix, method, arg1, arg2)
	if err != nil {
		return false, err
	}
	return client.Api.RPC.State.GetStorageLatest(key, result)
}

func getConst(meta *types.Metadata, prefix, name string, res interface{}) error {
	data, err := meta.FindConstantValue(prefix, name)
	if err != nil {
		return err
	}
	return types.DecodeFromBytes(data, res)
}

// QueryConst looks up a constant in the metadata
func QueryConst(client *Client, prefix, name string, res interface{}) error {
	err := getConst(client.Meta, prefix, name, res)
	if err != nil {
		return err
	}
	return nil
}

// BalanceOf returns the free balance of an account
func BalanceOf(client *Client, publicKey []byte) (*big.Int, error) {
	var acct types.AccountInfo

	ok, err := QueryStorage(client, "System", "Account", publicKey, nil, &acct)
	if err != nil {
		return nil, err
	} else if !ok {
		return big.NewInt(0), nil
	}
	return acct.Data.Free.Int, nil
}

func GetErc721Token(client *Client, id types.U256) (*Erc721Token, error) {
	var res Erc721Token
	tokenIdBz, err := types.EncodeToBytes(id)
	if err != nil {
		return nil, err
	}
	exists, err := QueryStorage(client, "TokenStorage", "Tokens", tokenIdBz, nil, &res)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("token %s does not exist", id.String())
	}
	return &res, nil
}

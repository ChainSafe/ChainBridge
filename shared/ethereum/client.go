// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"context"
	"math/big"

	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000

type Client struct {
	Client   *ethclient.Client
	Opts     *bind.TransactOpts
	CallOpts *bind.CallOpts
}

func NewClient(endpoint string, kp *secp256k1.Keypair) (*Client, error) {
	ctx := context.Background()
	rpcClient, err := rpc.DialWebsocket(ctx, endpoint, "/ws")
	if err != nil {
		return nil, err
	}
	client := ethclient.NewClient(rpcClient)

	nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(kp.Address()))
	if err != nil {
		return nil, err
	}
	opts := bind.NewKeyedTransactor(kp.PrivateKey())
	opts.Nonce = big.NewInt(int64(nonce - 1)) // -1 since we always increment before calling
	opts.Value = big.NewInt(0)                // in wei
	opts.GasLimit = uint64(DefaultGasLimit)   // in units
	opts.GasPrice = big.NewInt(DefaultGasPrice)
	opts.Context = ctx

	return &Client{
		Client: client,
		Opts:   opts,
		CallOpts: &bind.CallOpts{
			From: opts.From,
		},
	}, nil
}

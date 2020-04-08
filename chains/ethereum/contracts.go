// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20 "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc721 "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BridgeContract struct {
	BridgeFilterer
	BridgeCaller
	BridgeRaw
	BridgeTransactor
}

type ERC20HandlerContract struct {
	ERC20HandlerFilterer
	ERC20HandlerCaller
	ERC20HandlerRaw
}

type ERC721HandlerContract struct {
	ERC721HandlerFilterer
	ERC721HandlerCaller
	ERC721HandlerRaw
}

type BridgeFilterer interface {
}

type BridgeTransactor interface {
	Deposit(opts *bind.TransactOpts, chainID *big.Int, originChainHandlerAddress common.Address, data []byte) (*types.Transaction, error)
}

type BridgeCaller interface {
	GetDepositProposal(opts *bind.CallOpts, destinationChainID *big.Int, depositNonce *big.Int) (bridge.BridgeDepositProposal, error)
}

type BridgeRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

type ERC20HandlerFilterer interface {
}

type ERC20HandlerCaller interface {
	GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (erc20.ERC20HandlerDepositRecord, error)
}

type ERC20HandlerRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

type ERC721HandlerFilterer interface{}

type ERC721HandlerCaller interface {
	GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (erc721.ERC721HandlerDepositRecord, error)
}

type ERC721HandlerRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{})
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

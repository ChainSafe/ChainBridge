// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20 "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc721 "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
)

// genericDepositRecord is the return value from the solidity function getGenericDepositRecord()
type genericDepositRecord struct {
	OriginChainTokenAddress   common.Address
	OriginChainHandlerAddress common.Address
	DestChainID               *big.Int
	DestChainHandlerAddress   common.Address
	DestRecipientAddress      common.Address
	Data                      []byte
}

type erc20DepositRecord struct {
	OriginChainTokenAddress   common.Address
	OriginChainHandlerAddress common.Address
	DestChainID               *big.Int
	DestChainHandlerAddress   common.Address
	DestRecipientAddress      common.Address
	Amount                    *big.Int
}

// depositProposal is the return value from the solidity function getDepositProposal()
type depositProposal struct {
	OriginChainID *big.Int
	DepositNonce  *big.Int
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        string
}

type BridgeContract struct {
	BridgeFilterer
	BridgeCaller
	BridgeRaw
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

type BridgeCaller interface {
	GetDepositProposal(*bind.CallOpts, *big.Int, common.Address, *big.Int) (bridge.BridgeDepositProposal, error)
}

type BridgeRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)

	// 	FilterERCTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*bridge.BridgeERCTransferIterator, error)
	// 	FilterGenericTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*bridge.BridgeGenericTransferIterator, error)
	// 	FilterNFTTransfer(opts *bind.FilterOpts, _destChain []*big.Int, _depositNonce []*big.Int) (*bridge.BridgeNFTTransferIterator, error)
}

type ERC20HandlerFilterer interface {
}

type ERC20HandlerCaller interface {
	GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (erc20.ERC20HandlerDepositRecord, error)
}

type ERC20HandlerRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{})
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

type ERC721HandlerFilterer interface {
}

type ERC721HandlerCaller interface {
	GetDepositRecord(opts *bind.CallOpts, depositID *big.Int) (erc721.ERC721HandlerDepositRecord, error)
}

type ERC721HandlerRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{})
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

func UnpackGenericDepositRecord(args ...interface{}) (genericDepositRecord, error) {
	if args[6] != nil {
		return genericDepositRecord{}, args[6].(error)
	}
	return genericDepositRecord{
			OriginChainTokenAddress:   args[0].(common.Address),
			OriginChainHandlerAddress: args[1].(common.Address),
			DestChainID:               args[2].(*big.Int),
			DestChainHandlerAddress:   args[3].(common.Address),
			DestRecipientAddress:      args[4].(common.Address),
			Data:                      args[5].([]byte),
		},
		nil
}

func UnpackErc20DepositRecord(args ...interface{}) (erc20DepositRecord, error) {
	if args[6] != nil {
		return erc20DepositRecord{}, args[6].(error)
	}
	return erc20DepositRecord{
			OriginChainTokenAddress:   args[0].(common.Address),
			OriginChainHandlerAddress: args[1].(common.Address),
			DestChainID:               args[2].(*big.Int),
			DestChainHandlerAddress:   args[3].(common.Address),
			DestRecipientAddress:      args[4].(common.Address),
			Amount:                    args[5].(*big.Int),
		},
		nil
}

func UnpackDepositProposal(args ...interface{}) (depositProposal, error) {
	if args[6] != nil {
		return depositProposal{}, args[6].(error)
	}
	return depositProposal{
			OriginChainID: args[0].(*big.Int),
			DepositNonce:  args[1].(*big.Int),
			DataHash:      args[2].([32]byte),
			YesVotes:      args[3].([]common.Address),
			NoVotes:       args[4].([]common.Address),
			Status:        args[5].(string),
		},
		nil
}

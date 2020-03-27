// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

type BridgeFilterer interface {
}

type BridgeCaller interface {
	GetGenericDepositRecord(opts *bind.CallOpts, originChainID *big.Int, depositNonce *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, []byte, error)
	GetERC20DepositRecord(opts *bind.CallOpts, originChainID *big.Int, depositNonce *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, *big.Int, error)
	GetDepositProposal(opts *bind.CallOpts, destinationChainID *big.Int, depositNonce *big.Int) (*big.Int, *big.Int, [32]byte, []common.Address, []common.Address, string, error)
}

type BridgeRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)

	// 	FilterERCTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*bridge.BridgeERCTransferIterator, error)
	// 	FilterGenericTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*bridge.BridgeGenericTransferIterator, error)
	// 	FilterNFTTransfer(opts *bind.FilterOpts, _destChain []*big.Int, _depositNonce []*big.Int) (*bridge.BridgeNFTTransferIterator, error)
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

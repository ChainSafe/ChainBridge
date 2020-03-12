// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type GenericDepositRecord struct {
	OriginChainTokenAddress   common.Address
	OriginChainHandlerAddress common.Address
	DestChainID               *big.Int
	DestChainHandlerAddress   common.Address
	DestRecipientAddress      common.Address
	Data                      []byte
}

type Erc20DepositRecord struct {
	OriginChainTokenAddress   common.Address
	OriginChainHandlerAddress common.Address
	DestChainID               *big.Int
	DestChainHandlerAddress   common.Address
	DestRecipientAddress      common.Address
	Amount                    *big.Int
}

type DepositProposal struct {
	OriginChainID *big.Int
	DepositID     *big.Int
	DataHash      [32]byte
	NumYes        *big.Int
	NumNo         *big.Int
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
	GetGenericDepositRecord(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, []byte, error)
	GetERC20DepositRecord(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int) (common.Address, common.Address, *big.Int, common.Address, common.Address, *big.Int, error)
	GetDepositProposal(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int) (*big.Int, *big.Int, [32]byte, *big.Int, *big.Int, string, error)
}

type BridgeRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)

	// 	FilterERCTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*bridge.BridgeERCTransferIterator, error)
	// 	FilterGenericTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*bridge.BridgeGenericTransferIterator, error)
	// 	FilterNFTTransfer(opts *bind.FilterOpts, _destChain []*big.Int, _depositId []*big.Int) (*bridge.BridgeNFTTransferIterator, error)
}

func UnpackGenericDepositRecord(args ...interface{}) (GenericDepositRecord, error) {
	if args[6] != nil {
		return GenericDepositRecord{}, args[6].(error)
	}
	return GenericDepositRecord{
			OriginChainTokenAddress:   args[0].(common.Address),
			OriginChainHandlerAddress: args[1].(common.Address),
			DestChainID:               args[2].(*big.Int),
			DestChainHandlerAddress:   args[3].(common.Address),
			DestRecipientAddress:      args[4].(common.Address),
			Data:                      args[5].([]byte),
		},
		nil
}

func UnpackErc20DepositRecord(args ...interface{}) (Erc20DepositRecord, error) {
	if args[6] != nil {
		return Erc20DepositRecord{}, args[6].(error)
	}
	log15.Trace("erc20 record", "record", args)
	return Erc20DepositRecord{
			OriginChainTokenAddress:   args[0].(common.Address),
			OriginChainHandlerAddress: args[1].(common.Address),
			DestChainID:               args[2].(*big.Int),
			DestChainHandlerAddress:   args[3].(common.Address),
			DestRecipientAddress:      args[4].(common.Address),
			Amount:                    args[5].(*big.Int),
		},
		nil
}

func UnpackDepositProposal(args ...interface{}) (DepositProposal, error) {
	if args[6].(error) != nil {
		return DepositProposal{}, args[6].(error)
	}
	return DepositProposal{
			OriginChainID: args[0].(*big.Int),
			DepositID:     args[1].(*big.Int),
			DataHash:      args[2].([32]byte),
			NumYes:        args[3].(*big.Int),
			NumNo:         args[4].(*big.Int),
			Status:        args[5].(string),
		},
		nil
}

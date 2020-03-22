// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/contracts/Bridge"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// genericDepositRecord is the return value from the solidity function getGenericDepositRecord()

// depositProposal is the return value from the solidity function getDepositProposal()
type depositProposal struct {
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
	GetDepositProposal(opts *bind.CallOpts, originChainID *big.Int, depositID *big.Int) (Bridge.BridgeReturnDepositProposal, error)
}

type BridgeRaw interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transfer(opts *bind.TransactOpts) (*types.Transaction, error)
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)

	// 	FilterERCTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*bridge.BridgeERCTransferIterator, error)
	// 	FilterGenericTransfer(*bind.FilterOpts, []*big.Int, []*big.Int) (*bridge.BridgeGenericTransferIterator, error)
	// 	FilterNFTTransfer(opts *bind.FilterOpts, _destChain []*big.Int, _depositId []*big.Int) (*bridge.BridgeNFTTransferIterator, error)
}

func UnpackDepositProposal(args ...interface{}) (depositProposal, error) {
	if args[6] != nil {
		return depositProposal{}, args[6].(error)
	}
	return depositProposal{
			OriginChainID: args[0].(*big.Int),
			DepositID:     args[1].(*big.Int),
			DataHash:      args[2].([32]byte),
			NumYes:        args[3].(*big.Int),
			NumNo:         args[4].(*big.Int),
			Status:        args[5].(string),
		},
		nil
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"fmt"

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
	DestChainHandlerAddress   common.Address
	DestChainTokenAddress     common.Address
	DestRecipientAddress      common.Address
	Amount                    *big.Int
}

// depositProposal is the return value from the solidity function getDepositProposal()
type depositProposal struct {
	OriginChainID *big.Int
	DestChainID	  *big.Int
	DepositNonce  *big.Int
	DataHash      [32]byte
}

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
	Deposit(opts *bind.TransactOpts, chainID *big.Int, originChainHandlerAddress common.Address, data []byte)(*types.Transaction, error)
}

type BridgeCaller interface {
	GetDepositProposal(opts *bind.CallOpts, destinationChainID *big.Int,  depositNonce *big.Int) (bridge.BridgeDepositProposal, error)}

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
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) (error)
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

// originChainTokenAddress        := mload(add(data, 0x20))
// destinationChainHandlerAddress := mload(add(data, 0x40))
// destinationChainTokenAddress   := mload(add(data, 0x60))
// destinationRecipientAddress    := mload(add(data, 0x80))
// amount                         := mload(add(data, 0xA0))

func UnpackErc20DepositRecord(args ...interface{}) (erc20DepositRecord, error) {
	fmt.Println("WE ARE HERE NOW")
	if args[1] != nil {
		return erc20DepositRecord{}, args[1].(error)
	}
	// OriginChainTokenAddress        common.Address
	// DestinationChainID             *big.Int
	// DestinationChainHandlerAddress common.Address
	// DestinationChainTokenAddress   common.Address
	// DestinationRecipientAddress    common.Address
	// Depositer                      common.Address
	// Amount                         *big.Int

	return erc20DepositRecord{
			OriginChainTokenAddress:   args[0].(erc20.ERC20HandlerDepositRecord).OriginChainTokenAddress,
			DestChainHandlerAddress:   args[0].(erc20.ERC20HandlerDepositRecord).DestinationChainHandlerAddress,
			DestChainTokenAddress:     args[0].(erc20.ERC20HandlerDepositRecord).DestinationChainTokenAddress,
			DestRecipientAddress:      args[0].(erc20.ERC20HandlerDepositRecord).DestinationRecipientAddress,
			Amount:                    args[0].(erc20.ERC20HandlerDepositRecord).Amount,
		},
		nil
}

//        emit DepositProposalCreated(_chainID, destinationChainID, depositNonce, dataHash);

func UnpackDepositProposal(args ...interface{}) (depositProposal, error) {
	if args[6] != nil {
		return depositProposal{}, args[6].(error)
	}
	return depositProposal{
			OriginChainID: args[0].(*big.Int),
			DestChainID:   args[1].(*big.Int),
			DepositNonce:  args[2].(*big.Int),
			DataHash:      args[3].([32]byte),
		},
		nil
}

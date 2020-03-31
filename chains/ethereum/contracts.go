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
	generic "github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
)

// genericDepositRecord is the return struct from the solidity function getDepositRecord() in the GenericHandler contract
type genericDepositRecord struct {
	DestinationChainID          *big.Int
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	MetaData                    []byte
}

// erc20DepositRecord is the return struct from the solidity function getDepositRecord() in the ERC20Handler contract
type erc20DepositRecord struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	Amount                         *big.Int
}

// erc721DepositRecord is the return value from the solidity function getDepositRecord() in the ERC721Handler contract
type erc721DepositRecord struct {
	OriginChainTokenAddress        common.Address
	DestinationChainID             *big.Int
	DestinationChainHandlerAddress common.Address
	DestinationChainTokenAddress   common.Address
	DestinationRecipientAddress    common.Address
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}

// depositProposal is the return value from the solidity function getDepositProposal()
type depositProposal struct {
	OriginChainID *big.Int
	DestChainID   *big.Int
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
	Deposit(opts *bind.TransactOpts, chainID *big.Int, originChainHandlerAddress common.Address, data []byte) (*types.Transaction, error)
}

type BridgeCaller interface {
	GetDepositProposal(opts *bind.CallOpts, destinationChainID *big.Int, depositNonce *big.Int) (bridge.BridgeDepositProposal, error)
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
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
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

//UnpackErc20DepositRecord unpacks a DepositRecord from an Generic Handler
func UnpackGenericDepositRecord(args ...interface{}) (genericDepositRecord, error) {
	if args[4] != nil {
		return genericDepositRecord{}, args[4].(error)
	}

	depositRecord := args[0].(generic.GenericHandlerDepositRecord)

	return genericDepositRecord{
			DestinationChainID:          depositRecord.DestinationChainID,
			DestinationRecipientAddress: depositRecord.DestinationRecipientAddress,
			Depositer:                   depositRecord.Depositer,
			MetaData:                    depositRecord.MetaData,
		},
		nil
}

//UnpackErc20DepositRecord unpacks a DepositRecord from an Generic Handler
func UnpackERC721DepositRecord(args ...interface{}) (erc721DepositRecord, error) {
	if args[8] != nil {
		return erc721DepositRecord{}, args[8].(error)
	}

	depositRecord := args[0].(erc721.ERC721HandlerDepositRecord)

	return erc721DepositRecord{
			OriginChainTokenAddress:        depositRecord.OriginChainTokenAddress,
			DestinationChainID:             depositRecord.DestinationChainID,
			DestinationChainHandlerAddress: depositRecord.DestinationChainHandlerAddress,
			DestinationChainTokenAddress:   depositRecord.DestinationChainTokenAddress,
			DestinationRecipientAddress:    depositRecord.DestinationRecipientAddress,
			Depositer:                      depositRecord.Depositer,
			TokenID:                        depositRecord.TokenID,
			MetaData:                       depositRecord.MetaData,
		},
		nil
}

//UnpackErc20DepositRecord unpacks a DepositRecord from an ERC20 Handler
func UnpackErc20DepositRecord(args ...interface{}) (erc20DepositRecord, error) {
	if args[1] != nil {
		return erc20DepositRecord{}, args[1].(error)
	}

	depositRecord := args[0].(erc20.ERC20HandlerDepositRecord)

	return erc20DepositRecord{
			OriginChainTokenAddress:        depositRecord.OriginChainTokenAddress,
			DestinationChainID:             depositRecord.DestinationChainID,
			DestinationChainHandlerAddress: depositRecord.DestinationChainHandlerAddress,
			DestinationChainTokenAddress:   depositRecord.DestinationChainTokenAddress,
			DestinationRecipientAddress:    depositRecord.DestinationRecipientAddress,
			Depositer:                      depositRecord.Depositer,
			Amount:                         depositRecord.Amount,
		},
		nil
}

//UnpackDepositProposal unpacks a Deposit Proposal
func UnpackDepositProposal(args ...interface{}) (depositProposal, error) {
	if args[4] != nil {
		return depositProposal{}, args[4].(error)
	}
	return depositProposal{
			OriginChainID: args[0].(*big.Int),
			DestChainID:   args[1].(*big.Int),
			DepositNonce:  args[2].(*big.Int),
			DataHash:      args[3].([32]byte),
		},
		nil
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	// "math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	//DepositAsset                    = "DepositAsset"
	//NftTransfer                     = "NftTransfer"
	//ErcTransfer                     = "ErcTransfer"
	//DepositProposalCreated          = "DepositProposalCreated"
	//DepositedErc20                  = "DepositedErc20"
	Deposit EventSig = "Deposit(uint256,uint256,address,uint256)"
	//DepositAssetSignature           = "DepositAsset(address,bytes32)"
	//NftTransferSignature            = "NFTTransfer(uint256,uint256,address,address,uint256,bytes)"
	//ErcTransferSignature            = "ERCTransfer(uint256,uint256,address,uint256,address)"
	DepositProposalCreated   EventSig = "DepositProposalCreated(uint256,uint256,uint256,bytes32)"
	DepositProposalVote      EventSig = "DepositProposalVote(uint256,uint256,uint256,uint8)"
	DepositProposalFinalized EventSig = "DepositProposalFinalized(uint256,uint256,uint256)"
	DepositProposalExecuted  EventSig = "DepositProposalExecuted(uint256,uint256,uint256)"
)

type evtHandlerFn func(ethtypes.Log) msg.Message

func (l *Listener) handleErc20DepositedEvent(event ethtypes.Log) msg.Message {
	l.log.Debug("Handling deposited event")

	depositNonce := event.Topics[2].Big() // Only item in log is indexed.

	deposit, err := UnpackErc20DepositRecord(l.erc20HandlerContract.ERC20HandlerCaller.GetDepositRecord(&bind.CallOpts{}, depositNonce))

	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "err", err)
	}

	return msg.Message{
		Type:         msg.FungibleTransfer,
		Source:       l.cfg.id,
		Destination:  msg.ChainId(deposit.DestinationChainID.Uint64()),
		DepositNonce: uint32(depositNonce.Uint64()),
		Metadata: []interface{}{
			deposit.DestinationRecipientAddress,
			deposit.Amount.Bytes(),
			deposit.TokenId,
		},
	}
}

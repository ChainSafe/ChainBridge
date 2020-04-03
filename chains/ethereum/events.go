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
	Deposit                  EventSig = "Deposit(uint256,uint256,address,uint256)"
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

	return msg.NewFungibleTransfer(
		l.cfg.id,
		msg.ChainId(deposit.DestinationChainID.Uint64()),
		uint32(depositNonce.Uint64()),
		deposit.Amount,
		deposit.TokenId,
		deposit.DestinationRecipientAddress,
	)
}

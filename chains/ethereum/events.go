// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	// "math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type evtHandlerFn func(ethtypes.Log) msg.Message

func (l *Listener) handleErc20DepositedEvent(destId msg.ChainId, nonce msg.Nonce) msg.Message {
	l.log.Debug("Handling deposited event")

	record, err := l.erc20HandlerContract.GetDepositRecord(&bind.CallOpts{}, nonce.Big())
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "err", err)
	}

	return msg.NewFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.Amount,
		record.ResourceID,
		record.DestinationRecipientAddress,
	)
}

func (l *Listener) handleErc721DepositedEvent(destId msg.ChainId, nonce msg.Nonce) msg.Message {
	l.log.Debug("Handling deposited event")

	record, err := l.erc721HandlerContract.GetDepositRecord(&bind.CallOpts{}, nonce.Big())
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "err", err)
	}

	return msg.NewNonFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.TokenID.Bytes(),
		record.DestinationRecipientAddress,
		record.MetaData,
	)
}

func (l *Listener) handleGenericDepositedEvent(destId msg.ChainId, nonce msg.Nonce) msg.Message {
	l.log.Debug("Handling deposited event")

	record, err := l.genericHandlerContract.GetDepositRecord(&bind.CallOpts{}, nonce.Big())
	if err != nil {
		l.log.Error("Error Unpacking Generic Deposit Record", "err", err)
	}

	return msg.NewGenericTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.MetaDataHash[:],
	)
}

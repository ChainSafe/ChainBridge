// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (l *listener) handleErc20DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Debug("Handling deposited event", "dest", destId, "nonce", nonce)

	record, err := l.erc20HandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.kp.CommonAddress()}, nonce.Big(), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "err", err)
		return msg.Message{}, err
	}

	return msg.NewFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.Amount,
		record.ResourceID,
		record.DestinationRecipientAddress,
	), nil
}

func (l *listener) handleErc721DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Debug("Handling deposited event")

	record, err := l.erc721HandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.kp.CommonAddress()}, nonce.Big(), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "err", err)
		return msg.Message{}, err
	}

	recipient := record.DestinationRecipientAddress[:record.LenDestinationRecipientAddress.Int64()]

	return msg.NewNonFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.TokenID,
		recipient,
		record.MetaData,
	), nil
}

func (l *listener) handleGenericDepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Debug("Handling deposited event")

	record, err := l.genericHandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.kp.CommonAddress()}, nonce.Big(), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking Generic Deposit Record", "err", err)
		return msg.Message{}, nil
	}

	return msg.NewGenericTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.MetaData[:],
	), nil
}

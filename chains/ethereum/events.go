// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (l *listener) handleErc20DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	if l.cfg.erc20HandlerContract == utils.ZeroAddress {
		l.log.Warn("You are attempting to perform an ERC20 transfer without a deployed handler")
	}

	l.log.Info("Handling fungible deposit event", "dest", destId, "nonce", nonce)

	record, err := l.erc20HandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
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
	if l.cfg.erc20HandlerContract == utils.ZeroAddress {
		l.log.Warn("You are attempting to perform an ERC721 transfer without a deployed handler")
	}

	l.log.Info("Handling nonfungible deposit event")

	record, err := l.erc721HandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC721 Deposit Record", "err", err)
		return msg.Message{}, err
	}

	return msg.NewNonFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.TokenID,
		record.DestinationRecipientAddress,
		record.MetaData,
	), nil
}

func (l *listener) handleGenericDepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	if l.cfg.erc20HandlerContract == utils.ZeroAddress {
		l.log.Warn("You are attempting to perform a generic transfer without a deployed handler")
	}

	l.log.Info("Handling generic deposit event")

	record, err := l.genericHandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
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

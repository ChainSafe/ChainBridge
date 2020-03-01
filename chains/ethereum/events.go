// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"strings"

	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (l *Listener) handleNftTransferEvent(eventI interface{}) msg.Message {
	log15.Debug("Handling deposit proposal event")
	event := eventI.(ethtypes.Log)

	contractAbi, err := abi.JSON(strings.NewReader(emitter.EmitterABI))
	if err != nil {
		log15.Error("Unable to decode event", err)
	}

	var nftEvent emitter.EmitterNFTTransfer
	err = contractAbi.Unpack(&nftEvent, "NFTTransfer", event.Data)
	if err != nil {
		log15.Error("Unable to unpack NFTTransfer", err)
	}

	// Capture indexed values
	nftEvent.DestChain = event.Topics[1].Big()
	nftEvent.DepositId = event.Topics[2].Big()

	msg := msg.Message{
		Type:        msg.CreateDepositProposalType,
		Source:      l.cfg.id,
		Destination: msg.ChainId(uint8(nftEvent.DestChain.Uint64())),
		// TODO: Can we safely downsize?
		DepositId: uint32(nftEvent.DepositId.Uint64()),
		To:        nftEvent.To.Bytes(),
		Metadata:  nftEvent.Data,
	}

	return msg
}

func (l *Listener) handleErcTransferEvent(eventI interface{}) msg.Message {
	log15.Debug("Handling deposit proposal event")
	event := eventI.(ethtypes.Log)

	contractAbi, err := abi.JSON(strings.NewReader(emitter.EmitterABI))
	if err != nil {
		log15.Error("Unable to decode event", err)
	}

	var ercEvent emitter.EmitterERCTransfer
	err = contractAbi.Unpack(&ercEvent, "ERCTransfer", event.Data)
	if err != nil {
		log15.Error("Unable to unpack ERCTransfer", err)
	}

	// Capture indexed values
	ercEvent.DestChain = event.Topics[1].Big()
	ercEvent.DepositId = event.Topics[2].Big()

	msg := msg.Message{
		Type:        msg.CreateDepositProposalType,
		Source:      l.cfg.id,
		Destination: msg.ChainId(uint8(ercEvent.DestChain.Uint64())),
		// TODO: Can we safely downsize?
		DepositId: uint32(ercEvent.DepositId.Uint64()),
		To:        ercEvent.To.Bytes(),
		// Metadata:  ercEvent.Data,
	}

	return msg
}

func (l *Listener) handleVoteEvent(eventI interface{}) msg.Message {
	log15.Debug("Handling vote event")
	event := eventI.(ethtypes.Log)

	contractAbi, err := abi.JSON(strings.NewReader(string(receiver.ReceiverABI)))
	if err != nil {
		log15.Error("Unable to decode event", err)
	}

	var depositEvent receiver.ReceiverDepositProposalCreated
	err = contractAbi.Unpack(&depositEvent, "DepositProposalCreated", event.Data)
	if err != nil {
		log15.Error("Unable to unpack DepositProposalCreated", err)
	}

	return msg.Message{
		Source:      msg.ChainId(uint8(depositEvent.OriginChain.Uint64())), // Todo handle safely
		Destination: l.cfg.id,                                              // We are reading from the receiver, must write to the same contract
		Type:        msg.VoteDepositProposalType,
		DepositId:   uint32(depositEvent.DepositId.Int64()),
		Metadata:    depositEvent.Hash[:],
	}
}

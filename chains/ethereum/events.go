// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	DepositAsset           = "DepositAsset"
	NftTransfer            = "NftTransfer"
	ErcTransfer            = "ErcTransfer"
	DepositProposalCreated = "DepositProposalCreated"
	DepositedErc20         = "DepositedErc20"

	DepositedErc20Signature         = "ERC20Deposited(uint256)"
	DepositAssetSignature           = "DepositAsset(address,bytes32)"
	NftTransferSignature            = "NFTTransfer(uint256,uint256,address,address,uint256,bytes)"
	ErcTransferSignature            = "ERCTransfer(uint256,uint256,address,uint256,address)"
	DepositProposalCreatedSignature = "DepositProposalCreated(uint256,uint256,bytes32)"
)

func (l *Listener) handleErc20DepositedEvent(eventI interface{}) msg.Message {
	log15.Debug("Handling deposited event")
	event := eventI.(ethtypes.Log)

	depositID := event.Topics[1].Big() // Only item in log is indexed.

	// TODO remove when issue addressed https://github.com/ChainSafe/ChainBridgeV2/issues/173
	var destID msg.ChainId
	if l.cfg.id == 0 {
		destID = msg.ChainId(1)
	} else {
		destID = msg.ChainId(0)
	}

	deposit, err := UnpackErc20DepositRecord(l.bridgeContract.BridgeCaller.GetERC20DepositRecord(&bind.CallOpts{}, destID.Big(), depositID))
	if err != nil {
		log15.Error("Error Unpacking ERC20 Deposit Record", "err", err)
	}

	return msg.Message{
		Type:        msg.CreateDepositProposalType,
		Source:      l.cfg.id,
		Destination: msg.ChainId(deposit.DestChainID.Uint64()),
		DepositId:   uint32(depositID.Uint64()),
		To:          deposit.DestChainHandlerAddress.Bytes(),
		Metadata:    deposit.Amount.Bytes(),
	}
}

func (l *Listener) handleVoteEvent(eventI interface{}) msg.Message {
	log15.Debug("Handling vote event")
	event := eventI.(ethtypes.Log)

	originChainID := event.Topics[1].Big()
	depositID := event.Topics[2].Big()
	hash := event.Topics[3]

	return msg.Message{
		Source:      msg.ChainId(uint8(originChainID.Uint64())), // Todo handle safely
		Destination: l.cfg.id,                                   // Must write to the same contract
		Type:        msg.VoteDepositProposalType,
		DepositId:   uint32(depositID.Int64()),
		Metadata:    hash[:],
	}
}

// func (l *Listener) handleErcTransferEvent(eventI interface{}) msg.Message {
// 	log15.Debug("Handling deposit proposal event")
// 	event := eventI.(ethtypes.Log)

// 	contractAbi, err := abi.JSON(strings.NewReader(emitter.EmitterABI))
// 	if err != nil {
// 		log15.Error("Unable to decode event", err)
// 	}

// 	var ercEvent emitter.EmitterERCTransfer
// 	err = contractAbi.Unpack(&ercEvent, "ERCTransfer", event.Data)
// 	if err != nil {
// 		log15.Error("Unable to unpack ERCTransfer", err)
// 	}

// 	// Capture indexed values
// 	ercEvent.DestChain = event.Topics[1].Big()
// 	ercEvent.DepositId = event.Topics[2].Big()

// 	return msg.Message{
// 		Type:        msg.CreateDepositProposalType,
// 		Source:      l.cfg.id,
// 		Destination: msg.ChainId(uint8(ercEvent.DestChain.Uint64())),
// 		// TODO: Can we safely downsize?
// 		DepositId: uint32(ercEvent.DepositId.Uint64()),
// 		To:        ercEvent.To.Bytes(),
// 		// Metadata:  ercEvent.Data,
// 	}
// }

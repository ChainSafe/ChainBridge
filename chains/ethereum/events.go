// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

const (
	DepositAsset = "DepositAsset"
	NftTransfer  = "NftTransfer"
	ErcTransfer  = "ErcTransfer"

	DepositAssetSignature = "DepositAsset(address,bytes32)"
	NftTransferSignature  = "NFTTransfer(uint256,uint256,address,address,uint256,bytes)"
	ErcTransferSignature  = "ERCTransfer(uint256,uint256,address,uint256,address)"
)

func (l *Listener) handleTransferEvent(eventI interface{}) msg.Message {
	log15.Debug("Handling deposit proposal event")
	event := eventI.(ethtypes.Log)

	depositID := event.Topics[1].Big() // Only item in log is indexed.
	deposit, err := UnpackGenericDepositRecord(l.bridgeContract.BridgeCaller.GetGenericDepositRecord(&bind.CallOpts{}, l.cfg.id.Big(), depositID))
	if err != nil {
		log15.Error("Unable to decode event", err)
	}

	msg := msg.Message{
		Type:        msg.CreateDepositProposalType,
		Source:      l.cfg.id,
		Destination: msg.ChainId(deposit.DestChainID.Uint64()),
		DepositId:   uint32(depositID.Uint64()),
		To:          deposit.DestChainHandlerAddress.Bytes(),
		Metadata:    deposit.Data,
	}

	return msg
}

func (l *Listener) handleTestDeposit(eventI interface{}) msg.Message {
	event := eventI.(ethtypes.Log)
	data := ethcrypto.Keccak256Hash(event.Topics[0].Bytes()).Bytes()
	return msg.Message{
		Type:     msg.DepositAssetType,
		Metadata: data,
	}
}

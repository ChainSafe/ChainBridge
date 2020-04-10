// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package msg

import (
	"math/big"
)

//type MessageType string

//var DepositAssetType MessageType = "substrate_deposit_asset"
//var CreateDepositProposalType MessageType = "create_deposit_proposal"
//var VoteDepositProposalType MessageType = "vote_deposit_proposal"
//var ExecuteDepositType MessageType = "execute_deposit"

type TransferType string
type Nonce uint32
type ResourceId [32]byte

var FungibleTransfer TransferType = "FungibleTransfer"
var NonFungibleTransfer TransferType = "NonFungibleTransfer"
var GenericTransfer TransferType = "GenericTransfer"

// Message is used as a generic format to communicate between chains
type Message struct {
	Source       ChainId      // Source where message was initiated
	Destination  ChainId      // Destination chain of message
	Type         TransferType // type of bridge transfer
	DepositNonce Nonce        // Nonce for the deposit
	ResourceId   ResourceId
	Payload      []interface{} // data associated with event sequence
}

func NewFungibleTransfer(source, dest ChainId, nonce Nonce, amount *big.Int, resourceId ResourceId, recipient []byte) Message {
	return Message{
		Source:       source,
		Destination:  dest,
		Type:         FungibleTransfer,
		DepositNonce: nonce,
		ResourceId:   resourceId,
		Payload: []interface{}{
			amount.Bytes(),
			recipient,
		},
	}
}

func NewNonFungibleTransfer(source, dest ChainId, nonce Nonce, resourceId ResourceId, tokenId, recipient, metadata []byte) Message {
	return Message{
		Source:       source,
		Destination:  dest,
		Type:         NonFungibleTransfer,
		DepositNonce: nonce,
		Payload: []interface{}{
			resourceId,
			tokenId,
			recipient,
			metadata,
		},
	}
}

func NewGenericTransfer(source, dest ChainId, nonce Nonce, resourceId ResourceId, metadata []byte) Message {
	return Message{
		Source:       source,
		Destination:  dest,
		Type:         GenericTransfer,
		DepositNonce: nonce,
		ResourceId:   resourceId,
		Payload: []interface{}{
			metadata,
		},
	}
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package msg

//type MessageType string

//var DepositAssetType MessageType = "substrate_deposit_asset"
//var CreateDepositProposalType MessageType = "create_deposit_proposal"
//var VoteDepositProposalType MessageType = "vote_deposit_proposal"
//var ExecuteDepositType MessageType = "execute_deposit"

type TransferType string

var FungibleTransfer TransferType = "FungibleTransfer"
var NonFungibleTransfer TransferType = "NonFungibleTransfer"
var GenericTransfer TransferType = "GenericTransfer"

// Message is used as a generic format to communicate between chains
type Message struct {
	Source       ChainId       // Source where message was initiated
	Destination  ChainId       // Destination chain of message
	Type         TransferType  // type of bridge transfer
	DepositNonce uint32        // Nonce for the deposit
	Metadata     []interface{} // data associated with event sequence
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type eventName string
type eventHandler func(interface{}) msg.Message

const AssetTx eventName = "AssetTx"
const ValidatorAdded eventName = "ValidatorAdded"
const ValidatorRemoved eventName = "ValidatorRemoved"
const VoteFor eventName = "VoteFor"
const VoteAgainst eventName = "VoteAgainst"
const ProposalSucceeded eventName = "ProposalSucceeded"
const ProposalFailed eventName = "ProposalFailed"

var Subscriptions = []eventName{AssetTx}

type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

type EventAssetTransfer struct {
	Phase        types.Phase
	Destination  types.Bytes
	DepositNonce types.U32
	To           types.Bytes
	TokenID      types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventValidatorAdded struct {
	Phase     types.Phase
	Valdiator types.Hash
	Topics    []types.Hash
}

type EventValidatorRemoved struct {
	Phase     types.Phase
	Validator types.Hash
	Topics    []types.Hash
}

type EventVoteFor struct {
	Phase  types.Phase
	Hash   types.Hash
	Voter  types.AccountID
	Topics []types.Hash
}

type EventVoteAgainst struct {
	Phase  types.Phase
	Hash   types.Hash
	Voter  types.AccountID
	Topics []types.Hash
}

type EventProposalSucceeded struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type EventProposalFailed struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type Events struct {
	types.EventRecords
	Bridge_AssetTransfer     []EventAssetTransfer     //nolint:stylecheck,golint
	Bridge_ValidatorAdded    []EventValidatorAdded    //nolint:stylecheck,golint
	Bridge_ValidatorRemoved  []EventValidatorRemoved  //nolint:stylecheck,golint
	Bridge_VoteFor           []EventVoteFor           //nolint:stylecheck,golint
	Bridge_VoteAgainst       []EventVoteAgainst       //nolint:stylecheck,golint
	Bridge_ProposalSucceeded []EventProposalSucceeded //nolint:stylecheck,golint
	Bridge_ProposalFailed    []EventProposalFailed    //nolint:stylecheck,golint
}

func nftHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventNFTDeposited)
	if !ok {
		log15.Error("failed to cast EventNFTDeposited type")
	}

	log15.Info("Got nfts event!", "evt", evt.Asset.Hex())

	return msg.Message{
		Source:      msg.CentrifugeId,
		Destination: msg.EthereumId,
		Type:        msg.DepositAssetType,
		Metadata:    evt.Asset[:],
	}
}

func assetTransferHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventAssetTransfer)
	if !ok {
		log15.Error("failed to cast EventAssetTransfer type")
	}

	log15.Info("Got asset transfer event!", "destination", evt.Destination)

	return msg.Message{
		Source:      msg.CentrifugeId,
		Destination: msg.EthereumId,
		Type:        msg.DepositAssetType,
		Metadata:    evt.Destination[:], // TODO: Pack data
	}
}

func validatorAddedHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventValidatorAdded)
	if !ok {
		log15.Error("failed to cast EventValidatorAdded type")
	}

	log15.Info("Got ValidatorAdded event!", "address", evt.Valdiator)

	return msg.Message{}
}

func validatorRemovedHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventValidatorRemoved)
	if !ok {
		log15.Error("failed to cast EventValidatorRemoved type")
	}

	log15.Info("Got ValidatorRemoved event!", "address", evt.Validator)

	return msg.Message{}
}

func voteForHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventVoteFor)
	if !ok {
		log15.Error("failed to cast EventVoteFor type")
	}

	log15.Info("Got VoteFor event!", "hash", evt.Hash)

	return msg.Message{}
}

func voteAgainstHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventVoteAgainst)
	if !ok {
		log15.Error("failed to cast EventVoteAgainst type")
	}

	log15.Info("Got VoteAgainst event!", "hash", evt.Hash)

	return msg.Message{}
}

func proposalSucceededHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventProposalSucceeded)
	if !ok {
		log15.Error("failed to cast EventProposalSucceeded type")
	}

	log15.Info("Got ProposalSucceeded event!", "hash", evt.Hash)

	return msg.Message{}
}

func proposalFailedHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventProposalFailed)
	if !ok {
		log15.Error("failed to cast EventProposalFailed type")
	}

	log15.Info("Got ProposalFailed event!", "hash", evt.Hash)

	return msg.Message{}
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"encoding/binary"

	msg "github.com/ChainSafe/ChainBridge/message"

	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type eventName string
type eventHandler func(interface{}) msg.Message

const RelayerThresholdSet eventName = "RelayerThresholdSet"
const ChainWhitelisted eventName = "ChainWhitelsited"
const AssetTransfer eventName = "AssetTransfer"
const RelayerAdded eventName = "RelayerAdded"
const RelayerRemoved eventName = "RelayerRemoved"
const VoteFor eventName = "VoteFor"
const VoteAgainst eventName = "VoteAgainst"
const ProposalApproved eventName = "ProposalApproved"
const ProposalRejected eventName = "ProposalRejected"
const ProposalSucceeded eventName = "ProposalSucceeded"
const ProposalFailed eventName = "ProposalFailed"

var Subscriptions = []struct {
	name    eventName
	handler eventHandler
}{
	{RelayerThresholdSet, nil},
	{ChainWhitelisted, nil},
	{RelayerAdded, nil},
	{RelayerRemoved, nil},
	{AssetTransfer, assetTransferHandler},
	{VoteFor, nil},
	{VoteAgainst, nil},
	{ProposalApproved, nil},
	{ProposalRejected, nil},
	{ProposalSucceeded, nil},
	{ProposalFailed, nil},
}

type EventRelayerThresholdSet struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

type EventChainWhitelisted struct {
	Phase   types.Phase
	ChainId types.U32
	Topics  []types.Hash
}

type EventRelayerAdded struct {
	Phase   types.Phase
	Relayer types.Hash
	Topics  []types.Hash
}

type EventRelayerRemoved struct {
	Phase   types.Phase
	Relayer types.Hash
	Topics  []types.Hash
}

type EventAssetTransfer struct {
	Phase        types.Phase
	Destination  types.U32
	DepositNonce types.U32
	To           types.Bytes
	TokenID      types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventVoteFor struct {
	Phase        types.Phase
	DepositNonce types.U32
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventVoteAgainst struct {
	Phase        types.Phase
	DepositNonce types.U32
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventProposalApproved struct {
	Phase        types.Phase
	DepositNonce types.U32
	Topics       []types.Hash
}

type EventProposalRejected struct {
	Phase        types.Phase
	DepositNonce types.U32
	Topics       []types.Hash
}

type EventProposalSucceeded struct {
	Phase        types.Phase
	DepositNonce types.U32
	Topics       []types.Hash
}

type EventProposalFailed struct {
	Phase        types.Phase
	DepositNonce types.U32
	Topics       []types.Hash
}

type Events struct {
	types.EventRecords
	Bridge_RelayerThresholdSet []EventRelayerThresholdSet //nolint:stylecheck,golint
	Bridge_ChainWhitelisted    []EventChainWhitelisted    //nolint:stylecheck,golint
	Bridge_RelayerAdded        []EventRelayerAdded        //nolint:stylecheck,golint
	Bridge_RelayerRemoved      []EventRelayerRemoved      //nolint:stylecheck,golint
	Bridge_AssetTransfer       []EventAssetTransfer       //nolint:stylecheck,golint
	Bridge_VoteFor             []EventVoteFor             //nolint:stylecheck,golint
	Bridge_VoteAgainst         []EventVoteAgainst         //nolint:stylecheck,golint
	Bridge_ProposalApproved    []EventProposalApproved    //nolint:stylecheck,golint
	Bridge_ProposalRejected    []EventProposalRejected    //nolint:stylecheck,golint
	Bridge_ProposalSucceeded   []EventProposalSucceeded   //nolint:stylecheck,golint
	Bridge_ProposalFailed      []EventProposalFailed      //nolint:stylecheck,golint
	Sudo_Sudid                 []EventSudid               //nolint:stylecheck,golint
}

func assetTransferHandler(evtI interface{}) msg.Message {
	evt, ok := evtI.(EventAssetTransfer)
	if !ok {
		log15.Error("failed to cast EventAssetTransfer type")
	}

	log15.Info("Got asset transfer event!", "destination", evt.Destination)

	meta := make([]byte, 8)
	binary.LittleEndian.PutUint32(meta, uint32(evt.Destination))

	return msg.Message{
		Destination:  msg.ChainId(evt.Destination),
		Type:         msg.CreateDepositProposalType,
		DepositNonce: uint32(evt.DepositNonce),
		// TODO: What should To actually be?
		To:       evt.To,
		Metadata: meta,
	}
}

// TODO: These should be added directly to GSRPC

type EventSudid struct {
	Phase   types.Phase
	Success types.Bool
	Topics  []types.Hash
}

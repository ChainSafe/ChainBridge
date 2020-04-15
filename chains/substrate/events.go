// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"

	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type eventName string
type eventHandler func(interface{}, log15.Logger) (msg.Message, error)

const RelayerThresholdChanged eventName = "RelayerThresholdChanged"
const ChainWhitelisted eventName = "ChainWhitelsited"
const RelayerAdded eventName = "RelayerAdded"
const RelayerRemoved eventName = "RelayerRemoved"

const FungibleTransfer eventName = "FungibleTransfer"
const NonFungibleTransfer eventName = "NonFungibleTransfer"
const GenericTransfer eventName = "GenericTransfer"

const VoteFor eventName = "VoteFor"
const VoteAgainst eventName = "VoteAgainst"

const ProposalApproved eventName = "ProposalApproved"
const ProposalRejected eventName = "ProposalRejected"
const ProposalSucceeded eventName = "ProposalSucceeded"
const ProposalFailed eventName = "ProposalFailed"

const CodeUpdated eventName = "CodeUpdated"

var Subscriptions = []struct {
	name    eventName
	handler eventHandler
}{
	{RelayerThresholdChanged, nil},
	{ChainWhitelisted, nil},
	{RelayerAdded, nil},
	{RelayerRemoved, nil},
	{FungibleTransfer, fungibleTransferHandler},
	{NonFungibleTransfer, nonFungibleTransferHandler},
	{GenericTransfer, genericTransferHandler},
	{VoteFor, nil},
	{VoteAgainst, nil},
	{ProposalApproved, nil},
	{ProposalRejected, nil},
	{ProposalSucceeded, nil},
	{ProposalFailed, nil},
}

type EventRelayerThresholdChanged struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

type EventChainWhitelisted struct {
	Phase   types.Phase
	ChainId types.U8
	Topics  []types.Hash
}

type EventRelayerAdded struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventRelayerRemoved struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Amount       types.U32
	Recipient    types.Bytes
	Topics       []types.Hash
}

type EventNonFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	TokenId      types.Bytes
	Recipient    types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventGenericTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventVoteFor struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventVoteAgainst struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventProposalApproved struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalRejected struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalSucceeded struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalFailed struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventCodeUpdated struct {
	Phase  types.Phase
	Topics []types.Hash
}

// TODO: This should be added directly to GSRPC
type EventSudid struct {
	Phase   types.Phase
	Success types.Bool
	Topics  []types.Hash
}

type Events struct {
	types.EventRecords
	Bridge_RelayerThresholdChanged []EventRelayerThresholdChanged //nolint:stylecheck,golint
	Bridge_ChainWhitelisted        []EventChainWhitelisted        //nolint:stylecheck,golint
	Bridge_RelayerAdded            []EventRelayerAdded            //nolint:stylecheck,golint
	Bridge_RelayerRemoved          []EventRelayerRemoved          //nolint:stylecheck,golint
	Bridge_FungibleTransfer        []EventFungibleTransfer        //nolint:stylecheck,golint
	Bridge_NonFungibleTransfer     []EventNonFungibleTransfer     //nolint:stylecheck,golint
	Bridge_GenericTransfer         []EventGenericTransfer         //nolint:stylecheck,golint
	Bridge_VoteFor                 []EventVoteFor                 //nolint:stylecheck,golint
	Bridge_VoteAgainst             []EventVoteAgainst             //nolint:stylecheck,golint
	Bridge_ProposalApproved        []EventProposalApproved        //nolint:stylecheck,golint
	Bridge_ProposalRejected        []EventProposalRejected        //nolint:stylecheck,golint
	Bridge_ProposalSucceeded       []EventProposalSucceeded       //nolint:stylecheck,golint
	Bridge_ProposalFailed          []EventProposalFailed          //nolint:stylecheck,golint
	System_CodeUpdated             []EventCodeUpdated             //nolint:stylecheck,golint
	Sudo_Sudid                     []EventSudid                   //nolint:stylecheck,golint
}

func fungibleTransferHandler(evtI interface{}, log log15.Logger) (msg.Message, error) {
	evt, ok := evtI.(EventFungibleTransfer)
	if !ok {
		return msg.Message{}, fmt.Errorf("failed to cast EventFungibleTransfer type")
	}

	resourceId := msg.ResourceId(evt.ResourceId)
	log.Info("Got fungible transfer event!", "destination", evt.Destination, "resourceId", resourceId.Hex(), "amount", evt.Amount)

	return msg.NewFungibleTransfer(
		0, // Unset
		msg.ChainId(evt.Destination),
		msg.Nonce(evt.DepositNonce),
		big.NewInt(int64(evt.Amount)),
		resourceId,
		evt.Recipient,
	), nil
}

func nonFungibleTransferHandler(evtI interface{}, log log15.Logger) (msg.Message, error) {
	evt, ok := evtI.(EventNonFungibleTransfer)
	if !ok {
		return msg.Message{}, fmt.Errorf("failed to cast EventNonFungibleTransfer type")
	}

	log.Info("Got non-fungible transfer event!", "destination", evt.Destination, "resourceId", evt.ResourceId)

	return msg.NewNonFungibleTransfer(
		0, // Unset
		msg.ChainId(evt.Destination),
		msg.Nonce(evt.DepositNonce),
		msg.ResourceId(evt.ResourceId),
		evt.TokenId,
		evt.Recipient,
		evt.Metadata,
	), nil
}

func genericTransferHandler(evtI interface{}, log log15.Logger) (msg.Message, error) {
	evt, ok := evtI.(EventGenericTransfer)
	if !ok {
		return msg.Message{}, fmt.Errorf("failed to cast EventGenericTransfer type")
	}

	log.Info("Got generic transfer event!", "destination", evt.Destination, "resourceId", evt.ResourceId)

	return msg.NewGenericTransfer(
		0, // Unset
		msg.ChainId(evt.Destination),
		msg.Nonce(evt.DepositNonce),
		msg.ResourceId(evt.ResourceId),
		evt.Metadata,
	), nil
}

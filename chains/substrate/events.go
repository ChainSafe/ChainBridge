// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/log15"

	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type eventName string
type eventHandler func(interface{}, log15.Logger) (msg.Message, error)

const FungibleTransfer eventName = "FungibleTransfer"
const NonFungibleTransfer eventName = "NonFungibleTransfer"
const GenericTransfer eventName = "GenericTransfer"

var Subscriptions = []struct {
	name    eventName
	handler eventHandler
}{
	{FungibleTransfer, fungibleTransferHandler},
	{NonFungibleTransfer, nonFungibleTransferHandler},
	{GenericTransfer, genericTransferHandler},
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

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.BlockNumber
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventCodeUpdated struct {
	Phase  types.Phase
	Topics []types.Hash
}

type Events struct {
	utils.Events
	ChainBridge_FungibleTransfer    []EventFungibleTransfer    //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer []EventNonFungibleTransfer //nolint:stylecheck,golint
	ChainBridge_GenericTransfer     []EventGenericTransfer     //nolint:stylecheck,golint
	System_CodeUpdated         []EventCodeUpdated         //nolint:stylecheck,golint
	Nfts_DepositAsset                []EventNFTDeposited                   //nolint:stylecheck,golint
	Council_Proposed                 []types.EventCollectiveProposed       //nolint:stylecheck,golint
	Council_Voted                    []types.EventCollectiveProposed       //nolint:stylecheck,golint
	Council_Approved                 []types.EventCollectiveApproved       //nolint:stylecheck,golint
	Council_Disapproved              []types.EventCollectiveDisapproved    //nolint:stylecheck,golint
	Council_Executed                 []types.EventCollectiveExecuted       //nolint:stylecheck,golint
	Council_MemberExecuted           []types.EventCollectiveMemberExecuted //nolint:stylecheck,golint
	Council_Closed                   []types.EventCollectiveClosed         //nolint:stylecheck,golint
	Fees_FeeChanged                  []EventFeeChanged                     //nolint:stylecheck,golint
	MultiAccount_NewMultiAccount     []EventNewMultiAccount                //nolint:stylecheck,golint
	MultiAccount_MultiAccountUpdated []EventMultiAccountUpdated            //nolint:stylecheck,golint
	MultiAccount_MultiAccountRemoved []EventMultiAccountRemoved            //nolint:stylecheck,golint
	MultiAccount_NewMultisig         []EventNewMultisig                    //nolint:stylecheck,golint
	MultiAccount_MultisigApproval    []EventMultisigApproval               //nolint:stylecheck,golint
	MultiAccount_MultisigExecuted    []EventMultisigExecuted               //nolint:stylecheck,golint
	MultiAccount_MultisigCancelled   []EventMultisigCancelled              //nolint:stylecheck,golint
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
		big.NewInt(0).SetBytes(evt.TokenId[:]),
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

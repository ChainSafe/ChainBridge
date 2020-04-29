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
	Amount       types.U256
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

type Events struct {
	utils.Events
	ChainBridge_FungibleTransfer    []EventFungibleTransfer    //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer []EventNonFungibleTransfer //nolint:stylecheck,golint
	ChainBridge_GenericTransfer     []EventGenericTransfer     //nolint:stylecheck,golint
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
		evt.Amount.Int,
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

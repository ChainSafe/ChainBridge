// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

// accoutData is the on-chain representation of an account
type AccountData struct {
	Nonce    types.U32
	Refcount types.UCompact
	Data     struct {
		Free       types.U128
		Reserved   types.U128
		MiscFrozen types.U128
		FreeFrozen types.U128
	}
}

type voteState struct {
	VotesFor     []types.AccountID
	VotesAgainst []types.AccountID
}

// proposal represents an on-chain proposal
type proposal struct {
	DepositNonce types.U32
	Call         types.Call
	sourceId     types.U8
}

// encode takes only nonce and call and encodes them for storage queries
func (p *proposal) encode() ([]byte, error) {
	return types.EncodeToBytes(struct {
		types.U32
		types.Call
	}{p.DepositNonce, p.Call})
}

func (w *Writer) createFungibleProposal(m msg.Message) (*proposal, error) {
	amount64 := big.NewInt(0).SetBytes(m.Payload[0].([]byte)).Uint64()
	amount := types.U32(uint32(amount64))
	recipient := types.NewAccountID(m.Payload[1].([]byte))
	depositNonce := types.U32(m.DepositNonce)

	meta := w.conn.getMetadata()
	method, err := w.resolveResourceId(m.ResourceId)
	if err != nil {
		return nil, err
	}

	call, err := types.NewCall(
		&meta,
		method,
		recipient,
		amount,
	)
	if err != nil {
		return nil, err
	}

	return &proposal{
		DepositNonce: depositNonce,
		Call:         call,
		sourceId:     types.U8(m.Source),
	}, nil
}

func (w *Writer) createGenericProposal(m msg.Message) (*proposal, error) {
	meta := w.conn.getMetadata()
	method, err := w.resolveResourceId(m.ResourceId)

	call, err := types.NewCall(
		&meta,
		method,
		m.Payload[0].([]byte),
	)

	if err != nil {
		return nil, err
	}
	return &proposal{
		DepositNonce: types.U32(m.DepositNonce),
		Call:         call,
		sourceId:     types.U8(m.Source),
	}, nil
}

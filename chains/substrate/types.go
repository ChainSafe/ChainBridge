// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"encoding/binary"
	"fmt"
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
}

func createFungibleProposal(m msg.Message, meta *types.Metadata) (*proposal, error) {
	amount64 := big.NewInt(0).SetBytes(m.Metadata[0].([]byte)).Uint64()
	amount := types.U32(uint32(amount64))
	recipient := types.NewAccountID(m.Metadata[2].([]byte))
	depositNonce := types.U32(m.DepositNonce)

	call, err := types.NewCall(
		meta,
		ExampleTransfer.String(),
		recipient,
		amount,
	)

	if err != nil {
		return nil, err
	}
	return &proposal{
		DepositNonce: depositNonce,
		Call:         call,
	}, nil
}

func createGenericProposal(m msg.Message, meta *types.Metadata) (*proposal, error) {
	// The token id field is repurposed here to indicate some generic action.
	// It consists of a chain id and a method id
	_ = binary.LittleEndian.Uint32(m.Metadata[0].([]byte)[:4])
	methodId := binary.LittleEndian.Uint32(m.Metadata[0].([]byte)[:4])
	depositNonce := types.U32(m.DepositNonce)

	method, err := getMethod(methodId)
	if err != nil {
		return nil, err
	}

	call, err := types.NewCall(
		meta,
		method.String(),
		m.Metadata[0].([]byte),
	)

	if err != nil {
		return nil, err
	}
	return &proposal{
		DepositNonce: depositNonce,
		Call:         call,
	}, nil
}

func getMethod(id uint32) (Method, error) {
	switch id {
	case 1:
		return ExampleRemark, nil
	default:
		return "", fmt.Errorf("unknown method id for generic Call: %d", id)
	}
}

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"encoding/binary"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type accountData struct {
	Nonce    types.U32
	Refcount types.UCompact
	Data     struct {
		Free       types.U128
		Reserved   types.U128
		MiscFrozen types.U128
		FreeFrozen types.U128
	}
}

// assetTxProposal represents an on-chain assetTxProposal
type assetTxProposal struct {
	depositNonce types.U32
	recipient    types.AccountID
	amount       types.U32
	call         types.Call
}

type proposalKey struct {
	DepositNonce types.U32
	Call         types.Call
}

// createAssetTxProposal requires a DepositAsset message to construct a assetTxProposal
func createAssetTxProposal(m msg.Message, meta *types.Metadata) (*assetTxProposal, error) {
	recipient := types.NewAccountID(m.Metadata[0:32])
	amount := types.U32(binary.LittleEndian.Uint32(m.Metadata[32:36]))
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
	return &assetTxProposal{
		depositNonce: depositNonce,
		recipient:    recipient,
		amount:       amount,
		call:         call,
	}, nil
}

func (p assetTxProposal) getKey() *proposalKey {
	return &proposalKey{p.depositNonce, p.call}
}

type voteState struct {
	VotesFor     []types.AccountID
	VotesAgainst []types.AccountID
}

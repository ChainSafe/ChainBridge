// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"encoding/binary"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

// proposal represents an on-chain proposal
type proposal struct {
	depositId types.U32
	recipient types.AccountID
	amount    types.U32
	hash      types.Hash
	call      types.Call
}

// createProposalFromAssetTx requires a DepositAsset message to construct a proposal
func createProposalFromAssetTx(m msg.Message, meta *types.Metadata) (*proposal, error) {
	recipient := types.NewAccountID(m.Metadata[0:32])
	amount := types.U32(binary.LittleEndian.Uint32(m.Metadata[32:36]))
	depositId := types.U32(m.DepositId)

	// Create hash
	data := struct {
		DepositId types.U32
		Recipient types.AccountID
		Amount    types.U32
	}{depositId, recipient, amount}
	hash, err := types.GetHash(data)
	if err != nil {
		return nil, err
	}

	call, err := types.NewCall(
		meta,
		NativeTransfer.String(),
		recipient,
		amount,
	)

	if err != nil {
		return nil, err
	}
	return &proposal{
		depositId: depositId,
		recipient: recipient,
		amount:    amount,
		hash:      hash,
		call:      call,
	}, nil
}

type VoteState struct {
	VotesFor     []types.AccountID
	VotesAgainst []types.AccountID
	Hash         types.Hash
}

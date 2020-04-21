// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

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

type EventErc721Minted struct {
	Phase   types.Phase
	Owner   types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Transferred struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Burned struct {
	Phase   types.Phase
	TokenId types.AccountID
	Topics  []types.Hash
}

type EventExampleRemark struct {
	Phase  types.Phase
	Hash   types.Hash
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
	Bridge_VoteFor                 []EventVoteFor                 //nolint:stylecheck,golint
	Bridge_VoteAgainst             []EventVoteAgainst             //nolint:stylecheck,golint
	Bridge_ProposalApproved        []EventProposalApproved        //nolint:stylecheck,golint
	Bridge_ProposalRejected        []EventProposalRejected        //nolint:stylecheck,golint
	Bridge_ProposalSucceeded       []EventProposalSucceeded       //nolint:stylecheck,golint
	Bridge_ProposalFailed          []EventProposalFailed          //nolint:stylecheck,golint
	Erc721_Minted                  []EventErc721Minted            //nolint:stylecheck,golint
	Erc721_Transferred             []EventErc721Transferred       //nolint:stylecheck,golint
	Erc721_Burned                  []EventErc721Burned            //nolint:stylecheck,golint
	Example_Remark                 []EventExampleRemark           //nolint:stylecheck,golint
	Sudo_Sudid                     []EventSudid                   //nolint:stylecheck,golint
}

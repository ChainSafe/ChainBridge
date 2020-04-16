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
}

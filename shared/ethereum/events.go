// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EventSig string

func (es EventSig) GetTopic() common.Hash {
	return crypto.Keccak256Hash([]byte(es))
}

const (
	Deposit           EventSig = "Deposit(uint8,address,uint256)"
	ProposalCreated   EventSig = "ProposalCreated(uint8,uint8,uint256,bytes32)"
	ProposalVote      EventSig = "ProposalVote(uint8,uint8,uint256,uint8)"
	ProposalFinalized EventSig = "ProposalFinalized(uint8,uint8,uint256)"
	ProposalExecuted  EventSig = "ProposalExecuted(uint8,uint8,uint256)"
)

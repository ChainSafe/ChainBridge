// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EventSig string

func (en *EventSig) GetTopic() common.Hash {
	return crypto.Keccak256Hash([]byte(*en))
}

const (
	Deposit                  EventSig = "Deposit(uint8,address,uint256)"
	DepositProposalCreated   EventSig = "DepositProposalCreated(uint8,uint8,uint256,bytes32)"
	DepositProposalVote      EventSig = "DepositProposalVote(uint8,uint8,uint256,uint8)"
	DepositProposalFinalized EventSig = "DepositProposalFinalized(uint8,uint8,uint256)"
	DepositProposalExecuted  EventSig = "DepositProposalExecuted(uint8,uint8,uint256)"
)

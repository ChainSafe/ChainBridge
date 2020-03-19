// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	cfg            Config
	conn           *Connection
	bridgeContract BridgeContract // instance of bound receiver contract
	gasPrice       *big.Int
	gasLimit       *big.Int
}

func NewWriter(conn *Connection, cfg *Config) *Writer {
	return &Writer{
		cfg:      *cfg,
		conn:     conn,
		gasPrice: cfg.gasPrice,
		gasLimit: cfg.gasLimit,
	}
}

func (w *Writer) Start() error {
	log15.Debug("Starting ethereum writer...")
	return nil
}

func (w *Writer) SetBridgeContract(bridge BridgeContract) {
	w.bridgeContract = bridge
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *Writer) ResolveMessage(m msg.Message) bool {
	log15.Trace("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination)

	switch m.Type {
	case msg.DepositAssetType:
		return w.depositAsset(m)
	case msg.CreateDepositProposalType:
		return w.createDepositProposal(m)
	case msg.VoteDepositProposalType:
		return w.voteDepositProposal(m)
	case msg.ExecuteDepositType:
		return w.executeDeposit(m)
	default:
		log15.Warn("Unknown message type received", "type", m.Type)
		return false
	}
}

func (w *Writer) Stop() error {
	return nil
}

func keccakHash(data []byte) [32]byte {
	hash := [32]byte{}
	copy(hash[:], ethcrypto.Keccak256Hash(data).Bytes())
	return hash
}

func u32toBigInt(n uint32) *big.Int {
	return big.NewInt(int64(n))
}

func byteSliceTo32Bytes(in []byte) [32]byte {
	out := [32]byte{}
	// Note: this is safe as copy uses the min length of the two slices
	copy(out[:], in)
	return out
}

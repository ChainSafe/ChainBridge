// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
)

var _ chains.Writer = &writer{}

// https://github.com/ChainSafe/chainbridge-solidity/blob/b5ed13d9798feb7c340e737a726dd415b8815366/contracts/Bridge.sol#L20
var PassedStatus uint8 = 2

type writer struct {
	cfg            Config
	conn           *Connection
	bridgeContract *Bridge.Bridge // instance of bound receiver bridgeContract
	gasPrice       *big.Int
	gasLimit       *big.Int
	log            log15.Logger
}

func NewWriter(conn *Connection, cfg *Config, log log15.Logger) *writer {
	return &writer{
		cfg:      *cfg,
		conn:     conn,
		gasPrice: cfg.gasPrice,
		gasLimit: cfg.gasLimit,
		log:      log,
	}
}

func (w *writer) start() error {
	w.log.Debug("Starting ethereum writer...")
	return nil
}

func (w *writer) setContract(bridge *Bridge.Bridge) {
	w.bridgeContract = bridge
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *writer) ResolveMessage(m msg.Message) bool {
	w.log.Trace("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "rId", m.ResourceId.Hex())

	switch m.Type {
	case msg.FungibleTransfer:
		return w.createErc20Proposal(m)
	case msg.NonFungibleTransfer:
		return w.createErc721Proposal(m)
	case msg.GenericTransfer:
		return w.createGenericDepositProposal(m)
	default:
		w.log.Warn("Unknown message type received", "type", m.Type)
		return false
	}
}

func (w *writer) stop() error {
	return nil
}

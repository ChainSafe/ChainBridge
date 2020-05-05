// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"sync"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var _ chains.Writer = &writer{}

// https://github.com/ChainSafe/chainbridge-solidity/blob/b5ed13d9798feb7c340e737a726dd415b8815366/contracts/Bridge.sol#L20
var PassedStatus uint8 = 2

type writer struct {
	cfg            Config
	conn           *Connection
	bridgeContract *Bridge.Bridge // instance of bound receiver bridgeContract
	callOpts       *bind.CallOpts
	opts           *bind.TransactOpts
	nonce          uint64
	nonceLock      sync.Mutex
	log            log15.Logger
}

// NewWriter creates and returns writer
func NewWriter(conn *Connection, cfg *Config, log log15.Logger) *writer {
	return &writer{
		cfg:  *cfg,
		conn: conn,
		log:  log,
	}
}

// start adds contract call options and transaction options to the writer
func (w *writer) start() error {
	w.log.Debug("Starting ethereum writer...")

	opts, _, err := w.conn.newTransactOpts(big.NewInt(0), w.cfg.gasLimit, w.cfg.gasPrice)
	if err != nil {
		return err
	}

	w.opts = opts
	w.nonce = 0
	w.callOpts = &bind.CallOpts{From: w.conn.kp.CommonAddress()}
	return nil
}

// setContract adds the bound receiver bridgeContract to the writer
func (w *writer) setContract(bridge *Bridge.Bridge) {
	w.bridgeContract = bridge
}

func (w *writer) lockAndUpdateNonce() error {
	w.nonceLock.Lock()
	nonce, err := w.conn.conn.PendingNonceAt(w.conn.ctx, w.opts.From)
	if err != nil {
		w.nonceLock.Unlock()
		return err
	}
	w.opts.Nonce.SetUint64(nonce)
	return nil
}

func (w *writer) unlockNonce() {
	w.nonceLock.Unlock()
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *writer) ResolveMessage(m msg.Message) bool {
	w.log.Trace("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())

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

// stop stops the writer
func (w *writer) stop() error {
	return nil
}

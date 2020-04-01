// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/crypto"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	cfg                  Config
	conn                 *Connection
	bridgeContract       *BridgeContract // instance of bound receiver contract
	erc20HandlerContract *ERC20HandlerContract
	gasPrice             *big.Int
	gasLimit             *big.Int
	log                  log15.Logger
}

func NewWriter(conn *Connection, cfg *Config, log log15.Logger) *Writer {
	return &Writer{
		cfg:      *cfg,
		conn:     conn,
		gasPrice: cfg.gasPrice,
		gasLimit: cfg.gasLimit,
		log:      log,
	}
}

func (w *Writer) Start() error {
	w.log.Debug("Starting ethereum writer...")
	return nil
}

func (w *Writer) SetContracts(bridge *BridgeContract, erc20Handler *ERC20HandlerContract) {
	w.bridgeContract = bridge
	w.erc20HandlerContract = erc20Handler
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *Writer) ResolveMessage(m msg.Message) bool {
	w.log.Trace("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination)

	switch m.Type {
	case msg.FungibleTransfer:
		return w.createErc20DepositProposal(m)
	case msg.NonFungibleTransfer:
		panic("not yet implemented")
	case msg.GenericTransfer:
		panic("not yet implemented")
	default:
		w.log.Warn("Unknown message type received", "type", m.Type)
		return false
	}
}

func (w *Writer) Stop() error {
	return nil
}

func hash(data []byte) [32]byte {
	return crypto.Keccak256Hash(data)
}

func u32toBigInt(n uint32) *big.Int {
	return big.NewInt(int64(n))
}

func byteSliceTo32Bytes(in []byte) [32]byte { //nolint:deadcode,unused
	out := [32]byte{}
	// Note: this is safe as copy uses the min length of the two slices
	copy(out[:], in)
	return out
}

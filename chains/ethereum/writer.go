// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/chains/ethereum/config"
	"github.com/ChainSafe/chainbridge-utils/core"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

var _ core.Writer = &writer{}

// https://github.com/ChainSafe/chainbridge-solidity/blob/b5ed13d9798feb7c340e737a726dd415b8815366/contracts/Bridge.sol#L20
var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4

type writer struct {
	cfg            config.Config
	conn           WriterChainDataFetcher
	bridgeContract *Bridge.Bridge // instance of bound receiver bridgeContract
	log            log15.Logger
	stop           <-chan int
	sysErr         chan<- error // Reports fatal error to core
	metrics        *metrics.ChainMetrics
}

// TODO: Rename once in own package
type WriterChainDataFetcher interface {
	Opts() *bind.TransactOpts
	CallOpts() *bind.CallOpts
	LockAndUpdateOpts() error
	UnlockOpts()
	LatestBlock() (*big.Int, error)
	WaitForBlock(*big.Int, *big.Int) error
	FilterLogs(context.Context, eth.FilterQuery) ([]types.Log, error)
}

// NewWriter creates and returns writer
func NewWriter(cfg *config.Config, conn WriterChainDataFetcher, log log15.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *writer {
	return &writer{
		cfg:     *cfg,
		conn:    conn,
		log:     log,
		stop:    stop,
		sysErr:  sysErr,
		metrics: m,
	}
}

func (w *writer) start() error {
	w.log.Debug("Starting ethereum writer...")
	return nil
}

// setContract adds the bound receiver bridgeContract to the writer
func (w *writer) setContract(bridge *Bridge.Bridge) {
	w.bridgeContract = bridge
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *writer) ResolveMessage(m msg.Message) bool {
	w.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())

	switch m.Type {
	case msg.FungibleTransfer:
		return w.createErc20Proposal(m)
	case msg.NonFungibleTransfer:
		return w.createErc721Proposal(m)
	case msg.GenericTransfer:
		return w.createGenericDepositProposal(m)
	default:
		w.log.Error("Unknown message type received", "type", m.Type)
		return false
	}
}

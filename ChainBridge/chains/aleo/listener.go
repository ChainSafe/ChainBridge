/* Copyright (c) 2021 Forte Labs
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of
 * Forte Labs and their suppliers, if any.  The intellectual and technical
 * concepts contained herein are proprietary to Forte Labs and their suppliers
 * and may be covered by U.S. and Foreign Patents, patents in process, and are
 * protected by trade secret or copyright law. Dissemination of this information
 * or reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Forte Labs.
 */

package aleo

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ChainSafe/chainbridge-utils/msg"

	"github.com/ChainSafe/ChainBridge/chains"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/log15"
)

var CustodianRetryLimit = 5
var CustodianRetryInterval = time.Second * 5
var CustodianARC721Handler = "arc721"

type listener struct {
	cfg                Config
	conn               *Connection
	router             chains.Router
	log                log15.Logger
	latestBlock        metrics.LatestBlock
	stop               <-chan int
	sysErr             chan<- error // Reports fatal error to core
	metrics            *metrics.ChainMetrics
	blockConfirmations *big.Int
}

// NewListener creates and returns a listener
func NewListener(conn *Connection, cfg *Config, log log15.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *listener {
	return &listener{
		cfg:                *cfg,
		conn:               conn,
		log:                log,
		latestBlock:        metrics.LatestBlock{LastUpdated: time.Now()},
		stop:               stop,
		sysErr:             sysErr,
		metrics:            m,
		blockConfirmations: cfg.blockConfirmations,
	}
}

// sets the router
func (l *listener) setRouter(r chains.Router) {
	l.router = r
}

// start registers all subscriptions provided by the config
func (l *listener) start() error {
	l.log.Debug("Starting listener...")

	go func() {
		err := l.pollCustodian()
		if err != nil {
			l.log.Error("Polling custodian failed", "err", err)
		}
	}()

	return nil
}

// pollCustodian will poll the custodian for the latest data and parse the associated events.

func (l *listener) pollCustodian() error {
	var currentBlock = l.cfg.startBlock
	l.log.Info("Polling the Custodian...")

	var retry = CustodianRetryLimit
	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:
			// No more retries
			if retry == 0 {
				l.log.Error("Polling failed, retries exceeded")
				l.sysErr <- errors.New("polling failure")
				return nil
			}

			latestBlock, err := l.conn.LatestBlock()
			if err != nil {
				l.log.Error("rpc error: failed to query the latest block", "err", err)
				retry--
				time.Sleep(CustodianRetryInterval)
				continue
			}

			// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(l.blockConfirmations) == -1 {
				l.log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
				time.Sleep(CustodianRetryInterval)
				continue
			}

			if err := l.getDepositEventsForBlock(latestBlock); err != nil {
				l.log.Error("Failed to get events from custodian", "err", err)
				retry--
				time.Sleep(CustodianRetryInterval)
				continue
			}
			l.latestBlock.Height = big.NewInt(0).Set(latestBlock)
			l.latestBlock.LastUpdated = time.Now()

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = CustodianRetryLimit

			time.Sleep(CustodianRetryInterval)
		}
	}
}

func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying custodian for deposit events")
	logs, err := l.conn.DepositEvents(latestBlock)
	if err != nil {
		return fmt.Errorf("unable to get Deposit Events: %w", err)
	}

	for _, event := range logs {
		var m msg.Message
		var err error
		destId := msg.ChainId(event.DestinationChainID)
		nonce := msg.Nonce(event.Nonce)
		l.log.Info("Deposit Event found", "event", event)
		if event.Handler == CustodianARC721Handler {
			m, err = l.handleArc721DepositedEvent(destId, nonce)
		} else {
			l.log.Error("event has unrecognized handler", "handler", event.Handler)
			return nil
		}

		if err != nil {
			return err
		}

		l.log.Info("generated message to route", "message", m)

		err = l.router.Send(m)
		if err != nil {
			l.log.Error("subscription error: failed to route message", "err", err)
		}
	}

	return nil

}

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
	"time"

	"github.com/ChainSafe/ChainBridge/chains"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/log15"
)

var CustodianRetryLimit = 5

type listener struct {
	cfg         Config
	conn        *Connection
	router      chains.Router
	log         log15.Logger
	latestBlock metrics.LatestBlock
	stop        <-chan int
	sysErr      chan<- error // Reports fatal error to core
	metrics     *metrics.ChainMetrics
}

// NewListener creates and returns a listener
func NewListener(conn *Connection, cfg *Config, log log15.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *listener {
	return &listener{
		cfg:         *cfg,
		conn:        conn,
		log:         log,
		latestBlock: metrics.LatestBlock{LastUpdated: time.Now()},
		stop:        stop,
		sysErr:      sysErr,
		metrics:     m,
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

		}
	}
}

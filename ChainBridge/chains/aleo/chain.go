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

// Package aleo
/*
	The Aleo package contains the logic for interacting with the aleo custodian./
	The supported transfer type is only a Nonfungible token.

	There are 3 major components: the connection, the listener, and the writer.

	Connection

	The Connection handles connecting to the aleo client, and submitting transactions to the custodian.
	This is shared by the writer and listener.

	Listener

	The aleo listener will poll the custodian and parses the events for the transfer. It then forwards the transfer to the router.

	Writer

	The writer will receive messages from the router, builds proposals, then attempts to vote on those proposals.
*/
package aleo

import (
	"github.com/ChainSafe/chainbridge-utils/core"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
)

var _ core.Chain = &Chain{}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     *Connection       // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
}

func InitializeChain(chainCfg *core.ChainConfig, logger log15.Logger, sysErr chan<- error, m *metrics.ChainMetrics) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	stop := make(chan int)
	conn := NewConnection(cfg.endpoint, cfg.http, logger)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	if chainCfg.LatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.startBlock = curr
	}

	listener := NewListener(conn, cfg, logger, stop, sysErr, m)

	writer := NewWriter(conn, cfg, logger, stop, sysErr, m)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, nil
}

func (c *Chain) Start() error {
	err := c.listener.start()
	if err != nil {
		return err
	}

	err = c.writer.start()
	if err != nil {
		return err
	}

	c.conn.log.Debug("Successfully started chain", "chainId", c.cfg.Id)
	return nil
}

func (c *Chain) SetRouter(r *core.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.setRouter(r)
}

func (c *Chain) LatestBlock() metrics.LatestBlock {
	return c.listener.latestBlock
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Name() string {
	return c.cfg.Name
}

func (c *Chain) Stop() {
	close(c.stop)
}

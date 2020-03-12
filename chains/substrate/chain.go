// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/ChainBridgeV2/router"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
)

type Chain struct {
	cfg *core.ChainConfig // The config of the chain
	// TODO: Does this have to be an interface?
	conn     chains.Connection // THe chains connection
	listener *Listener         // The listener of this chain
	writer   *Writer           // The writer of the chain
}

func InitializeChain(cfg *core.ChainConfig) (*Chain, error) {
	kp, err := cfg.Keystore.KeypairFromAddress(cfg.From, keystore.SubChain)
	if err != nil {
		return nil, err
	}
	key := signature.KeyringPair{
		URI:     kp.Public().Address(),
		Address: kp.Public().Address(),
		// TODO: Check that this works
		PublicKey: kp.Public().Encode(),
	}

	conn := NewConnection(cfg.Endpoint, &key)
	l := NewListener(conn, cfg)
	w := NewWriter(conn)
	return &Chain{
		cfg:      cfg,
		conn:     conn,
		listener: l,
		writer:   w,
	}, nil
}

func (c *Chain) Start() error {
	if c.conn == nil {
		return fmt.Errorf("no connection specified")
	}
	if c.listener == nil {
		return fmt.Errorf("no listener specified")
	}
	if c.writer == nil {
		return fmt.Errorf("no Writer specified")
	}

	err := c.listener.Start()
	if err != nil {
		return err
	}

	err = c.writer.Start()
	if err != nil {
		return err
	}

	log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) SetRouter(r *router.Router) {
	r.Listen(c.cfg.Id, c.writer)
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Stop() error {
	err := c.listener.Stop()
	if err != nil {
		return err
	}

	err = c.writer.Stop()
	if err != nil {
		return err
	}

	c.conn.Close()

	return nil
}

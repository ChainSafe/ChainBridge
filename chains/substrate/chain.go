// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto/sr25519"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/ChainBridgeV2/router"
	log "github.com/ChainSafe/log15"
)

type Chain struct {
	cfg *core.ChainConfig // The config of the chain
	// TODO: Does this have to be an interface?
	conn     *Connection // THe chains connection
	listener *Listener   // The listener of this chain
	writer   *Writer     // The writer of the chain
}

func InitializeChain(cfg *core.ChainConfig) (*Chain, error) {
	kp, err := keystore.KeypairFromAddress(cfg.From, keystore.SubChain, cfg.KeystorePath, false)
	if err != nil {
		return nil, err
	}
	krp := kp.(*sr25519.Keypair).AsKeyringPair()
	conn := NewConnection(cfg.Endpoint, krp)
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
	err := c.listener.Start()
	if err != nil {
		return err
	}

	err = c.writer.Start()
	if err != nil {
		return err
	}

	log.Debug("Successfully started chain", "id", c.cfg.Id.String())
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

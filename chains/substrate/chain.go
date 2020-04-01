// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"errors"

	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/crypto/sr25519"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/ChainBridge/router"
	"github.com/ChainSafe/log15"
)

type Chain struct {
	cfg *core.ChainConfig // The config of the chain
	// TODO: Does this have to be an interface?
	conn     *Connection // THe chains connection
	listener *Listener   // The listener of this chain
	writer   *Writer     // The writer of the chain
}

func InitializeChain(cfg *core.ChainConfig) (*Chain, error) {
	kp, err := keystore.KeypairFromAddress(cfg.From, keystore.SubChain, cfg.KeystorePath, cfg.Insecure)
	if err != nil {
		return nil, err
	}

	logger := log15.Root().New("chain", cfg.Name)
	krp := kp.(*sr25519.Keypair).AsKeyringPair()

	// Setup connection
	conn := NewConnection(cfg.Endpoint, cfg.Name, krp, logger)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}
	var id msg.ChainId
	ok, err := conn.queryStorage("Bridge", "ChainIdentifier", krp.PublicKey, nil, &id)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("Unable to find ChainId")
	} else if id != cfg.Id {
		return nil, errors.New("ChainID is incorrect")
	}

	// Setup listener & writer
	startBlock := parseStartBlock(cfg)
	l := NewListener(conn, cfg.Name, cfg.Id, startBlock, logger)
	w := NewWriter(conn, logger)
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

	c.conn.log.Debug("Successfully started chain", "chainId", c.cfg.Id)
	return nil
}

func (c *Chain) SetRouter(r *router.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.SetRouter(r)
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Name() string {
	return c.cfg.Name
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

// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"github.com/ChainSafe/ChainBridge/blockstore"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/crypto/sr25519"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/ChainBridge/router"
	"github.com/ChainSafe/log15"
)

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     *Connection       // THe chains connection
	listener *Listener         // The listener of this chain
	writer   *Writer           // The writer of the chain
}

// checkBlockstore queries the blockstore for the latest known block. If the latest block is
// greater than startBlock, then the latest block is returned, otherwise startBlock is.
func checkBlockstore(bs *blockstore.Blockstore, startBlock uint64) (uint64, error) {
	latestBlock, err := bs.TryLoadLatestBlock()
	if err != nil {
		return 0, err
	}

	if latestBlock.Uint64() > startBlock {
		return latestBlock.Uint64(), nil
	} else {
		return startBlock, nil
	}
}

func InitializeChain(cfg *core.ChainConfig, logger log15.Logger) (*Chain, error) {
	kp, err := keystore.KeypairFromAddress(cfg.From, keystore.SubChain, cfg.KeystorePath, cfg.Insecure)
	if err != nil {
		return nil, err
	}

	krp := kp.(*sr25519.Keypair).AsKeyringPair()

	// Attempt to load latest block
	bs, err := blockstore.NewBlockstore(cfg.BlockstorePath, cfg.Id, kp.Address())
	if err != nil {
		return nil, err
	}
	startBlock := parseStartBlock(cfg)
	if !cfg.FreshStart {
		startBlock, err = checkBlockstore(bs, startBlock)
		if err != nil {
			return nil, err
		}
	}

	// Setup connection
	conn := NewConnection(cfg.Endpoint, cfg.Name, krp, logger)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	err = conn.checkChainId(cfg.Id)
	if err != nil {
		return nil, err
	}

	// Setup listener & writer
	l := NewListener(conn, cfg.Name, cfg.Id, startBlock, logger, bs)
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

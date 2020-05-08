// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

/*
The substrate package contains the logic for interacting with substrate chains.
The current supported events are Fungible, Nonfungible, and generic transfer events.

There are 4 major components: The chain, the connection, the listener, and the writer.

The Chain

The Chain handles initializing the Chain, which includes creating the connection, writer, and listener.
There is also functionality for starting and stopping the Chain.

The Connection

The Connection handles connecting to the substrate client, and submitting transactions to the client.
In addition, the connection also can query and obtain both data and metadata from the client.
The connection is shared by the writer and listener.

The Listener

The listener listens to event coming from the client and posts them to the router.
To accomplish this, the listener requests the data within the blocks on the chain (known as polling the blocks) and reads the data to determine if any events have happened.
If any events do happen, the listener sends a message to the router, which will be passed onto the destination chains writer.

The Writer

The writer listens to event coming from the router and writes them to the chain.
The writer upon hearing from the router posts a proposal message to the client.
If the proposal is found to be active, then the writer executes the event on the chain.

*/
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
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
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

func (c *Chain) SetRouter(r *router.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.setRouter(r)
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

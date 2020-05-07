// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
The ethereum package contains the logic for interacting with ethereum chains.

There are 4 major components: The chain, the connection, the listener, and the writer.
The currently supported actions are The currently supported events are Erc20, Erc721, and generic deposit events.

The Chain

The Chain handles the overheard and initalization of the Chain. Calling the Chain to create and connect to the proper ports and set the listener and writer.
In addition to setting up the Chain, it also has the functionality to start the Chain and send transactions to the ethereum chain.

The Connection

The connection connects to the ethereum client and can be accessed by both the writer and listener.

The Listener

The listener's job is to handle events happening on the local chain and send the events to the bridge.
The listen vs write verbs are relative to the chain that the service is attached too.
The ethereum listener "listens" to the chain and writes to the router.
The ethereum writer "writes" to the chain and listens to the router.


The Writer

The writer's job is to obtain events from the router, and create the corresponding event on the ethereum chain and then post that event to the chain.
An example flow of the writer would be:
 - A listener on a source chain hears an event and parses the event into a message and sends it the router.
 - The router passes the message to the ethereum writer.
 - The writer recieves the message and creates a proposal on-chain.
 - The writer then watches for a finalization event.
 - After seeing the finalization event, the writer executes the proposal.
*/
package ethereum

import (
	"context"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc721Handler "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
	"github.com/ChainSafe/ChainBridge/blockstore"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
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
// greater than cfg.startBlock, then cfg.startBlock is replaced with the latest known block.
func setupBlockstore(cfg *Config, kp *secp256k1.Keypair) (*blockstore.Blockstore, error) {
	bs, err := blockstore.NewBlockstore(cfg.blockstorePath, cfg.id, kp.Address())
	if err != nil {
		return nil, err
	}

	if !cfg.freshStart {
		latestBlock, err := bs.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}

		if latestBlock.Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = latestBlock
		}
	}

	return bs, nil
}

func InitializeChain(chainCfg *core.ChainConfig, logger log15.Logger, ctx context.Context) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	kpI, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath, chainCfg.Insecure)
	if err != nil {
		return nil, err
	}
	kp, _ := kpI.(*secp256k1.Keypair)

	bs, err := setupBlockstore(cfg, kp)
	if err != nil {
		return nil, err
	}

	conn := NewConnection(cfg, kp, logger, ctx)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	err = conn.ensureHasBytecode(cfg.bridgeContract)
	if err != nil {
		return nil, err
	}
	err = conn.ensureHasBytecode(cfg.erc20HandlerContract)
	if err != nil {
		return nil, err
	}
	err = conn.ensureHasBytecode(cfg.genericHandlerContract)
	if err != nil {
		return nil, err
	}

	bridgeContract, err := bridge.NewBridge(cfg.bridgeContract, conn.conn)
	if err != nil {
		return nil, err
	}

	erc20HandlerContract, err := erc20Handler.NewERC20Handler(cfg.erc20HandlerContract, conn.conn)
	if err != nil {
		return nil, err
	}

	erc721HandlerContract, err := erc721Handler.NewERC721Handler(cfg.erc721HandlerContract, conn.conn)
	if err != nil {
		return nil, err
	}

	genericHandlerContract, err := GenericHandler.NewGenericHandler(cfg.genericHandlerContract, conn.conn)
	if err != nil {
		return nil, err
	}

	listener := NewListener(conn, cfg, logger, bs)
	listener.setContracts(bridgeContract, erc20HandlerContract, erc721HandlerContract, genericHandlerContract)

	writer := NewWriter(conn, cfg, logger)
	writer.setContract(bridgeContract)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
	}, nil
}

func (c *Chain) SetRouter(r *router.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.setRouter(r)
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

	c.writer.log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Name() string {
	return c.cfg.Name
}

func (c *Chain) Stop() error {
	err := c.listener.stop()
	if err != nil {
		return err
	}

	err = c.writer.stop()
	if err != nil {
		return err
	}

	c.conn.Close()

	return nil
}

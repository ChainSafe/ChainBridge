// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
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

var _ core.Chain = &Chain{}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     *Connection       // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
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

func InitializeChain(chainCfg *core.ChainConfig, logger log15.Logger, sysErr chan<- error) (*Chain, error) {
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

	stop := make(chan int)
	conn := NewConnection(cfg, kp, logger, stop)
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

	listener := NewListener(conn, cfg, logger, bs, stop, sysErr)
	listener.setContracts(bridgeContract, erc20HandlerContract, erc721HandlerContract, genericHandlerContract)

	writer := NewWriter(conn, cfg, logger, stop, sysErr)
	writer.setContract(bridgeContract)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
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

// Stop signals to any running routines to exit
func (c *Chain) Stop() {
	close(c.stop)
}

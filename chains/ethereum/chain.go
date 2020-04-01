// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	"github.com/ChainSafe/ChainBridge/chains"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/ChainBridge/router"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
)

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     *Connection       // THe chains connection
	listener *Listener         // The listener of this chain
	writer   *Writer           // The writer of the chain
}

func createBridgeContract(addr common.Address, conn *Connection) (*BridgeContract, error) {
	bridgeInstance, err := bridge.NewBridge(addr, conn.conn)
	if err != nil {
		return nil, err
	}

	raw := &bridge.BridgeRaw{
		Contract: bridgeInstance,
	}

	return &BridgeContract{
		BridgeRaw:        raw,
		BridgeCaller:     &bridgeInstance.BridgeCaller,
		BridgeTransactor: &bridgeInstance.BridgeTransactor,
	}, nil
}

func createErc20HandlerContract(addr common.Address, conn *Connection) (*ERC20HandlerContract, error) {
	instance, err := erc20Handler.NewERC20Handler(addr, conn.conn)
	if err != nil {
		return nil, err
	}

	raw := &erc20Handler.ERC20HandlerRaw{
		Contract: instance,
	}

	return &ERC20HandlerContract{
		ERC20HandlerRaw:    raw,
		ERC20HandlerCaller: &instance.ERC20HandlerCaller,
	}, nil
}

func InitializeChain(chainCfg *core.ChainConfig) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	kpI, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath, chainCfg.Insecure)
	if err != nil {
		return nil, err
	}

	kp, _ := kpI.(*secp256k1.Keypair)

	logger := log15.Root().New("chain", cfg.name)

	conn := NewConnection(cfg, kp, logger)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	err = conn.checkBridgeContract(cfg.contract)
	if err != nil {
		return nil, err
	}

	bridgeContract, err := createBridgeContract(cfg.contract, conn)
	if err != nil {
		return nil, err
	}

	erc20HandlerContract, err := createErc20HandlerContract(cfg.erc20HandlerContract, conn)
	if err != nil {
		return nil, err
	}

	listener := NewListener(conn, cfg, logger)
	listener.SetContracts(bridgeContract, erc20HandlerContract)

	writer := NewWriter(conn, cfg, logger)
	writer.SetContracts(bridgeContract, erc20HandlerContract)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
	}, nil
}

func (c *Chain) SetRouter(r *router.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.SetRouter(r)
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

	c.writer.log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Name() string {
	return c.cfg.Name
}

func (c *Chain) GetWriter() chains.Writer {
	return c.writer
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

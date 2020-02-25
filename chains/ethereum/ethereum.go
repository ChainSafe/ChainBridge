// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	log "github.com/ChainSafe/log15"
)

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     chains.Connection // THe chains connection
	listener *Listener         // The listener of this chain
	writer   *Writer           // The writer of the chain
}

func InitializeChain(chainCfg *core.ChainConfig) *Chain {
	cfg := ParseChainConfig(chainCfg)

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		panic(err)
	}

	emitterContract, err := emitter.NewEmitter(cfg.emitter, conn.conn)
	if err != nil {
		panic(err)
	}

	listener := NewListener(conn, cfg)
	listener.SetEmitterContract(&emitterContract.EmitterFilterer)

	receiverContract, err := receiver.NewReceiver(cfg.receiver, conn.conn)
	if err != nil {
		panic(err)
	}

	instance := &receiver.ReceiverRaw{
		Contract: receiverContract,
	}

	writer := NewWriter(conn, cfg)
	writer.SetReceiverContract(instance)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
	}
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

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Connection() chains.Connection {
	return c.conn
}

func (c *Chain) GetWriter() chains.Writer {
	return c.writer
}

func (c *Chain) GetListner() chains.Listener {
	return c.listener
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

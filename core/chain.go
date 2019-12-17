package core

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

type Chain struct {
	id       msg.ChainId // Unique chain identifier (see package message
	home     string      // home bridge address
	away     string      // away bridge address
	conn     chains.Connection
	listener chains.Listener
	writer   chains.Writer
}

type ChainConfig struct {
	Endpoint string `toml:"endpoint"` // url for rpc endpoint
	Home     string `toml:"home"`     // home bridge address
	Away     string `toml:"away"`     // away bridge address
	From     string `toml:"from"`     // address of key to use
}

func NewChain(id msg.ChainId, cfg *ChainConfig) *Chain {
	return &Chain{
		id:   id,
		home: cfg.Home,
		away: cfg.Away,
	}
}

func (c *Chain) SetConnection(conn chains.Connection) {
	c.conn = conn
}

func (c *Chain) SetListener(listener chains.Listener) {
	c.listener = listener
}

func (c *Chain) SetWriter(Writer chains.Writer) {
	c.writer = Writer
}

func (c *Chain) GetWriter() chains.Writer {
	return c.writer
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

	return c.conn.Connect()
}

func (c *Chain) Id() msg.ChainId {
	return c.id
}

func (c *Chain) Connection() chains.Connection {
	return c.conn
}

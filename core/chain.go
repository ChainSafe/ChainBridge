package core

import (
	"fmt"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

type Chain struct {
	cfg      *ChainConfig // away bridge address
	conn     Connection
	listener Listener
	writer   Writer
	//subscriptions []string
}

type ChainConfig struct {
	Id            msg.ChainId `toml:"id"`       // ChainID
	Endpoint      string      `toml:"endpoint"` // url for rpc endpoint
	Home          string      `toml:"home"`     // home bridge address
	Away          string      `toml:"away"`     // away bridge address
	From          string      `toml:"from"`     // address of key to use
	Subscriptions []string    `toml:"subscriptions"`
}

func NewChain(cfg *ChainConfig) *Chain {
	return &Chain{
		cfg: cfg,
	}
}

func (c *Chain) SetConnection(conn Connection) {
	c.conn = conn
}

func (c *Chain) SetListener(listener Listener) {
	c.listener = listener
}

func (c *Chain) SetWriter(Writer Writer) {
	c.writer = Writer
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

	err := c.conn.Connect()
	if err != nil {
		return err
	}

	err = c.listener.Start()
	if err != nil {
		return err
	}

	err = c.writer.Start()
	if err != nil {
		return err
	}
	return nil
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Connection() Connection {
	return c.conn
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

package core

import (
	msg "ChainBridgeV2/message"
	"fmt"
)

type Chain struct {
	id       msg.ChainId // Unique chain identifier (see package message
	endpoint string      // url for rpc endpoint
	home     []byte      // home bridge address
	away     []byte      // away bridge address
	conn     Connection
	listener Listener
	pusher   Pusher
}

func NewChain(id msg.ChainId, endpoint string, home, away []byte) *Chain {
	return &Chain{
		id:       id,
		endpoint: endpoint,
		home:     home,
		away:     away,
	}
}

func (c *Chain) SetConnection(conn Connection) {
	c.conn = conn
}

func (c *Chain) SetListener(listener Listener) {
	c.listener = listener
}

func (c *Chain) SetPusher(pusher Pusher) {
	c.pusher = pusher
}

func (c *Chain) Start() error {
	if c.conn == nil {
		return fmt.Errorf("no connection specified")
	}
	if c.listener == nil {
		return fmt.Errorf("no listener specified")
	}
	if c.pusher == nil {
		return fmt.Errorf("no pusher specified")
	}

	return c.conn.Connect()
}

func (c *Chain) Id() msg.ChainId {
	return c.id
}

func (c *Chain) Connection() Connection {
	return c.conn
}

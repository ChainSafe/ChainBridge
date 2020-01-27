package ethereum

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
)

func InitializeChain(chainCfg *core.ChainConfig) *core.Chain {
	cfg := ParseChainConfig(chainCfg)

	c := core.NewChain(chainCfg)
	conn := NewConnection(cfg)
	c.SetConnection(conn)

	emitterContract, err := emitter.NewEmitter(cfg.emitter, conn.conn)
	if err != nil {
		panic(err)
	}

	listener := NewListener(conn, cfg)
	listener.SetEmitterContract(&emitterContract.EmitterFilterer)
	c.SetListener(listener)

	receiverContract, err := receiver.NewReceiver(cfg.receiver, conn.conn)
	if err != nil {
		panic(err)
	}

	instance := &receiver.ReceiverRaw{
		Contract: receiverContract,
	}

	writer := NewWriter(conn, cfg)
	writer.SetReceiverContract(instance)
	c.SetWriter(writer)

	return c
}

package ethereum

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ethereum/go-ethereum/common"
)

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	id            msg.ChainId    // ChainID
	endpoint      string         // url for rpc endpoint
	receiver      common.Address // bridge address to call
	emitter       common.Address // bridge address where events occur
	from          string         // address of key to use
	subscriptions []string       // list of events to subscribe to
	keystore      *keystore.Keystore
}

// ParseChainConfig uses a core.ChainConfig to construct a corresponding Config
func ParseChainConfig(chainCfg *core.ChainConfig) *Config {
	return &Config{
		id:            chainCfg.Id,
		endpoint:      chainCfg.Endpoint,
		receiver:      common.HexToAddress(chainCfg.Receiver),
		emitter:       common.HexToAddress(chainCfg.Emitter),
		from:          chainCfg.From,
		subscriptions: chainCfg.Subscriptions,
		keystore:      chainCfg.Keystore,
	}
}

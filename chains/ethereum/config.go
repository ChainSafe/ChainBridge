package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ethereum/go-ethereum/common"
)

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	id            msg.ChainId        // ChainID
	chainID       *big.Int           // Ethereum chain ID
	endpoint      string             // url for rpc endpoint
	receiver      common.Address     // bridge address to call
	emitter       common.Address     // bridge address where events occur
	from          string             // address of key to use
	subscriptions []string           // list of events to subscribe to
	keystore      *keystore.Keystore // Location of keyfiles
}

type configOpts struct {
	chainID *big.Int // Ethereum chain ID
}

// ParseChainConfig uses a core.ChainConfig to construct a corresponding Config
func ParseChainConfig(chainCfg *core.ChainConfig) *Config {

	opts := parseOpts(chainCfg.Opts)

	return &Config{
		id:            chainCfg.Id,
		chainID:       opts.chainID,
		endpoint:      chainCfg.Endpoint,
		receiver:      common.HexToAddress(chainCfg.Receiver),
		emitter:       common.HexToAddress(chainCfg.Emitter),
		from:          chainCfg.From,
		subscriptions: chainCfg.Subscriptions,
		keystore:      chainCfg.Keystore,
	}
}

func parseOpts(opts map[string]string) configOpts {

	config := defaultConfig()

	if chainID, ok := opts["ChainID"]; ok {
		config.chainID = big.NewInt(0)
		config.chainID.SetString(chainID, 10)
	}

	return config
}

func defaultConfig() configOpts {
	return configOpts{
		chainID: big.NewInt(1),
	}
}

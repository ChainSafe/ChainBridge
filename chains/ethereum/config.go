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
	var chainID *big.Int

	if opts["ChainID"] != "" {
		chainID := new(big.Int)
		chainID.SetString(opts["ChainID"], 10)
	}

	return configOpts{
		chainID,
	}
}

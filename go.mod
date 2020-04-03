module github.com/ChainSafe/ChainBridge

go 1.13

require (
	github.com/ChainSafe/log15 v1.0.0
	github.com/btcsuite/btcutil v1.0.1
	github.com/centrifuge/go-substrate-rpc-client v1.1.1-0.20200310150725-e088497b941f
	github.com/ethereum/go-ethereum v1.9.10
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/pierrec/xxHash v0.1.5 // indirect
	github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4
	github.com/urfave/cli v1.22.2
	golang.org/x/crypto v0.0.0-20200128174031-69ecbb4d6d5d
	gotest.tools v2.2.0+incompatible
)

replace github.com/centrifuge/go-substrate-rpc-client => github.com/ansermino/go-substrate-rpc-client v1.1.1-0.20200326234327-118ba514039a

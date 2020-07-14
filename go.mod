module github.com/ChainSafe/ChainBridge

go 1.13

require (
	github.com/ChainSafe/chainbridge-substrate-events v0.0.0-20200527185312-f0db52b1c793
	github.com/ChainSafe/log15 v1.0.0
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/centrifuge/go-substrate-rpc-client v2.0.0-alpha.3+incompatible
	github.com/ethereum/go-ethereum v1.9.13
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/stretchr/testify v1.4.0
	github.com/urfave/cli/v2 v2.2.0
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79
)

replace github.com/ChainSafe/chainbridge-substrate-events v0.0.0-20200527185312-f0db52b1c793 => github.com/mikiquantum/chainbridge-substrate-events v0.0.0-20200714203812-fb6d256611a4

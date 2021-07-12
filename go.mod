module github.com/ChainSafe/ChainBridge

go 1.15

require (
	github.com/ChainSafe/chainbridge-substrate-events v0.0.0-20200715141113-87198532025e
	github.com/ChainSafe/chainbridge-utils v1.0.6
	github.com/ChainSafe/log15 v1.0.0
	github.com/centrifuge/go-substrate-rpc-client/v3 v3.0.2-0.20210617193115-00abec4508f4
	github.com/ethereum/go-ethereum v1.10.4
	github.com/prometheus/client_golang v1.4.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
)

replace github.com/ChainSafe/chainbridge-substrate-events => github.com/ParthDesai/chainbridge-substrate-events v0.0.0-20210617050925-1781bf7f0e70

replace github.com/ChainSafe/chainbridge-utils => github.com/ParthDesai/chainbridge-utils v1.0.7-0.20210617044003-5cd33cb31ddb

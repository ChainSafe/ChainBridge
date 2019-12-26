# ChainBridge V2

#**[WIP]**

## Run the bridge
1. `make run`
2. `make build`

## Configuring the bridge

### Selecting Chains

Presently chains must be manually inserted into `cmd/chainbridge/main.go`. In future this will be more flexible and require less manual configuration.

For the time being we have `ethereum` and `centrifuge` which are both standard Ethereum chains.

### Configuring Chains

A chain configurations take this form:
```toml
[[chains]]
id = <see ./message/id.go>
endpoint = <RPC (WS) endpoint>
receiver = <bridge contract address>
emitter = <bridge contract address>
from = <public key to use for txs>
```

See `config.toml` for an example configuration. 

Note: presently a home and away contracts can be specified, these can be the same contract.


### Keystore

To manage key ChainBridge uses a keystore specificed with the `--keystore <path>` flag. By default it uses `./keys`. The addresses specified in the config will be used to identify which keys to load.

Keys can be managed with the `account` sub-command. Please see `chainbridge account --help` for documentation.

## Ethereum Dev Environment 

To start a ganache instance:
```
make start_eth
```

Bridge contracts can then be deployed with:
```
make deploy_eth
```

Note: The environment variable `PORT=<port>` can be provided for these commands (default `PORT=8545`)

## Tests

### Go tests
To run the go tests a fresh ganache instance is required (tests depend on deterministic addresses). 
A new instance can be started by running these in seperate terminals:
```
make start_eth
make deploy_eth
```
Go tests can then be run with:
```
make test
```
### Contract Tests
Truffle tests can be run with just:
```
make truffle_test
```
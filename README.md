# ChainBridge V2

#**[WIP]**

## Run the bridge
1. `make run`

or

1. `make build`
2. `./build/chainbridge`

## Configuring the bridge

### Selecting Chains

Presently chains must be manually inserted into `cmd/chainbridge/main.go`. In future this will be more flexible and require less manual configuration.

For the time being we have `ethereum` and `ethereum2` which are both standard Ethereum chains.

### Configuring Chains

A chain configurations take this form:
```toml
[<name>]
endpoint = <RPC (WS) endpoint>
receiver = <bridge contract address>
emitter = <bridge contract address>
from = <address of key to use for txs>
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

Go tests can run with:
```
make test
```

Truffle tests can be run with:
```
make truffle_test
```
# ChainBridge V2

# **[WIP]**

# Chain configs
For chain specific configs, please check [this subdirectory](./chain-documents)

## Run the bridge
 
`make run`

See `--help` for CLI options.

## Configuring the bridge

### Selecting Chains

Presently chains must be manually inserted into `cmd/chainbridge/main.go`. In future this will be more flexible and require less manual configuration.

For the time being we have `ethereum` and `centrifuge` chains configured.

### Configuring Chains

A chain configurations take this form:
```toml
[[chains]]
id = 0 # see ./message/id.go
endpoint = "ws://localhost:8545" # RPC (WS) endpoint
receiver = "0x1234" # bridge receiver contract address
emitter = "0x1234" # bridge emitter contract address>
from = "0x1234" # public key to use for txs
```

See `config.toml` for an example configuration. 

Note: presently a home and away contracts can be specified, these can be the same contract.


### Keystore

To manage keys ChainBridge uses a keystore specificed with the `--keystore <path>` flag. By default it uses `./keys`. The public key specified in the config will be used to identify which keys to load.

Keys can be managed with the `account` sub-command. Please see `chainbridge account --help` for documentation.

Alternatively, an environemnet variable can be used with the key `KEYSTORE_PASSWORD`.

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

To build the go bindings for the Ethereum contracts:
```
make bindings
```

## Centrifuge Dev Environment

To fetch, build and run centrifuge-chain run:
```
make start_cent
```

Note: The build process takes a while, but will only run once.

You can submit a random asset hash for validation (ie. transfer) with:
```
make cent_asset_tx
```
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

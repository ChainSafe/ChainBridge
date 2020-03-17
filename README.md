# ChainBridge V2

[![Build Status](https://travis-ci.com/ChainSafe/ChainBridgeV2.svg?branch=master)](https://travis-ci.com/ChainSafe/ChainBridgeV2)

# **[WIP]**

# Chain configs
For chain specific configs, please check [this subdirectory](./chain-documents)

## Dependencies

- [Subkey](https://github.com/paritytech/substrate): 
Required for substrate key management.

  `make install-subkey`

- [abigen](https://geth.ethereum.org/docs/dapp/native-bindings):
Builds go bindings for Solidity contract interactions.

  See [Installation Instructions](https://geth.ethereum.org/docs/install-and-build/installing-geth).

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

Note: The build process takes a while, but will only run once. It currently uses a modified fork of centrifuge-chain

You can run several commands to interact with the bridge module:

### Emitter Address

You can set and get the emitter address with:
 ```
 make cent_get_emitter
``` 
and 
```
make cent_set_emitter CENT_EMITTER_ADDR=<HEX VALUE>
```
### Whitelist Chain

A chain ID can be whitelisted as a destination with:

```
make cent_whitelist_chain CENT_CHAIN_ID=<HEX VALUE>
```

### Asset Transfer

An asset transfer can be executed with:

```
make cent_asset_tx CENT_CHAIN_ID=<HEX VALUE> CENT_TO=<HEX VALUE> CENT_TOKEN_ID=<HEX VALUE> CENT_METADATA=<HEX VALUE>
```

### Auto Run

Setting an emitter address, whitelisting a chain and submitting an asset tx can be executed using default values with:
```
make cent_auto_run 
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

## Simulations
If you have a bridge instance running, and access to an ethereum node, transactions can be run that simulate different types of transfer events. These scripts can be found at `./on-chain/evm-contracts/scripts/cli`. 

### Running the simulations
1. `./on-chain/evm-contracts`
2. Run one of the following:
###### Deploy contracts
`node ./scripts/cli/index.js -p <port_number>`
###### Mint tokens
`node ./scripts/cli/index.js --test-only --mint-erc20 --value <amount of token>`
###### ERC transfer (note must mint tokens first) 
`node ./scripts/cli/index.js --test-only --deposit-erc --value <amount_to_deposit> --dest <destination_chain_id>`
###### NFT transfer
// Outdated

Notes: 
- `--test-only` ensures we don't re-deploy the contracts (this must be refactored out in favor of commands)
- `--dest` allows you to specify which chain_id you want to the transfer to go to 
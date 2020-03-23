# ChainBridge V2

[![Build Status](https://travis-ci.com/ChainSafe/ChainBridgeV2.svg?branch=master)](https://travis-ci.com/ChainSafe/ChainBridgeV2)

<h3><b>[WIP]</b></h3>

# Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Running](#running)
- [Testing](#testing)
- [Simulations](#simulations)

# Installation

## Dependencies

- [Subkey](https://github.com/paritytech/substrate): 
Required for substrate key management.

  `make install-subkey`

- [abigen](https://geth.ethereum.org/docs/dapp/native-bindings):
Builds go bindings for Solidity contract interactions.

  See [Installation Instructions](https://geth.ethereum.org/docs/install-and-build/installing-geth).

## Building

`make build`: Builds `chainbridge` in `./build`.

`make install`: Uses `go install` to add `chainbridge` to your GOBIN.

## Configuration

### Configuring Chains

A chain configurations take this form:
```toml
[[chains]]
name = "ethereum" // Human-readable name
type = "ethereum" // Either "ethereum" or "substrate"
id = 0            // Chain Id
endpoint = "ws://host:port" // API endpoint
from = "029b67ec8aba36421137e22d874a897f8aa2a47e2d479d772d96ca8c5744b5a95c" // Public key of desired key, not required for test keys
opts = {}         // Chain-specific configuration options (see below)
```

See `config.toml.example` for an example configuration. 

#### Ethereum Options

Ethereum chains support the following additional options:

```
contract = ""0x12345..." // The address of the bridge contract
gasPrice = "0x1234"      // Gas price for transactions 
gasLimit = "0x1234"      // Gas limit for transactions
http = "true"            // Whether the chain connection is ws or http
```

#### Substrate Options

There are presently no additional config options for substrate chains.

### Keystore

ChainBridge requires keys to sign and submit transactions, and to identify each bridge node on chain.

To use secure keys, see `chainbridge accounts --help`. The keystore password can be supplied with the `KEYSTORE_PASSWORD` environment variable.

For testing purposes, chainbridge provides 5 test keys. The can be used with `--testkey <name>`, where `name` is one of `Alice`, `Bob`, `Charlie`, `Dave`, or `Eve`. 

# Testing

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
### Ethereum ERC20 Transfer
Start chain 1 (terminal 1)
```shell
make start_eth
```
Start chain 2 (terminal 2)
```shell
PORT=8546 make start_eth
```
Deploy the contracts (terminal 3)
```shell
make deploy_eth && PORT=8546 make deploy_eth
```
Build the latest ChainBridge binary & run it (terminal 3)
```shell
make build
./build/chainbridge --verbosity=trace --config ./scripts/configs/config1.toml --testkey alice
```
Mint & make a deposit (terminal 4)
```shell
node on-chain/evm-contracts/scripts/cli/index.js --test-only --mint-erc20 --value 100
node on-chain/evm-contracts/scripts/cli/index.js --test-only --deposit-erc —dest 1
```

Notes: 
- Alice (from the keyring) is always the deployer, if that key changes, then the constants will be different
- Validators start from the keyring and move alphabetically down the list. For example if you specify `--validators 3`, the validators would be `Alice`, `Bob`, `Charlie`. If you said 4, `Dave` would join
- `--test-only` ensures we don't re-deploy the contracts
- `--dest` allows you to specify which chain_id you want to the transfer to go to

#### Debugging
Node script errors:
"Contract not found" or similar:
- Check the deployments in step 3, do the addresses listed there match with the addresses saved in `.on-chain/evm-contracts/scripts/cli/constants.js`? The constants file should be updated accordingly
"Sender doesn't have funds" or similar when executing an erc20 transfer:
- Check that the you ran `--mint <value>` (step 4) if you didn't the account has no tokens to deposit

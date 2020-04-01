# ChainBridge

[![Build Status](https://travis-ci.com/ChainSafe/ChainBridge.svg?branch=master)](https://travis-ci.com/ChainSafe/ChainBridge)

<h3><b>[WIP]</b></h3>

# Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Running](#running)
- [Chain Implementations](#chain-implementations)
- [Testing](#testing)
- [Simulations](#simulations)

# Installation

## Dependencies

- [Subkey](https://github.com/paritytech/substrate): 
Required for substrate key management.

  `make install-subkey`


## Building

`make build`: Builds `chainbridge` in `./build`.

**or**

`make install`: Uses `go install` to add `chainbridge` to your GOBIN.

## Configuration

### Configuring Chains

A chain configurations take this form:
```toml
[[chains]]
name = "ethereum" # Human-readable name
type = "ethereum" # Either "ethereum" or "substrate"
id = 0            # Chain Id
endpoint = "ws://host:port" # API endpoint
from = "029b67ec8aba36421137e22d874a897f8aa2a47e2d479d772d96ca8c5744b5a95c" # Public key of desired key, not required for test keys
opts = {}         # Chain-specific configuration options (see below)
```

See `config.toml.example` for an example configuration. 

#### Ethereum Options

Ethereum chains support the following additional options:

```
contract = "0x12345..." // The address of the bridge contract (required)
gasPrice = "0x1234"      // Gas price for transactions (default: 20000000000)
gasLimit = "0x1234"      // Gas limit for transactions (default: 6721975)
http = "true"            // Whether the chain connection is ws or http (default: false)
```

#### Substrate Options

Substrate supports the following additonal options:

```
startBlock = "1234" // The block to start processing event from (default: 0)
```

### Keystore

ChainBridge requires keys to sign and submit transactions, and to identify each bridge node on chain.

To use secure keys, see `chainbridge accounts --help`. The keystore password can be supplied with the `KEYSTORE_PASSWORD` environment variable.

To import external ethereum keys, such as those generated with geth, use `chainbridge accounts import --ethereum /path/to/key`.

For testing purposes, chainbridge provides 5 test keys. The can be used with `--testkey <name>`, where `name` is one of `Alice`, `Bob`, `Charlie`, `Dave`, or `Eve`. 

# Chain Implementations

- Ethereum (Solidity): [chainbridge-solidity](https://github.com/ChainSafe/chainbridge-solidity) 

    The Solidity contracts required for chainbridge. Includes deployment and interaction CLI.
    
    The bindings for the contracts live in `bindings/`. To update the bindings modify `scripts/setup-contracts.sh` and then run `make clean && make setup-contracts`

- Substrate: [chainbridge-substrate](https://github.com/ChainSafe/chainbridge-substrate)

    A substrate pallet that can be integrated into a chain, as well as an example pallet to demonstrate chain integration.

# Testing

## Ethereum Dev Environment 

First, run `make setup-sol-cli` to fetch the necessary scripts. Requires `truffle` and `ganache-cli`.

To start a ganache instance:
```
make start-eth
```

Bridge contracts can then be deployed with:
```
make deploy-eth
```

Note: The environment variable `PORT=<port>` can be provided for these commands (default `PORT=8545`)

To build the go bindings for the Ethereum contracts:
```
make bindings
```

## Centrifuge Dev Environment

To fetch, build and run centrifuge-chain run:
```
make start-cent
```

Note: The build process takes a while, but will only run once. It currently uses a modified fork of centrifuge-chain

You can run several commands to interact with the bridge module:

### Emitter Address

You can set and get the emitter address with:
 ```
 make cent-get-emitter
``` 
and 
```
make cent-set-emitter CENT_EMITTER_ADDR=<HEX VALUE>
```
### Whitelist Chain

A chain ID can be whitelisted as a destination with:

```
make cent-whitelist-chain CENT_CHAIN_ID=<HEX VALUE>
```

### Asset Transfer

An asset transfer can be executed with:

```
make cent-asset-tx CENT_CHAIN_ID=<HEX VALUE> CENT_TO=<HEX VALUE> CENT_TOKEN_ID=<HEX VALUE> CENT_METADATA=<HEX VALUE>
```

### Auto Run

Setting an emitter address, whitelisting a chain and submitting an asset tx can be executed using default values with:
```
make cent-auto-run 
```

## Tests

To run the go tests a fresh ganache instance is required (tests depend on deterministic addresses). 
A new instance can be started by running these in seperate terminals:
```
make start-eth
make deploy-eth
```
Go tests can then be run with:
```
make test
```

**Note: Substrate tests are not yet able to be run locally and will fail.**

## Simulations
### Ethereum ERC20 Transfer
Start chain 1 (terminal 1)
```shell
make setup-contracts
make start-eth
```
Start chain 2 (terminal 2)
```shell
PORT=8546 make start-eth
```
Deploy the contracts (terminal 3)
```shell
make deploy_eth && PORT=8546 make deploy-eth
```
Build the latest ChainBridge binary & run it (terminal 3)
```shell
make build
./build/chainbridge --verbosity=trace --config ./scripts/configs/config1.toml --testkey alice
```
Mint & make a deposit (terminal 4)
```shell
node solidity/scripts/cli/index.js --test-only --mint-erc20 --value 100
node solidity/scripts/cli/index.js --test-only --deposit-erc —dest 1
```

Notes: 
- Alice (from the keyring) is always the deployer, if that key changes, then the constants will be different
- Validators start from the keyring and move alphabetically down the list. For example if you specify `--validators 3`, the validators would be `Alice`, `Bob`, `Charlie`. If you said 4, `Dave` would join
- `--test-only` ensures we don't re-deploy the contracts
- `--dest` allows you to specify which chain_id you want to the transfer to go to

#### Debugging
Node script errors:
"Contract not found" or similar:
- Check the deployments in step 3, do the addresses listed there match with the addresses saved in `solidity/scripts/cli/constants.js`? The constants file should be updated accordingly
"Sender doesn't have funds" or similar when executing an erc20 transfer:
- Check that the you ran `--mint <value>` (step 4) if you didn't the account has no tokens to deposit

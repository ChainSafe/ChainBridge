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
Used for substrate key management. Only required if connecting to a substrate chain.

  `make install-subkey`


## Building

`make build`: Builds `chainbridge` in `./build`.

**or**

`make install`: Uses `go install` to add `chainbridge` to your GOBIN.

# Configuration

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

### Ethereum Options

Ethereum chains support the following additional options:

```
bridge = "0x12345..." // Address of the bridge contract (required)
erc20Handler = "0x1234..." // Address of erc20 handler (required)
erc721Handler = "0x1234..." // Address of erc721 handler (required)
genericHandler = "0x1234..." // Address of generic handler (required)
gasPrice = "0x1234"      // Gas price for transactions (default: 20000000000)
gasLimit = "0x1234"      // Gas limit for transactions (default: 6721975)
http = "true"            // Whether the chain connection is ws or http (default: false)
startBlock = "1234" // The block to start processing events from (default: 0)
```

### Substrate Options

Substrate supports the following additonal options:

```
startBlock = "1234" // The block to start processing events from (default: 0)
```

## Blockstore

The blockstore is used to record the last block the relayer processed, so it can pick up where it left off. 

If a `startBlock` option is provided (see [Configuration](#configuration)), then the greater of `startBlock` and the latest block in the blockstore is used at startup.

To disable loading from the blockstore specify the `--fresh` flag. A custom path for the blockstore can be provided with `--blockstore <path>`

## Keystore

ChainBridge requires keys to sign and submit transactions, and to identify each bridge node on chain.

To use secure keys, see `chainbridge accounts --help`. The keystore password can be supplied with the `KEYSTORE_PASSWORD` environment variable.

To import external ethereum keys, such as those generated with geth, use `chainbridge accounts import --ethereum /path/to/key`.

To import private keys as keystores, use `chainbridge account import --privateKey key`.

For testing purposes, chainbridge provides 5 test keys. The can be used with `--testkey <name>`, where `name` is one of `Alice`, `Bob`, `Charlie`, `Dave`, or `Eve`. 

# Chain Implementations

- Ethereum (Solidity): [chainbridge-solidity](https://github.com/ChainSafe/chainbridge-solidity) 

    The Solidity contracts required for chainbridge. Includes deployment and interaction CLI.
    
    The bindings for the contracts live in `bindings/`. To update the bindings modify `scripts/setup-contracts.sh` and then run `make clean && make setup-contracts`

- Substrate: [chainbridge-substrate](https://github.com/ChainSafe/chainbridge-substrate)

    A substrate pallet that can be integrated into a chain, as well as an example pallet to demonstrate chain integration.

# Testing

Unit tests require an ethereum node running on `localhost:8545` and a substrate node running on `localhost:9944`. E2E tests require an additional ethereum node on `localhost:8546`.

See [chainbridge-solidity](https://github.com/chainsafe/chainbridge-solidity) and [chainbridge-substrate-chain](https://github.com/ChainSafe/chainbridge-substrate-chain) for more information.

Go tests can be run with:
```
make test
```
Go tests for end-to-end, ethereum and substrate can be run with
```
make test-e2e
make test-eth
make test-sub
```

# 🌉 <b> Overview </b>

[![Build Status](https://travis-ci.com/ChainSafe/ChainBridge.svg?branch=master)](https://travis-ci.com/ChainSafe/ChainBridge)

## Relevant repos

### ChainBridge

[chainbridge](https://github.com/ChainSafe/ChainBridge)

A Modular Multi-Directional Blockchain Bridge to interact with Multiple Networks; Ethereum, Ethereum Classic, Substrate, Cosmos-SDK based chains.
### Ethereum (Solidity) 

[chainbridge-solidity](https://github.com/ChainSafe/chainbridge-solidity) 

 The Solidity contracts required for chainbridge. Includes deployment and interaction CLI.
    

 The bindings for the contracts live in `bindings/`. To update the bindings modify `scripts/setup-contracts.sh` and then run `make clean && make setup-contracts`

### Substrate
[chainbridge-substrate](https://github.com/ChainSafe/chainbridge-substrate)

A substrate pallet that can be integrated into a chain, as well as an example pallet to demonstrate chain integration.

## Notion

To learn more about ChainBridge check out the [Notion](https://www.notion.so/chainsafe/ChainBridge-Public-Documentation-be78788285f54270b481cfdc811d3eae) page.
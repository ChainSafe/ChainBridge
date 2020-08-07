# Ethereum Chain

The ethereum package contains the logic for interacting with ethereum chains.

There are 3 major components: the [connection](https://github.com/ChainSafe/ChainBridge/blob/master/connections/ethereum/connection.go), the [listener](https://github.com/ChainSafe/ChainBridge/blob/master/chains/ethereum/listener.go), and the [writer](https://github.com/ChainSafe/ChainBridge/blob/master/chains/ethereum/writer.go).
The currently supported transfer types are Fungible (ERC20), Non-Fungible (ERC721), and generic.

## Connection

The connection contains the ethereum RPC client and can be accessed by both the writer and listener.

## Listener

The listener polls for each new block and looks for deposit events in the bridge contract. If a deposit occurs, the listener will fetch additional information from the handler before constructing a message and forwarding it to the router.

## Writer

The writer recieves the message and creates a proposals on-chain. Once a proposal is made, the writer then watches for a finalization event and will attempt to execute the proposal if a matching event occurs. The writer skips over any proposals it has already seen.

## Configuration
### Ethereum Options

Ethereum chains have [configurations](../configuration.md) that also support the following additional options:

```
{
    "bridge": "0x12345..."          // Address of the bridge contract (required)
    "erc20Handler": "0x1234..."     // Address of erc20 handler (required)
    "erc721Handler": "0x1234..."    // Address of erc721 handler (required)
    "genericHandler": "0x1234..."   // Address of generic handler (required)
    "maxGasPrice": "0x1234"            // Gas price for transactions (default: 20000000000)
    "gasLimit": "0x1234"            // Gas limit for transactions (default: 6721975)
    "http": "true"                  // Whether the chain connection is ws or http (default: false)
    "startBlock": "1234"            // The block to start processing events from (default: 0)
}
```


## Chain Implementations
 Ethereum (Solidity): [chainbridge-solidity](https://github.com/ChainSafe/chainbridge-solidity) 

The Solidity contracts required for chainbridge. Includes deployment and interaction CLI.

The bindings for the contracts live in `bindings/`. To update the bindings modify `scripts/setup-contracts.sh` and then run `make clean && make setup-contracts`

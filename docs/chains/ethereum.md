# Chain Implementations
 Ethereum (Solidity): [chainbridge-solidity](https://github.com/ChainSafe/chainbridge-solidity) 

The Solidity contracts required for chainbridge. Includes deployment and interaction CLI.

The bindings for the contracts live in `bindings/`. To update the bindings modify `scripts/setup-contracts.sh` and then run `make clean && make setup-contracts`

#Introduction

The ethereum package contains the logic for interacting with ethereum chains.

There are 3 major components: the connection, the listener, and the writer.
The currently supported transfer types are Fungible (ERC20), Non-Fungible (ERC721), and generic.

Connection

The connection contains the ethereum RPC client and can be accessed by both the writer and listener.

Listener

The listener polls for each new block and looks for deposit events in the bridge contract. If a deposit occurs, the listener will fetch additional information from the handler before constructing a message and forwarding it to the router.

Writer

The writer recieves the message and creates a proposals on-chain. Once a proposal is made, the writer then watches for a finalization event and will attempt to execute the proposal if a matching event occurs. The writer skips over any proposals it has already seen.

# On-Chain Setup (Ethereum)

## Deploy Contracts

You can deploy the contracts to the geth chain like this:

```bash
cb-sol-cli deploy --all --relayerThreshold 1
```

You should see an output such as this:

```bash
================================================================
Url:        http://localhost:8545
Deployer:   0xff93B45308FD417dF303D6515aB04D9e89a750Ca
Gas Limit:   8000000
Gas Price:   20000000
Deploy Cost: 0.0

Options
=======
Chain Id:    0
Threshold:   2
Relayers:    0xff93B45308FD417dF303D6515aB04D9e89a750Ca,0x8e0a907331554AF72563Bd8D43051C2E64Be5d35,0x24962717f8fA5BA3b931bACaF9ac03924EB475a0,0x148FfB2074A9e59eD58142822b3eB3fcBffb0cd7,0x4CEEf6139f00F9F4535Ad19640Ff7A0137708485
Bridge Fee:  0
Expiry:      100

Contract Addresses
================================================================
Bridge:             0x62877dDCd49aD22f5eDfc6ac108e9a4b5D2bD88B
----------------------------------------------------------------
Erc20 Handler:      0x3167776db165D8eA0f51790CA2bbf44Db5105ADF
----------------------------------------------------------------
Erc721 Handler:     0x3f709398808af36ADBA86ACC617FeB7F5B7B193E
----------------------------------------------------------------
Generic Handler:    0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07
----------------------------------------------------------------
Erc20:              0x21605f71845f372A9ed84253d2D024B7B10999f4
----------------------------------------------------------------
Erc721:             0xd7E33e1bbf65dC001A0Eb1552613106CD7e40C31
----------------------------------------------------------------
Centrifuge Asset:   0xc279648CE5cAa25B9bA753dAb0Dfef44A069BaF4
================================================================
```

## Register Resources

```bash
# Register fungible resource ID with erc20 contract
cb-sol-cli bridge register-resource --resourceId "0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00" --targetContract "0x21605f71845f372A9ed84253d2D024B7B10999f4"

# Register non-fungible resource ID with erc721 contract
cb-sol-cli bridge register-resource --resourceId "0x000000000000000000000000000000e389d61c11e5fe32ec1735b3cd38c69501" --targetContract "0xd7E33e1bbf65dC001A0Eb1552613106CD7e40C31" --handler "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E"

# Register generic resource ID
cb-sol-cli bridge register-generic-resource --resourceId "0x000000000000000000000000000000f44be64d2de895454c3467021928e55e01" --targetContract "0xc279648CE5cAa25B9bA753dAb0Dfef44A069BaF4" --handler "0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07" --hash --deposit "" --execute "store(bytes32)"
```

## Specify Token Semantics

To allow for a variety of use cases, the ethereum contracts support both the `transfer` and the `mint/burn` ERC methods.

To simplify these examples we should use `mint/burn`:

```bash
# Register the erc20 contract as mintable/burnable
cb-sol-cli bridge set-burn --tokenContract "0x21605f71845f372A9ed84253d2D024B7B10999f4"

# Register the associated handler as a minter
cb-sol-cli erc20 add-minter --minter "0x3167776db165D8eA0f51790CA2bbf44Db5105ADF"

# Register the erc721 contract as mintable/burnable
cb-sol-cli bridge set-burn --tokenContract "0xd7E33e1bbf65dC001A0Eb1552613106CD7e40C31" --handler "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E"

# Add the handler as a minter
cb-sol-cli erc721 add-minter --minter "0x3f709398808af36ADBA86ACC617FeB7F5B7B193E"
```
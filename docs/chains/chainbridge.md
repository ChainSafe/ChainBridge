# ChainBridge Specification

## Summary

ChainBridge is a modular multi-directional blockchain bridge to allow data and value transfer between any number of
 blockchains. This should enable users to specify a destination blockchain from their source chain, and send data to
  that blockchain for consumption on the destination chain. This could be a token that is locked on ChainA and
   redeemed on ChainB, or an operation that is executed on a destination chain and initiated on the source chain. The bridge
    should maintain a set of relayers that are authorized to make & process transfers across the different blockchains. The bridge design should be modular enough that the addition of a new type of transfer should not require a full re-deployment of the suite of tools, rather small modular upgrades.


## Definitions

### Chain ID

Each chain has a unique 8-bit identifier. We presently define the following chain IDs (subject to change):

| ID | Chain |
| - | - |
| 0 | ETH Mainnet |
| 1 | ETC Mainnet |
| 2 | Centrifuge Network |
| 3 | Aragon Chain |
| .. | .. |

### Deposit Nonce

A nonce must be generated for every transfer to ensure uniqueness. All implementations must track a sequential nonce (unsigned 64-bit integer) for each possible destination chain. This is included as a standard parameter for each transfer. Order is not enforced.

### Resource ID

In order to provide generality, we need some way to associate some action on a source chain to some action on a destination chain. This may express tokenX on chain A is equivalent to tokenY on chain B, or to simply associate that some action performed on chain A should result in some other action occurring on chain B. 

All resource IDs are considered to have a Home Chain. The only strict requirements for Resource IDs is that they must be 32 bytes in length and the least significant byte must contain a chain ID. 

### Transfer Flow

1. User initiates a transfer on the source chain.
2. Relayers observing the chain parse the parameters of the transfer and format them into a message.
3. The message is parsed and then proposed on the destination chain.
4. If the vote threshold is met, the proposal will be executed to finalize the transfer.

After the initiation, a user should not be required to make any additional interactions.

## Transfer Types
In a effort to balance the goals of allowing simple integration and proving generalized transfers, multiple transfer types are defined. Some or all of these may implemented for a chain.

|Event|Description|
|-----|-----------|
|FungibleTransfer| Transfer of fungible assets |
|NonFungibleTransfer| Transfer of non-fungible assets |
|GenericTransfer| Transfer of arbitrary data |


All transfers contain a source chain, destination chain, deposit nonce, resource ID and transfer-specific parameters.

### Fungible
|Field|Type|Description|
|----|----|-----------|
| Amount | 256 bit uint | The total number of assets being transferred |
| Recipient | 32 bytes | The recipient address on the destination chain |

### Non-Fungible
|Field|Type|Description|
|----|----|-----------|
| Token ID | 256 bit uint | The unique identifier for the NFT |
| Recipient | 32 bytes | The recipient address on the destination chain |
| Metadata | variable sized bytes | Any additional data associated to the NFT |

### Generic
|Field|Type|Description|
|----|----|-----------|
| Metadata | variable sized bytes | An opaque payload to transmit |

*Note: Addresses are limited to 32bytes in size, but may be smaller. They must always be compatible with the destination chain.*

## Relayer Set
Each chain implementation must track a set of relayers, and allow updating of the set as necessary. A threshold should also be maintained to define how many relayers must vote for a proposed transfer before is can be executed. For this initial implementation, the relayer set may be controlled by a single party. Multi-signature wallets can be used to distribute risk, if available on the chain.


## Implementation

This sections defines the specifics of the ChainBridge implementation and the requirements for a chain integration.

### Components

#### Chain

A chain is loosely defined as consisting of three major components:

- **Connection**:
A container for on chain interactions. Shared by the listener and writer.
- **Listener**: 
Observes chain state transitions to watch for initiated transfers. When a transfer is encountered it should construct a message and pass it to the router.
- **Writer**:
Responsible for performing on-chain actions. This will parse a proposed transfer from a message and enact it on-chain. 

These vary considerably depending on the chain. As long as the on-chain components are compatible, following the internal message protocol should be sufficient to be compatible with the system. These components are intended for architectural guidance and are only loosely constrained.

#### Message

A message represents a single transfer and its associated parameters. 

```go
type Message struct {
	Source       ChainId   
	Destination  ChainId 
	Type         TransferType
	DepositNonce Nonce
	ResourceId   ResourceId
	Payload      []interface{}
}
```

The `payload` field contains the data for the specific transfer, as defined [above](#transfer-types).

#### Router

The router is responsible for taking messages from a source chain and routing them to their destination chain.

The router provides an interface to allow Listeners to submit constructed messages:
```go
type Router interface {
	Send(message msg.Message) error
}
```

All chains must fulfill a Writer interface to receive messages from the router:
```go
type Writer interface {
	ResolveMessage(message msg.Message) bool
}
```

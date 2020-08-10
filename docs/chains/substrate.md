# Substrate Implementation Specification

- [Events](#events)
- [Inter-Pallet Communication](#inter-pallet-communication)
- [Bridge Account ID & Origin Check](#bridge-account-id--origin-check)
- [Executing Calls](#executing-calls)

The ChainBridge Substrate implementation will consist of a Substrate pallet that can be integrated into a runtime to enable bridging of additional pallet functionality. 

Due to the complexities of the Substrate API we must define some limitations to the supported calls, however the pallet should define a `Proposal` type equivalent to a dispatchable call to theoretically allow for any call to be made.

```rust
pub trait Trait: system::Trait {
    type Proposal: Parameter + Dispatchable<Origin = Self::Origin> + EncodeLike + GetDispatchInfo;
}
```

# Events

To easily distinguish different transfer types we should define three event types:

```rust
/// FungibleTransfer is for relaying fungibles (dest_id, nonce, resource_id, amount, recipient, metadata)
FungibleTransfer(ChainId, DepositNonce, ResourceId, U256, Vec<u8>)

/// NonFungibleTransfer is for relaying NFTS (dest_id, nonce, resource_id, token_id, recipient, metadata)
NonFungibleTransfer(ChainId, DepositNonce, ResourceId, Vec<u8>, Vec<u8>, Vec<u8>)

/// GenericTransfer is for a generic data payload (dest_id, nonce, resource_id, metadata)
GenericTransfer(ChainId, DepositNonce, ResourceId, Vec<u8>)
```

These can be observed by relayers and should provide enough context to construct transfer messages. 

# Inter-Pallet Communication

The ChainBridge pallet is intended to be combined with other pallets to define what is being bridged. To allow for this we must define some methods that other pallets can call to initiate transfers:

```rust
pub fn transfer_fungible(dest_id: ChainId, resource_id: ResourceId, to: Vec<u8>, amount: U256,)
    
pub fn transfer_nonfungible(dest_id: ChainId, resource_id: ResourceId, token_id: Vec<u8>, to: Vec<u8>, metadata: Vec<u8>)
    
pub fn transfer_generic(dest_id: ChainId, resource_id: ResourceId, metadata: Vec<u8>)
```

These should result in the associated event being emitted with the correct parameters.

# Bridge Account ID & Origin Check

To allow the bridge pallet to take ownership of tokens a `ModuleId` should be used to derive an `AccountId`.

A bridge origin check (implementing `EnsureOrigin`) should also be provided. Other pallets should be able to use this to check the origin of call is the bridge pallet, indicating the execution of a proposal.

# Executing Calls

The pallet should support dispatching of arbitrary calls as the result of successful proposal. Resource IDs should be mapped to specific calls to define their behaviour. Relayers will need to resolve resource IDs to calls in order to submit a proposal. The pallet should provide a mapping of resource IDs to method names that can be updated by the admin.

Compatible calls are restrained to the following signature to allow relayers to understand how to construct the calls: 
- Fungible: `Call(origin, recipient: AccountId, amount: u128)`
- Non-Fungible: `Call(origin, recipient: AccountId, tokenId: U256, metadata: Vec<u8>)`
- Generic: `Call(origin, data: Vec<u8>)`


*Note: Calls in substrate are resolved based on a pallet and call index. The pallet index depends on the ordering of pallets in the runtime, and the call index on the ordering of calls in the pallet. As these may change during a runtime upgrade, relayers should use the actual method name string to reference calls* 


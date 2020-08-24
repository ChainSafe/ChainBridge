# Ethereum Implementation Specification

The solidity implementation of ChainBridge should consist of a central Bridge contract, and will delegate specific functionality to [handlers](#handler-contracts). Fungible and non-fungible compatibility should be focused on ERC20 and ERC721 tokens.

## Transfer Flow 

### As Source Chain
1.  Some user calls the `deposit` function on the bridge contract. A `depositRecord` is created on the bridge and a call is delgated to a handler contract specified by the provided `resourceID`.

2. The specified handler's `deposit` function validates the parameters provided by the user. If successful, a `depositRecord` is created on the handler. 
    
3. If the call delegated to the handler is succesful, the bridge emits a `Deposit` event.

4. Relayers parse the `Deposit` event and retrieve the associated `DepositRecord` from the handler to construct a message.


### As Destination Chain
1. A Relayer calls `voteProposal` on the bridge contract. If a `proposal` corresponding with the parameters passed in does not exist, it is created and the Relayer's vote is recorded. If the proposal already exists, the Relayer's vote is simply recorded.

3. Once we have met some vote threshold for a `proposal`, the bridge emits a `ProposalFinalized` event.

4. Upon seeing a `ProposalFinalized` event, Relayers call the`executeDeposit`function on the bridge. `executeDeposit` delegates a call to a handler contract specified by the associated `resourceID`.

5. The specified handler's `executeDeposit` function validates the paramters provided and makes a call to some contract to complete the transfer.


## Bridge Contract

Users and relayers will interact with the `Bridge` contract. This delegates calls to the handler contracts for deposits and executing proposals. 

```
function deposit (uint8 destinationChainID, bytes32 resourceID, bytes calldata data)
```

## Handler Contracts

To provide modularity and break out the necessary contract logic, the implementation uses a notion of handlers. A handler is defined for ERC20, ERC721 and generic transfers. These map directly to the Fungible, Non-Fungible, and generic transfer types.

A handler must fulfill two interfaces:
```
// Will be called by the bridge contract to initiate a transfer
function deposit(uint8 destinationChainID, uint64 depositNonce, address depositer, bytes calldata data)
```

```
// TODO: This would be more aptly named executeProposal
// Called by the bridge contract to complete a transfer
function executeDeposit(bytes calldata data)
```

The `calldata` field is the parameters required for the handler. The exact serialization is defined for each handler.

### ERC20 & ERC721 Handlers

These handlers share a lot of similarities.

These handlers are responsible for transferring ERC assets. They should provide the ability for the bridge to take ownership of tokens and release tokens to execute transfers. 

Different configurations may require different interface interactions. For example, it may make sense to mint and burn a token that is originally from another chain. If supply needs to be controlled, transferring tokens in and out of a reserve may be desired instead. To support either case handlers should associate each resource ID/token contract with one of these:

- `transferFrom()` - The user approves the handler to move the tokens prior to initiating the transfer. The handler will call `transferFrom()` as part of the transfer initiation. For the inverse, the handler will call `transfer()` to release tokens from the handlers ownership.
- `mint()`/`burn()` - The user approves the handler to move the tokens prior to initiating the transfer. The handler will call `burnFrom()` as part of the transfer initiation. For the inverse, the handler will call `mint()` to release tokens to the recipient (and must have privileges to do so).


### ERC20 Handler

#### Calldata for `deposit()`
| Data | Type | Location |
| - | - | - |
| Amount | uint256 | 0 - 31 |
| Recipient Address Length | uint256 | 32 - 63 |
| Recipient Address | bytes | 63 - END |


#### Calldata for `executeDeposit()`
| Data | Type | Location |
| - | - | - |
| Amount | uint256 | 0 - 31 |
| Recipient Address Length | uint256 | 32 - 63 |
| Recipient Address | bytes | 64 - END |

### ERC721 Handler

#### Metadata

The `tokenURI` should be used as the `metadata` field if the contract supports the Metadata extension (interface ID `0x5b5e139f`).


#### Calldata for `deposit()`
| Data | Type | Location |
| - | - | - |
| TokenID | uint256 | 0 - 31 |
| Recipient Address Length | uint256 | 32 - 63 |
| Recipient Address | bytes | 63 - END |

#### Calldata for `executeDeposit()`
| Data | Type | Location |
| - | - | - |
| TokenID | uint256 | 0 - 31 |
| Recipient Address Length | uint256 | 32 - 63 |
| Recipient Address | bytes | 64 - 95 |
| Metadata Length | uint256 | 96 - 127 | 
| Metadata | bytes | 128 - END |


### Generic Handler 

As well as associating a resource ID to a contract address, the generic handler should allow specific functions on those contracts to be used. To allow for this we must:

1. Use [function selectors](https://solidity.readthedocs.io/en/v0.6.4/abi-spec.html#function-selector) to identify functions.
2. Require functions that accept `bytes` as a the only parameter **OR** require the data already be ABI encoded for the function

#### Deposit

In a generic context, a deposit is simply the initiation of a transfer of a piece of data. To (optionally) allow this data to be validated for transfer the deposit mechanism should pass the data to a specified function and proceed with the transfer if the call succeeds (ie. does not revert). A function selector of `0x00` should skip the deposit function call.

#### Execute

An execution function must be specified. When `executeDeposit()` is called on the handler it should pass the `metadata` field to the specified function.


#### Calldata for `deposit()`
| Data | Type | Location |
| - | - | - |
| Metadata Length | uint256 | 0 - 31 |
| Metadata | bytes | 32 - END |

#### Calldata for `execute()`
| Data | Type | Location |
| - | - | - |
| Metadata Length | uint256 | 0 - 31 |
| Metadata | bytes | 32 - END |

## Administration

The contracts should be controlled by an admin account. This should control the relayer set, manage the resource IDs, and specify the handlers. It should also be able to pause and unpause transfers at any times.

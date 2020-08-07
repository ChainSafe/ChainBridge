# Chain Implementations
 Ethereum (Solidity): [chainbridge-solidity](https://github.com/ChainSafe/chainbridge-solidity) 

The Solidity contracts required for chainbridge. Includes deployment and interaction CLI.
    
The bindings for the contracts live in `bindings/`. To update the bindings modify `scripts/setup-contracts.sh` and then run `make clean && make setup-contracts`

# Overview

ChainBridge uses Solidity smart contracts to enable transfers to and from EVM compatible chains. These contracts consist of a core bridge contract (`Bridge.sol`) and a set of handler contracts (`ERC20Handler.sol`, `ERC721Handler.sol`, and `GenericHandler.sol`). The bridge contract is responsible for initiating, voting on, and executing proposed transfers. The handlers are used by the bridge contract to interact with other existing contracts.

# Bridge Contract

### Core Functions

```jsx
function deposit (uint8 destinationChainID, address handler, bytes memory data)
```

This is called by a user to initiate a transfer from this chain to another chain. The bridge will do some internal accounting and then call the handler to complete the deposit. The handler defines the exact structure of `data`.

```jsx
function voteProposal(uint8 originChainID, uint256 depositNonce, bytes32 dataHash) public onlyRelayers
```

Relayers use this method to create or vote on a proposed transfer from another chain destined for this chain. 

```jsx
function executeProposal(uint8 originChainID, uint256 depositNonce, address handler, bytes memory data) public onlyRelayers
```

Once the voting threshold has been met for a proposal, an event is emitted to signal to relayers that the proposal can be executed. Presently, all relayers will attempt to execute a finalized proposal, but only one will succeed. The execution of the proposal is completed by the associated handler.

### Configuration

There are several administrative functions provided by the bridge contract:

```jsx
// Modifies the number of votes required for a proposal to be executed
function adminChangeRelayerThreshold(uint newThreshold) public onlyAdmin
// Add address to relayer set
function adminAddRelayer(address relayerAddress) public onlyAdmin
// Remove address from relayer set
function adminRemoveRelayer(address relayerAddress) public onlyAdmin
// Register an ERC resource ID and token contract
function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) public onlyAdmin
// Register a generic resourceID and contract
function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) public onlyAdmin
// Enable minting/burning on a token contract
function adminSetBurnable(address handlerAddress, address tokenAddress) public onlyAdmin
```

# ERC Handlers

There are two modes of operation for ERC handlers:

## Mint/Burn

As tokens are transferred out of the chain, they are burned. Conversely, as token are transferred into the chain, they are minted. This requires the token contract to provide mint/burn functionality and the respective handler to have privileges on the contract to perform these operations.

## Transfer (Safe)

Deposited tokens are owned by the handler, and tokens transferred into the chain are taken from the handlers account. This requires a sufficient supply to be maintained in the handlers account. Some methods are provided to fund and withdraw from the handler.

By default, the transfer mode is enabled for all registered token contracts. To enable mint/burn a call must be made to the bridge:

```jsx
function adminSetBurnable(address handlerAddress, address tokenAddress) public onlyAdmin
```

 

## Registering Resource IDs

Resource IDs and token contracts can be associated by calling this method on the bridge contract:

```jsx
function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) public onlyAdmin
```
pragma solidity 0.6.4;

/**
 * @title SimpleEmitter
 * @dev Very simple version of the Emitter Contract. It generates an event for a random hash.
 */
contract SimpleEmitter {
	event DepositAsset(address indexed _addr, bytes32 _hash);

<<<<<<< HEAD:on-chain/evm-contracts/contracts/test/SimpleEmitter.sol
	fallback() external {
=======
	fallback () external {
>>>>>>> master:on-chain/evm-contracts/contracts/tests/SimpleEmitter.sol
		bytes32 hash = sha256(abi.encodePacked(msg.sender, block.number));
		emit DepositAsset(msg.sender, hash);
	}
}
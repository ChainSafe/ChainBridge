pragma solidity 0.5.12;

contract Emitter {
	event Transfer(address _addr, bytes32 indexed _hash);

	function () external {
		bytes32 hash = sha256(abi.encodePacked(msg.sender, block.number));
		emit Transfer(msg.sender, hash);
	}
}
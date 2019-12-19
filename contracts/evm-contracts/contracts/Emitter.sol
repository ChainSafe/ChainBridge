pragma solidity 0.5.12;

contract Emitter {
	event Transfer(address _addr, bytes32 _hash);

	function () external {
		bytes32 hash = sha256(abi.encodePacked(msg.sender));
		emit Transfer(msg.sender, hash);
	}
}
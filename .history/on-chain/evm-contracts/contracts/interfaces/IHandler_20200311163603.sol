pragma solidity ^0.5.12;

import "../helpers/Ownable.sol";

/**
 * @title IHandler
 * @dev An interface for handling transfers of data.
 * A receiver contract sends the _data to an arbitrary message.
 */
contract IHandler {
    /**
     * @dev A function that can be generically implemented, this is directly called by the Receiver.
     * eg: The receiver would
     * IHandler(_to).executeTransfer(0, 0, _data)
     */
    function executeTransfer(uint _originChain, bytes memory _data) public;
}
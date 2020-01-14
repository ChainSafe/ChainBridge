pragma solidity 0.5.12;

import "../helpers/Ownable.sol";

interface IHandler {
    function executeTransfer(uint _originChain, bytes calldata _data) external;
}
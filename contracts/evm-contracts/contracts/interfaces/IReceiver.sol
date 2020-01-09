pragma solidity 0.5.12;

import "../helpers/Ownable.sol";

interface IReceiver {
    function executeTransfer(uint _originChain, bytes calldata _data) external;
}
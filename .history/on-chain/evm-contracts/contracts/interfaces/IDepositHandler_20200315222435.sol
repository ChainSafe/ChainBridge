pragma solidity 0.6.4;

interface IDepositHandler {
    function deposit(uint256 depositID, bytes memory data) public;
    function executeDeposit(bytes calldata data) external;
}
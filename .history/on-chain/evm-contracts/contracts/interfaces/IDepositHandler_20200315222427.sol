pragma solidity 0.6.4;

interface IDepositHandler {
    function deposit(uint256 depositID, bytes calldata data) public;
    function executeDeposit(bytes calldata data) external;
}
pragma solidity ^0.5.12;

interface IDepositHandler {
    function executeDeposit(bytes calldata data) external;
}
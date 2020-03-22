pragma solidity 0.6.1;
pragma experimental ABIEncoderV2;

interface IDepositHandler {
    function deposit(uint256 depositID, bytes calldata data) external;
    function executeDeposit(bytes calldata data) external;
}
pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

interface IDepositHandler {
    function deposit(uint256 depositID, bytes calldata data) external;
    function executeDeposit(bytes calldata data) external;
}
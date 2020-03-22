pragma solidity 0.6.1;
pragma experimental ABIEncoderV2;

interface IDepositHandler {
    struct DepositRecord {
        address _originChainTokenAddress;
        address _originChainHandlerAddress;
        uint256 _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        uint256 _amount;
    }

    function getDepositRecord(uint256 depositID) external returns (DepositRecord);
    function deposit(uint256 depositID, bytes calldata data) external;
    function executeDeposit(bytes calldata data) external;
}
pragma solidity 0.6.1;
pragma experimental ABIEncoderV2;

interface IDepositHandler {
    struct DepositRecord {
        address _originChainTokenAddress;
        uint    _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        address _depositer;
        uint    _amount;
    }

    function getDepositRecord(uint256 depositID) external returns (DepositRecord memory || bool);
    function deposit(uint256 depositID, bytes calldata data) external;
    function executeDeposit(bytes calldata data) external;
}
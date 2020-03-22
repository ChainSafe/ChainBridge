pragma solidity 0.6.4;

interface IDepositHandler {
    struct ERC20DepositRecord {
        address _originChainTokenAddress;
        uint    _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        address _depositer;
        uint    _amount;
    }

    struct GenericDepositRecord {
        uint256 _destinationChainID;
        address _destinationRecipientAddress;
        bytes   _metaData;
    }

    function getDepositRecord(uint256 depositID) external returns (ERC20DepositRecord);
    function deposit(uint256 depositID, bytes calldata data) external;
    function executeDeposit(bytes calldata data) external;
}
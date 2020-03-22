pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "../ERC20Safe.sol";
import "../interfaces/IDepositHandler.sol";

contract GenericHandler is IDepositHandler, ERC20Safe {
    address public _bridgeAddress;

    struct DepositRecord {
        address _originChainTokenAddress;
        uint256 _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        address _depositer;
        bytes   _metaData;
    }

    // DepositID => Deposit Record
    mapping (uint256 => DepositRecord) public _depositRecords;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress, "sender must be bridge contract");
        _;
    }

    constructor(address bridgeAddress) public {
        _bridgeAddress = bridgeAddress;
    }

    function getDepositRecord(uint256 depositID) public view returns (DepositRecord memory) {
        return _depositRecords[depositID];
    }

    function deposit(uint256 depositID, bytes memory data) public override _onlyBridge {
        address       originChainTokenAddress;
        uint256       destinationChainID;
        address       destinationRecipientAddress;
        bytes memory  metaData;

        assembly {
            destinationChainID             := mload(add(data, 0x40))
            destinationRecipientAddress    := mload(add(data, 0x80))
            metaData                       := mload(add(data, 0xC0))
            let lenextra := mload(add(0x80, data))
            mstore(0x40, add(0x60, add(metaData, lenextra)))
                calldatacopy(
                metaData,                  // copy to extra
                0xA0,                      // copy from calldata @ 0xA0
                sub(calldatasize(), 0xA0)  // copy size (calldatasize - 0xA0)
            )
        }

        _depositRecords[depositID] = DepositRecord(
            originChainTokenAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            depositer,
            metaData
        );
    }

    // Todo: Implement example of generic deposit
    function executeDeposit(bytes memory data) public override {}
}

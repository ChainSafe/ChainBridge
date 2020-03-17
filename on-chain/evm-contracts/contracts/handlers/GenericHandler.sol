pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "../ERC20Safe.sol";
import "../interfaces/IDepositHandler.sol";

contract GenericHandler is IDepositHandler, ERC20Safe {
    address public _bridgeAddress;

    struct DepositRecord {
        uint256 _destinationChainID;
        address _destinationRecipientAddress;
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
        uint256       destinationChainID;
        address       destinationRecipientAddress;
        bytes memory  metaData;

        assembly {
            destinationChainID             := mload(add(data, 0x20))
            destinationRecipientAddress    := mload(add(data, 0x40))
            metaData                       := mload(add(data, 0x60))
            let lenExtra := mload(add(data, 0x80))
            mstore(0x60, add(0x60, add(metaData, lenExtra)))
                calldatacopy(
                metaData,                  // copy to extra
                0xA0,                      // copy from calldata @ 0xA0
                sub(calldatasize(), 0xA0)  // copy size (calldatasize - 0xA0)
            )
        }

        _depositRecords[depositID] = DepositRecord(
            destinationChainID,
            destinationRecipientAddress,
            metaData
        );
    }

    // Todo: Implement example of generic deposit
    function executeDeposit(bytes memory data) public override {}
}

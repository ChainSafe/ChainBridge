pragma solidity 0.6.4;

import "../interfaces/IDepositHandler.sol";

contract CentrifugeAssetHandler is IDepositHandler {
    address public _bridgeAddress;

    enum AssetDepositStatus { Uninitialized, Active, Confirmed }

    // metaDataHash => AssetDepositStatus
    mapping(bytes32 => AssetDepositStatus) _assetDepositStatuses;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress, "sender must be bridge contract");
        _;
    }

    constructor(address bridgeAddress) public {
        _bridgeAddress = bridgeAddress;
    }

    function depositAsset(bytes32 metaDataHash) public _onlyBridge {
        require(_assetDepositStatuses[metaDataHash] == AssetDepositStatus.Uninitialized,
        "asset has already been deposited and cannot be changed");
        _assetDepositStatuses[metaDataHash] = AssetDepositStatus.Active;
    }

    function executeDeposit(bytes memory data) public override {
        bytes32 metaDataHash;

        assembly {
            metaDataHash := mload(add(data, 0x20))
        }

        require(_assetDepositStatuses[metaDataHash] == AssetDepositStatus.Active, "asset hasn't been deposited or has already been finalized");
        _assetDepositStatuses[metaDataHash] = AssetDepositStatus.Confirmed;
    }
}

pragma solidity 0.6.4;

interface ICentrifugeAssetHandler {
    function depositAsset(bytes32 metaDataHash) external;
    function executeDeposit(bytes calldata data) external;
}
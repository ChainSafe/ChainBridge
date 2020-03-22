pragma solidity ^0.5.12;
pragma experimental ABIEncoderV2;

contract BridgeAsset {

  event AssetStored(
    bytes32 indexed asset
  );

  uint8 min_count; // n confirmations * 10

  // unconfirmed - (value > 0)
  // confirmed - (value = 1)
  mapping (bytes32 => uint8) public assets;

  constructor(uint8 mc) public {
    min_count = mc;
  }

  function store(bytes32 asset) public { // Add OnlyOperator modifier
    require(assets[asset] != 1, "Asset cannot be changed once confirmed");

    assets[asset] += 10;

    if (assets[asset] == min_count) {
      assets[asset] = 1;
      emit AssetStored(asset);
    }
  }

  function isAssetValid(bytes32 asset) external view returns (bool) {
    if (assets[asset] == 1) {
      return true;
    }

    return false;
  }

}
pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "../ERC721Safe.sol";
import "../interfaces/IDepositHandler.sol";
import "../erc/ERC721/ERC721Mintable.sol";

contract ERC721Handler is IDepositHandler, ERC721Safe {
    address public _bridgeAddress;

    struct DepositRecord {
        address _originChainTokenAddress;
        uint256 _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        address _depositer;
        uint256 _tokenID;
        bytes   _metaData;
    }

    // DepositID => Deposit Record
    mapping (uint256 => DepositRecord) public _depositRecords;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress, "seder must be bridge contract");
        _;
    }

    constructor(address bridgeAddress) public {
        _bridgeAddress = bridgeAddress;
    }

    function getDepositRecord(uint256 depositID) public view returns (DepositRecord memory) {
        return _depositRecords[depositID];
    }

    function deposit(uint256 depositID, bytes memory data) public override _onlyBridge {
        address      originChainTokenAddress;
        uint256      destinationChainID;
        address      destinationChainHandlerAddress;
        address      destinationRecipientAddress;
        address      depositer;
        uint256      tokenID;
        bytes memory metaData;

        assembly {
            originChainTokenAddress        := mload(add(data, 0x20))
            destinationChainID             := mload(add(data, 0x40))
            destinationChainHandlerAddress := mload(add(data, 0x60))
            destinationRecipientAddress    := mload(add(data, 0x80))
            depositer                      := mload(add(data, 0xA0))
            tokenID                        := mload(add(data, 0xC0))
            metaData                       := mload(add(data, 0xE0))
            let lenextra := mload(add(0x80, data))
            mstore(0x40, add(0x60, add(metaData, lenextra)))
                calldatacopy(
                metaData,                  // copy to extra
                0xA0,                      // copy from calldata @ 0xA0
                sub(calldatasize(), 0xA0)    // copy size (calldatasize - 0xA0)
            )
        }

        lockERC721(originChainTokenAddress, depositer, address(this), tokenID);

        _depositRecords[depositID] = DepositRecord(
            originChainTokenAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            depositer,
            tokenID,
            metaData
        );
    }

    function executeDeposit(bytes memory data) public override {
        address      destinationChainTokenAddress;
        address      destinationRecipientAddress;
        uint256      tokenID;
        bytes memory metaData;

        assembly {
            destinationChainTokenAddress := mload(add(data, 0x20))
            destinationRecipientAddress  := mload(add(data, 0x40))
            tokenID                      := mload(add(data, 0x60))
            metaData                     := mload(0x40)
            let lenextra                 := mload(add(0x80, data))
            mstore(0x40, add(0x60, add(metaData, lenextra)))
                calldatacopy(
                metaData,                  // copy to metaData
                0xA0,                      // copy from calldata @ 0xA0
                sub(calldatasize(), 0xA0)    // copy size (calldatasize - 0xA0)
            )
        }

        ERC721Mintable erc721 = ERC721Mintable(destinationChainTokenAddress);
        erc721.safeMint(destinationRecipientAddress, tokenID, metaData);
    }

    function withdraw(address tokenAddress, address recipient, uint tokenID) public _onlyBridge {
        releaseERC721(tokenAddress, address(this), recipient, tokenID);
    }
}

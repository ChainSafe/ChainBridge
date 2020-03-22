pragma solidity 0.6.1;
pragma experimental ABIEncoderV2;

import "../interfaces/IERC721Handler.sol";
import "../ERC721Safe.sol";
import "../interfaces/IDepositHandler.sol";
import "../erc/ERC721/ERC721Mintable.sol";

contract ERC721Handler is IERC721Handler, IDepositHandler, ERC721Safe {
    address public _bridgeAddress;

    struct DepositRecord {
        address _originChainTokenAddress;
        uint    _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        uint    _tokenID;
        bytes   _metaData;
    }

    // DepositID => Deposit Record
    mapping (uint256 => DepositRecord) public _depositRecords;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress, "sneder must be bridge contract");
        _;
    }

    constructor(address bridgeAddress) public {
        _bridgeAddress = bridgeAddress;
    }

    function getDepositRecord(uint256 depositID) public returns (DepositRecord memory) {
        return _depositRecords[depositID];
    }

    function depositERC721(address tokenAddress, address owner, uint tokenID) public _onlyBridge {
        lockERC721(tokenAddress, owner, address(this), tokenID);
    }

    function deposit(uint2546 depositID, bytes memory data) public _onlyBridge {
        address originChainTokenAddress;
        uint256 destinationChainID;
        address destinationChainHandlerAddress;
        address destinationRecipientAddress;
        address depositer;
        uint256 tokenID;
        bytes   metaData

        assembly {
            originChainTokenAddress        := mload(add(data, 0x20))
            destinationChainID             := mload(add(data, 0x40))
            destinationChainHandlerAddress := mload(add(data, 0x60))
            destinationRecipientAddress    := mload(add(data, 0x80))
            depositer                      := mload(add(data, 0xA0))
            tokenID                        := mload(add(data, 0xC0))
            metaData                       := mload(add(data, 0xE0))
        }

        lockERC20(originChainTokenAddress, depositer, address(this), amount);

        _depositRecords[depositID] = DepositRecord(
            originChainTokenAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            depositer,
            amount
        );
    }

    function executeDeposit(bytes memory data) public {
        address destinationChainTokenAddress;
        address destinationRecipientAddress;
        uint tokenID;
        bytes memory extraData;

        assembly {
            destinationChainTokenAddress := mload(add(data, 0x20))
            destinationRecipientAddress := mload(add(data, 0x40))
            tokenID := mload(add(data, 0x60))
            extraData := mload(0x40)
            let lenextra := mload(add(0x80, data))
            mstore(0x40, add(0x60, add(extraData, lenextra)))
                calldatacopy(
                extraData,                 // copy to extra
                0xA0,                      // copy from calldata @ 0xA0
                sub(calldatasize, 0xA0)    // copy size (calldatasize - 0xA0)
            )
        }

        ERC721Mintable erc721 = ERC721Mintable(destinationChainTokenAddress);
        erc721.safeMint(destinationRecipientAddress, tokenID, extraData);
    }

    function withdrawERC721(address tokenAddress, address recipient, uint tokenID) public _onlyBridge {
        releaseERC721(tokenAddress, address(this), recipient, tokenID);
    }
}

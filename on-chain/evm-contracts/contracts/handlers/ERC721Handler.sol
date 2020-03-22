pragma solidity 0.6.4;

import "../interfaces/IERC721Handler.sol";
import "../ERC721Safe.sol";
import "../interfaces/IDepositHandler.sol";
import "../erc/ERC721/ERC721Mintable.sol";

contract ERC721Handler is IERC721Handler, IDepositHandler, ERC721Safe {
    address public _bridgeAddress;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress, "sender must be bridge contract");
        _;
    }

    constructor(address bridgeAddress) public {
        _bridgeAddress = bridgeAddress;
    }

    function depositERC721(address tokenAddress, address owner, uint tokenID) public override _onlyBridge {
        lockERC721(tokenAddress, owner, address(this), tokenID);
    }

    function executeDeposit(bytes memory data) public override {
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
                sub(calldatasize(), 0xA0)    // copy size (calldatasize - 0xA0)
            )
        }

        ERC721Mintable erc721 = ERC721Mintable(destinationChainTokenAddress);
        erc721.safeMint(destinationRecipientAddress, tokenID, extraData);
    }

    function withdrawERC721(address tokenAddress, address recipient, uint tokenID) public override _onlyBridge {
        releaseERC721(tokenAddress, address(this), recipient, tokenID);
    }
}

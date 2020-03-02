pragma solidity ^0.5.12;

import "../interfaces/IERC721Handler.sol";
import "../ERC721Safe.sol";
import "../interfaces/IDepositHandler.sol";
import "../erc/ERC721/ERC721Mintable.sol";

contract ERC721Handler is IERC721Handler, IDepositHandler, ERC721Safe {
    address public _bridgeAddress;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress);
        _;
    }

    constructor(address bridgeAddress) public {
        _bridgeAddress = bridgeAddress;
    }

    function depositERC721(address tokenAddress, address owner, uint tokenID) public _onlyBridge {
        lockERC721(tokenAddress, owner, address(this), tokenID);
    }

    function executeDeposit(bytes memory data) public {
        address destinationChainTokenAddress;
        address destinationRecipientAddress;
        uint tokenID;
        bytes memory extraData;

        assembly {
            destinationChainTokenAddress := mload(add(data, 32))
            destinationRecipientAddress := mload(add(data, 64))
        }

        ERC721Mintable erc721 = ERC721Mintable(destinationChainTokenAddress);
        erc721.safeMint(destinationRecipientAddress, tokenID, extraData);
    }

    function withdrawERC721(address tokenAddress, address recipient, uint tokenID) public _onlyBridge {
        releaseERC721(tokenAddress, address(this), recipient, tokenID);
    }
}

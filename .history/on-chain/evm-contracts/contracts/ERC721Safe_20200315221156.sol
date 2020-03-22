pragma solidity 0.6.4;

import "./helpers/SafeMath.sol";
import "./erc/ERC721/IERC721.sol";

contract ERC721Safe {
    using SafeMath for uint;

    // ERC721 contract => amount of tokens owned by Safe
    mapping(address => uint) public _balances;

    function lockERC721(address tokenAddress, address owner, address recipient, uint tokenID) internal {
        IERC721 erc721 = IERC721(tokenAddress);
        erc721.transferFrom(owner, recipient, tokenID);

        _balances[tokenAddress] = _balances[tokenAddress].add(1);
    }

    function releaseERC721(address tokenAddress, address owner, address recipient, uint tokenID) internal {
        IERC721 erc721 = IERC721(tokenAddress);
        erc721.transferFrom(owner, recipient, tokenID);

        _balances[tokenAddress] = _balances[tokenAddress].sub(1);
    }
}

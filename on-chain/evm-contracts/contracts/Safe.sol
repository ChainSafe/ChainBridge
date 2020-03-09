pragma solidity ^0.5.12;

import "./helpers/SafeMath.sol";
import "./erc/ERC20/ERC20.sol";
import "./erc/ERC721/ERC721.sol";
import "./interfaces/ISafe.sol";

/**
 * @title Safe
 * @dev The safe is used to take custody, and give custody of ERC20 based tokens.
 */
contract Safe is ISafe {
    using SafeMath for uint;

    address Owner;

    mapping(address => uint) public balances;

    modifier onlyOwner {
        require(msg.sender == Owner);
        _;
    }

    constructor(address _owner) public {
        Owner = _owner;
    }

    function lockErc(address _tokenAddress, uint _value, address _to, address _from) internal {
        balances[_tokenAddress] = balances[_tokenAddress].add(_value);

        // Lock tokens
        ERC20 token = ERC20(_tokenAddress);
        token.transferFrom(_from, _to, _value);
    }

    function lockNFT(address _tokenAddress, address _to, address _from, uint _tokenId) internal {
        ERC721 nft = ERC721(_tokenAddress);
        nft.transferFrom(_from, _to, _tokenId);
    }

    function releaseErc(address _tokenAddress, uint _value, address _to) public onlyOwner {
        require(balances[_tokenAddress] >= _value, "Withdrawal amount is too high!");
        balances[_tokenAddress] = balances[_tokenAddress].sub(_value);

        // Transfer funds to user
        ERC20 token = ERC20(_tokenAddress);
        token.transfer(_to, _value);
    }
    
    function releaseNFT(address _tokenAddress, address _to, uint _tokenId) public onlyOwner {
        ERC721 nft = ERC721(_tokenAddress);
        nft.transferFrom(address(this), _to, _tokenId);
    }
    
    function owner() public view returns (address) {
       return Owner;
    }
}

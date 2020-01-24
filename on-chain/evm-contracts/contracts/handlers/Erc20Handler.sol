pragma solidity 0.5.12;

import "../erc/ERC20/ERC20Mintable.sol";
import "../helpers/Ownable.sol";

/**
 * @dev A handler for the ERC20 token standard.
 */
contract ERC20Handler is Ownable {
    // ChainId originChainAddress => currentChainAddress
    mapping(uint => mapping(address => address)) public whitelist;

    // ChainId originChainAddress => currentChainAddress
    mapping(uint => mapping(address => address)) public synthetics;
    
    function executeTransfer(uint _originChain, address _originTokenAddress, address _to, uint _value) public onlyOwner {
        if (whitelist[_originChain][_originTokenAddress] != address(0)) {
            // Token is whitelisted
            // tokenAddress might not be needed
            address tokenAddress = whitelist[_originChain][_originTokenAddress];
            ERC20Mintable token = ERC20Mintable(tokenAddress);
            // Mint new tokens
            token.mint(_to, _value);
        } else {
            if (synthetics[_originChain][_originTokenAddress] != address(0)) {
                // token exists
                // tokenAddress might not be needed 
                address tokenAddress = synthetics[_originChain][_originTokenAddress]; 
                ERC20Mintable token = ERC20Mintable(tokenAddress);
                // Mint new contract
                token.mint(_to, _value);
            } else {
                // Token doesn't exist
                ERC20Mintable token = new ERC20Mintable();
                // Create a relationship between the originAddress and the synthetic
                synthetics[_originChain][_originTokenAddress] = address(token);
                // Mint new tokens
                token.mint(_to, _value);
            }
        }
    }
}
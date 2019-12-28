pragma solidity 0.5.12;

import "./Ownable";

contract Erc20Receiver is Ownable {

    // ChainId originChainAddress => currentChainAddress
    mapping(uint => mapping(address => address)) public whitelist;

    // ChainId originChainAddress => currentChainAddress
    mapping(uint => mapping(address => address)) public synthetics;

    function transfer(uint _originChain, address _originTokenAddress, address _to, uint _value) public onlyOwner {
        if (whitelist[_originChain][_originTokenAddress] != address(0)) {
            // Token is whitelisted
            /**
             * TODO
             * How to represent native fuel? Perhaps address(0) could represent that?
             * interact with safe?
             */
        } else {
            if (synthetics[_originChain][_originTokenAddress] != address(0)) {
                // token exists
                ERC20Synthetic token = ERC20Synthetic(synthetics[_originChain][_originTokenAddress]);
                token.mint(_to, _value);
            } else {
                // Token doesn't exist
                ERC20Synthetic token = new ERC20Synthetic();
                // Create a relationship between the originAddress and the synthetic
                synthetics[_originChain][_orginTokenAddress] = token.address;
                token.mint(_to, _value);
            }
        }
    }
}
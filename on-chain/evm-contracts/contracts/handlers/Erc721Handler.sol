pragma solidity 0.5.12;

import "../erc/ERC721/ERC721Mintable.sol";
import "../interfaces/IHandler.sol";
import "../helpers/Ownable.sol";

/**
 * @dev A handler for the ERC71 token standard.
 */
contract ERC721Handler is Ownable, IHandler {
    // ChainId originChainAddress => currentChainAddress
    mapping(uint => mapping(address => address)) public whitelist;

    // ChainId originChainAddress => currentChainAddress
    mapping(uint => mapping(address => address)) public synthetics;

    function executeTransfer(uint _originChain, bytes calldata _data) public onlyOwner {
        // Decodes input bytes for _data struct
        //     address _originAddress   @ 0x44
        //     address _to              @ 0x64
        //     uint _value              @ 0x84
        //     bytes _extraData         @ 0xA4 ... arbitrary length
        assembly {
            // get addresses from calldata
            origin := mload(add(0x20, _data))
            to := mload(add(0x40, _data))

            // get uint from calldata
            tokenId := mload(add(0x60, _data))

            // get arbitrary length bytes from calldata
            extraData := mload(0x40)
            let lenextra := mload(add(0x80, _data))
            mstore(0x40, add(0x60, add(extra, lenextra)))
            calldatacopy(
                    extra,                     // copy to extra
                    0xA4,                      // copy from calldata @ 0xA4
                    sub(calldatasize, 0xA4)    // copy size (calldatasize - 0xA4)
            )
        }


        if (whitelist[_originChain][_originTokenAddress] != address(0)) {
            // Token is whitelisted
            // tokenAddress might not be needed
            address tokenAddress = whitelist[_originChain][_originTokenAddress];
            ERC721Mintable token = ERC721Mintable(tokenAddress);
            // Mint new tokens
            token.safeMint(to, tokenId, extraData);
        } else {
            if (synthetics[_originChain][_originTokenAddress] != address(0)) {
                // token exists
                // tokenAddress might not be needed
                address tokenAddress = synthetics[_originChain][_originTokenAddress];
                ERC721Mintable token = ERC721Mintable(tokenAddress);
                // Mint new contract
                token.safeMint(to, tokenId, extraData);
            } else {
                // Token doesn't exist
                ERC721Mintable token = new ERC721Mintable();
                // Create a relationship between the originAddress and the synthetic
                synthetics[_originChain][_originTokenAddress] = address(token);
                // Mint new tokens
                token.safeMint(to, tokenId, extraData);
            }
        }
    }
}
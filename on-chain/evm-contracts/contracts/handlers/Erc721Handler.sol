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

    function executeTransfer(uint _originChain, bytes memory _data) public onlyOwner {
        // Decodes input bytes for _data struct
        //     address _originAddress   @ 0x44
        //     address _to              @ 0x64
        //     uint _value              @ 0x84
        //     bytes _extraData         @ 0xA4 ... arbitrary length
        
        bytes memory _extraData;
        address _to;
        uint _tokenId;
        address _originTokenAddress;
        
        
        // NOTE: test using calldatacopy instead of mlad for all variables and manually setting the offset 
        // for the arbitrary length _extraData it might be more efficient
        assembly {
            // get addresses from calldata
            _originTokenAddress := mload(add(0x20, _data))
            _to := mload(add(0x40, _data))

            // get uint from calldata
            _tokenId := mload(add(0x60, _data))

            // get arbitrary length bytes from calldata
            _extraData := mload(0x40)
            let lenextra := mload(add(0x80, _data))
            mstore(0x40, add(0x60, add(_extraData, lenextra)))
            calldatacopy(
                    _extraData,                // copy to extra
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
            token.safeMint(_to, _tokenId, _extraData);
        } else {
            if (synthetics[_originChain][_originTokenAddress] != address(0)) {
                // token exists
                // tokenAddress might not be needed
                address tokenAddress = synthetics[_originChain][_originTokenAddress];
                ERC721Mintable token = ERC721Mintable(tokenAddress);
                // Mint new contract
                token.safeMint(_to, _tokenId, _extraData);
            } else {
                // Token doesn't exist
                ERC721Mintable token = new ERC721Mintable();
                // Create a relationship between the originAddress and the synthetic
                synthetics[_originChain][_originTokenAddress] = address(token);
                // Mint new tokens
                token.safeMint(_to, _tokenId, _extraData);
            }
        }
    }
}
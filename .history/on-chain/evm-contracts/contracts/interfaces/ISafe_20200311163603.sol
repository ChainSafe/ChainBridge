pragma solidity ^0.5.12;

/**
 * @title ISafe
 * @dev an interface for the Safe contract, it is used to lock and release ERC based tokens.
 */
interface ISafe {
    /**
     * @dev Takes custody of an ERC token
     */
    // function lockErc(address _tokenAddress, uint _value, address _to, address _from) external;
    
    // function lockNFT(address _tokenAddress, address _to, address _from, uint _tokenId) external;

    /**
     * @dev Transfers a custodied ERC token to a user.
     */
    function releaseErc(address _tokenAddress, uint _value, address _to) external;
    
    function releaseNFT(address _tokenAddress, address _to, uint _tokenId) external;
}
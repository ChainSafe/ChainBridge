pragma solidity 0.5.12;

// this is an active WIP
/**
 * @title ISafe
 * @dev an interface for the Safe contract, it is used to lock and release ERC based tokens.
 */
interface ISafe {
    /**
     * @dev Takes custody of an ERC token
     */
    function lock(address _tokenAddress, uint _value, address _to, address _from) external;

    /**
     * @dev Transfers a custodied ERC token to a user.
     */
    function release(address _tokenAddress, uint _value, address _to) external;
}
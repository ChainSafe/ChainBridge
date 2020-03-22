pragma solidity 0.6.4;

import "../erc/ERC20/ERC20.sol";
import "../erc/ERC20/ERC20Mintable.sol";

/**
 * @title SimpleToken
 * @dev Very simple ERC20 Token example, where all tokens are pre-assigned to the creator.
 * Note they can later distribute these tokens as they wish using `transfer` and other
 * `ERC20` functions.
 */
contract SimpleToken is ERC20, ERC20Mintable {

    /**
     * @dev Constructor
     */
    constructor ()
        ERC20Mintable() 
        public {
            
        }
}
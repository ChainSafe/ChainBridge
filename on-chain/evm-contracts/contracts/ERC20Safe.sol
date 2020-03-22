pragma solidity 0.6.4;

import "./helpers/SafeMath.sol";
import "./erc/ERC20/IERC20.sol";

contract ERC20Safe {
    using SafeMath for uint;

    // ERC20 contract => amount of tokens owned by Safe
    mapping(address => uint) public _balances;

    function lockERC20(address tokenAddress, address owner, address recipient, uint amount) internal {
        IERC20 erc20 = IERC20(tokenAddress);
        erc20.transferFrom(owner, recipient, amount);

        _balances[tokenAddress] = _balances[tokenAddress].add(amount);
    }

    function releaseERC20(address tokenAddress, address owner, address recipient, uint amount) internal {
        IERC20 erc20 = IERC20(tokenAddress);
        erc20.transferFrom(owner, recipient, amount);

        _balances[tokenAddress] = _balances[tokenAddress].sub(amount);
    }
}
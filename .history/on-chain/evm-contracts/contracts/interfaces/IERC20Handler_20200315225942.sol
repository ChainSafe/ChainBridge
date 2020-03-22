pragma solidity 0.6.4;

interface IERC20Handler {
    function deposit(address tokenAddress, address owner, uint amount) external;
    function withdrawERC20(address tokenAddress, address recipient, uint amount) external;
}

pragma solidity 0.6.4;

interface IERC20Handler {
    function depositERC20(address tokenAddress, address owner, uint amount) external;
    function withdrawERC20(address tokenAddress, address recipient, uint amount) external;
}

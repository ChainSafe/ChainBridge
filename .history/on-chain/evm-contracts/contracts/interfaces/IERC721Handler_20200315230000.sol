pragma solidity 0.6.4;

interface IERC721Handler {
    function deposit(address tokenAddress, address owner, uint tokenID) external;
    function withdraw(address tokenAddress, address recipient, uint tokenID) external;
}

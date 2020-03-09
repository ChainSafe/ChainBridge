pragma solidity ^0.5.12;

interface IERC721Handler {
    function depositERC721(address tokenAddress, address owner, uint tokenID) external;
    function withdrawERC721(address tokenAddress, address recipient, uint tokenID) external;
}

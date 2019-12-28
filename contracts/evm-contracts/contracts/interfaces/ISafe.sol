pragma solidity 0.5.12;

interface ISafe {
    function lock(address _tokenAddress, uint _value, address _to) public;
    function release(address _tokenAddress, uint _value, address _to) public;
}
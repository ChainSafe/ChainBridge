pragma solidity 0.5.12;

interface IReceiver {
    function transfer(uint _originChain, address _originTokenAddress, address _to, uint _value) external;
}
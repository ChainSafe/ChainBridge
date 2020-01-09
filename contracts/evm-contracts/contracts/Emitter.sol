pragma solidity 0.5.12;

import "./erc/ERC20.sol";
import "./Safe.sol";

contract Emitter is Safe {

    // ChainId => number of deposits
    mapping(uint => uint) DepositCounts;

    event GenericTransfer(uint _destChain, uint _depositId, address _destAddress, bytes _data);
    event ERCTransfer(uint _destChain, uint _depositId, address _to, uint _amount, address _tokenAddress);

    function deposit(uint _destChain, address _destAddress, bytes memory _data) public {
        // Incremnet deposit
        DepositCounts[_destChain]++;
        uint depositId = DepositCounts[_destChain];
        emit GenericTransfer(_destChain, depositId, _destAddress, _data);
    }

    function genericErc(
        uint _destChain,
        uint _value,
        address _to,
        address _from,
        address _tokenAddress
    ) public {
        // Incremnet deposit
        DepositCounts[_destChain]++;
        uint depositId = DepositCounts[_destChain];

        // Lock funds
        lock(_tokenAddress, _value, _to, _from);

        emit ERCTransfer(_destChain, depositId, _to, _value, _tokenAddress);
    }
}
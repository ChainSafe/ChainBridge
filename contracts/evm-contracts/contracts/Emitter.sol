pragma solidity 0.5.12;

import "./Safe.sol";

/**
 * @title Emitter
 * @dev The emitter is found on the origin chain,it is the on-chain
 * user interacted portion of the bridge, and it prepares an asset for
 * transfer. It is responsible for generating the necessary information
 * for the bridge validators to successfully perform a transfer.
 */
 contract Emitter is Safe {

    // ChainId => number of deposits
    mapping(uint => uint) DepositCounts;

    event GenericTransfer(uint _destChain, uint _depositId, address _destAddress, bytes _data);
    event ERCTransfer(uint _destChain, uint _depositId, address _to, uint _amount, address _tokenAddress);
    
    constructor()
        Safe(address(this)) 
        public {}
    
    function deposit(uint _destChain, address _destAddress, bytes memory _data) public {
        // Incremnet deposit
        DepositCounts[_destChain]++;
        uint depositId = DepositCounts[_destChain];
        emit GenericTransfer(_destChain, depositId, _destAddress, _data);
    }
    
    function depositGenericErc(
        uint _destChain,
        uint _value,
        address _to,
        address _tokenAddress
    ) public {
        // Incremnet deposit
        DepositCounts[_destChain]++;
        uint depositId = DepositCounts[_destChain];

        // Lock funds
        super.lockErc(_tokenAddress, _value, address(this), msg.sender);

        emit ERCTransfer(_destChain, depositId, _to, _value, _tokenAddress);
    }
}

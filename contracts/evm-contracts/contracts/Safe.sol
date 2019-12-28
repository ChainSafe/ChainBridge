pragma solidity 0.5.12;

import "./SafeMath";
import "./erc/ERC20";
import "./Ownable";
import "./interfaces/ISafe.sol";

contract Safe is ISafe, Ownable {
    using SafeMath for uint;

    mapping(address => uint) balances;

    function lock(address _tokenAddress, uint _value, address _to, address _from) public onlyowner {
        balances[_tokenAddress] = balances[_tokenAddress].add(_value);
        
        // Lock tokens
        ERC20 token = ERC20(_tokenAddress);
        token.transferFrom(_from, this.address, _value);
    }
    
    function release(address _tokenAddress, uint _value, address _to) public onlyowner {
        require(balances[_tokenAddress] >= _value, "Withdrawal amount is too high!");
        balances[_tokenAddress] = balances[_tokenAddress].sub(_value);

        // Transfer funds to user
        ERC20 token = ERC20(_tokenAddress);
        token.transfer(_to, _value)
    }
}
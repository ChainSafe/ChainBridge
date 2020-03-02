pragma solidity ^0.5.12;

import "../interfaces/IERC20Handler.sol";
import "../ERC20Safe.sol";
import "../erc/ERC20/ERC20Mintable.sol";

contract ERC20Handler is IERC20Handler, ERC20Safe {
    address public _bridgeAddress;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress);
        _;
    }

    constructor(address bridgeAddress) ERC20Safe() public {
        _bridgeAddress = bridgeAddress;
    }

    function depositERC20(address tokenAddress, address owner, uint amount) public _onlyBridge {
        lockERC20(tokenAddress, owner, address(this), amount);
    }

    function executeDeposit(bytes memory data) public {
        address destinationChainTokenAddress;
        address destinationRecipientAddress;
        uint amount;

        assembly {
            destinationChainTokenAddress := mload(add(data, 32))
            destinationRecipientAddress := mload(add(data, 64))
            amount := mload(add(data, 96))
        }

        ERC20Mintable erc20 = ERC20Mintable(destinationChainTokenAddress);
        erc20.mint(destinationRecipientAddress, amount);
    }
}

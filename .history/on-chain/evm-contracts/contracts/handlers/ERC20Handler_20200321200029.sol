pragma solidity 0.6.4;

import "../interfaces/IERC20Handler.sol";
import "../ERC20Safe.sol";
import "../erc/ERC20/ERC20Mintable.sol";
import "../interfaces/IDepositHandler.sol";

contract ERC20Handler is IERC20Handler, IDepositHandler, ERC20Safe {
    address public _bridgeAddress;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress);
        _;
    }

    constructor(address bridgeAddress) public {
        _bridgeAddress = bridgeAddress;
    }

    function depositERC20(address tokenAddress, address owner, uint amount) public override _onlyBridge {
        lockERC20(tokenAddress, owner, address(this), amount);
    }

    function executeDeposit(bytes memory data) public override {
        address destinationChainTokenAddress;
        address destinationRecipientAddress;
        uint amount;

        assembly {
            destinationChainTokenAddress := mload(add(data, 0x20))
            destinationRecipientAddress := mload(add(data, 0x40))
            amount := mload(add(data, 0x60))
        }

        ERC20Mintable erc20 = ERC20Mintable(destinationChainTokenAddress);
        erc20.mint(destinationRecipientAddress, amount);
    }

    function withdrawERC20(address tokenAddress, address recipient, uint amount) public override _onlyBridge {
        releaseERC20(tokenAddress, address(this), recipient, amount);
    }
}

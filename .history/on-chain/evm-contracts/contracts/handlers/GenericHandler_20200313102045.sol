pragma solidity 0.6.1;
pragma experimental ABIEncoderV2;

import "../interfaces/IERC20Handler.sol";
import "../ERC20Safe.sol";
import "../erc/ERC20/ERC20Mintable.sol";
import "../interfaces/IDepositHandler.sol";

contract ERC20Handler is IERC20Handler, IDepositHandler, ERC20Safe {
    address public _bridgeAddress;

    struct DepositRecord {
        address _originChainTokenAddress;
        address _originChainHandlerAddress;
        uint    _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        bytes   _data;
    }

    // DepositID => Deposit Record
    mapping (uint256 => DepositRecord) public _depositRecords;

    modifier _onlyBridge() {
        require(msg.sender == _bridgeAddress, "sender must be bridge contract");
        _;
    }

    constructor(address bridgeAddress) public {
        _bridgeAddress = bridgeAddress;
    }

    function getDepositRecord(uint256 depositID) public returns (DepositRecord memory) {
        return _depositRecords[depositID];
    }

    function deposit(uint2546 depositID, bytes memory data) public _onlyBridge {
        address originChainTokenAddress;
        uint256 destinationChainID;
        address destinationChainHandlerAddress;
        address destinationRecipientAddress;
        address depositer;
        uint256 amount;

        assembly {
            originChainTokenAddress        := mload(add(data, 0x20))
            destinationChainID             := mload(add(data, 0x40))
            destinationChainHandlerAddress := mload(add(data, 0x60))
            destinationRecipientAddress    := mload(add(data, 0x80))
            depositer                      := mload(add(data, 0xA0))
            amount                         := mload(add(data, 0xC0))
        }

        lockERC20(originChainTokenAddress, depositer, address(this), amount);

        _depositRecords[depositID] = DepositRecord(
            originChainTokenAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            depositer,
            amount
        );
    }

    function executeDeposit(bytes memory data) public {
        address destinationChainTokenAddress;
        address destinationRecipientAddress;
        uint256 amount;

        assembly {
            destinationChainTokenAddress := mload(add(data, 0x20))
            destinationRecipientAddress  := mload(add(data, 0x40))
            amount                       := mload(add(data, 0x60))
        }

        ERC20Mintable erc20 = ERC20Mintable(destinationChainTokenAddress);
        erc20.mint(destinationRecipientAddress, amount);
    }

    function withdrawERC20(address tokenAddress, address recipient, uint amount) public _onlyBridge {
        releaseERC20(tokenAddress, address(this), recipient, amount);
    }
}

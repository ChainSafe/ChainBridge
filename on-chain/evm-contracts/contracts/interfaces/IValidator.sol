pragma solidity ^0.5.12;

interface IValidator {
    function isValidator(address validatorAddress) external returns (bool);
    function getTotalValidators() external returns (uint);
}

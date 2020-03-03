pragma solidity ^0.5.12;

interface IValidator {
    enum ValidatorActionType {Remove, Add}
    enum ThresholdType {Validator, Deposit}
    enum Vote {No, Yes}
    enum VoteStatus {Inactive, Active}

    function isValidator(address validatorAddress) external returns (bool);
    function getTotalValidators() external returns (uint);
    function createValidatorProposal(address proposedAddress, ValidatorActionType action) external;
    function voteValidatorProposal(address proposedAddress, Vote vote) external;
    function createValidatorThresholdProposal(uint proposedValue) external;
    function voteValidatorThresholdProposal(Vote vote) external;
}

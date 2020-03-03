pragma solidity ^0.5.12;

interface IValidator {
    enum Vote {No, Yes}

    // ValidatorActionType and _validatorActionTypeStrings must be kept
    // the same length and order to function properly
    enum ValidatorActionType {Remove, Add}

    // VoteStatus and _voteStatusStrings must be kept
    // the same length and order to function properly
    enum VoteStatus {Inactive, Active}

    function isValidator(address validatorAddress) external returns (bool);
    function getTotalValidators() external returns (uint);
    function createValidatorProposal(address proposedAddress, ValidatorActionType action) external;
    function voteValidatorProposal(address proposedAddress, Vote vote) external;
    function createValidatorThresholdProposal(uint proposedValue) external;
    function voteValidatorThresholdProposal(Vote vote) external;
}

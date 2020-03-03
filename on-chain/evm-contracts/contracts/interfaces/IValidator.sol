pragma solidity ^0.5.12;

interface IValidator {
    enum ValidatorActionType {Remove, Add}
    enum ThresholdType {Validator, Deposit}
    enum Vote {No, Yes}
    enum VoteStatus {Inactive, Active}

    function isValidator(address validatorAddress) external returns (bool);
    function getTotalValidators() external returns (uint);
    function createValidatorProposal(address validatorAddress, ValidatorActionType action) external;
    function voteValidatorProposal(address validatorAddress, Vote vote) external;
    function createThresholdProposal(uint value, ThresholdType thresholdType) external;
    function voteThresholdProposal(Vote vote, ThresholdType thresholdType) external;
}

pragma solidity 0.6.4;

interface IValidator {
    enum Vote {No, Yes}
    enum ValidatorActionType {Remove, Add}
    enum VoteStatus {Inactive, Active}

    function isValidator(address validatorAddress) external returns (bool);
    function getTotalValidators() external returns (uint);
    function createValidatorProposal(address proposedAddress, ValidatorActionType action) external;
    function voteValidatorProposal(address proposedAddress, Vote vote) external;
    function createValidatorThresholdProposal(uint proposedValue) external;
    function voteValidatorThresholdProposal(Vote vote) external;
}

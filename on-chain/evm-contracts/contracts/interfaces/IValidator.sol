pragma solidity ^0.5.12;

interface IValidator {
    enum ValidatorActionType {Add, Remove}
    enum ThresholdType {}
    enum Vote {No, Yes}

    function isValidator(address validatorAddress) external returns (bool);
    function getTotalValidators() external returns (uint);
    function hasVoted(uint depositID, address validatorAddress) external returns (bool);
    function createValidatorProposal(address validatorAddress, ValidatorActionType action) external;
    function voteValidatorProposal(address validatorAddress, Vote vote) external;
    function createThresholdProposal(uint value, ThresholdType type) external;
    function voteThresholdProposal(Vote vote, ThresholdType type) external;
}

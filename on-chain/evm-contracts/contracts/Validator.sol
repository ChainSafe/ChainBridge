pragma solidity ^0.5.12;

import "./interfaces/IValidator.sol";

contract Validator is IValidator {

    uint public _validatorThreshold;
    uint public _totalValidators;

    enum Vote {Yes, No}
    enum ValidatorActionType {Add, Remove}
    enum VoteStatus {Inactive, Active, Finalized, Transferred}
    enum ThresholdType {Validator, Deposit}

    struct ValidatorProposal {
        // Address of the proposed validator
        address validator;
        // validator action
        ValidatorActionType action;
        // Keeps track if a user has voted
        mapping(address => bool) votes;
        // Number of votes in favour
        uint numYes;
        // Number of votes against
        uint numNo;
        // Status of a vote
        VoteStatus status;
    }

    mapping(address => bool) public Validators;
    mapping(address => ValidatorProposal) public ValidatorProposals;

    modifier _onlyValidators() {
        require(Validators[msg.sender], "sender is not a validator");
        _;
    }

    function isValidator(address validatorAddress) public returns (bool) {
        return Validators[validatorAddress];
    }

    function getTotalValidators() public returns (uint) {
        return _totalValidators;
    }
}

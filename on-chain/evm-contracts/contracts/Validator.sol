pragma solidity ^0.5.12;

import "./interfaces/IValidator.sol";

contract Validator is IValidator {

    uint public _validatorThreshold;
    uint public _totalValidators;

    struct ValidatorProposal {
        address                  _proposedAddress;
        ValidatorActionType      _action;
        mapping(address => bool) _votes;
        uint                     _numYes;
        uint                     _numNo;
        VoteStatus               _status;
    }

    // Validator Address => whether they are a Validator
    mapping(address => bool) public _validators;
    // Validator Address => ValidatorProposal
    mapping(address => ValidatorProposal) public _validatorProposals;

    modifier _onlyValidators() {
        require(_validators[msg.sender], "sender is not a validator");
        _;
    }

    function isValidator(address validatorAddress) public returns (bool) {
        return _validators[validatorAddress];
    }

    function getTotalValidators() public returns (uint) {
        return _totalValidators;
    }

    function createValidatorProposal(address proposedAddress, ValidatorActionType action) public {
        require(uint(action) <= 1, "action out of the vote enum range");
        require(action == ValidatorActionType.Remove && _validators[proposedAddress] == true, "address is not a validator");
        require(action == ValidatorActionType.Add && _validators[proposedAddress] == false, "address is currently a validator");
        require(_validatorProposals[proposedAddress]._status == VoteStatus.Inactive, "there is already an active proposal for this address");

        _validatorProposals[proposedAddress] = ValidatorProposal({
            _proposedAddress: proposedAddress,
            _action: action,
            _numYes: 1, // Creator always votes in favour
            _numNo: 0,
            _status: VoteStatus.Active
            });

        if (_validatorThreshold == 1) {
            _validatorProposals[proposedAddress]._status = VoteStatus.Inactive;
            if (action == ValidatorActionType.Add) {
                // Add validator
                _validators[proposedAddress] = true;
                _totalValidators++;
            } else {
                // Remove validator
                _validators[proposedAddress] = false;
                _totalValidators--;
            }
        }
        // Record vote
        _validatorProposals[proposedAddress]._votes[msg.sender] = true;
    }

    // function voteValidatorProposal(address validatorAddress, Vote vote) external;

    // function createThresholdProposal(uint value, ThresholdType thresholdType) external;

    // function voteThresholdProposal(Vote vote, ThresholdType thresholdType) external;
}

pragma solidity ^0.5.12;

import "./interfaces/IValidator.sol";
import "./helpers/SafeMath.sol";

contract Validator is IValidator {
    using SafeMath for uint;

    uint public _validatorThreshold;
    uint public _totalValidators;
    ValidatorThresholdProposal private _currentValidatorThresholdProposal;
    // ValidatorActionType and _validatorActionTypeStrings must be kept
    // the same length and order to function properly
    string[] _validatorActionTypeStrings = ["remove", "add"];
    // VoteStatus and _voteStatusStrings must be kept
    // the same length and order to function properly
    string[] _voteStatusStrings = ["inactive", "active"];

    struct ValidatorProposal {
        address                  _proposedAddress;
        ValidatorActionType      _action;
        mapping(address => bool) _votes;
        uint                     _numYes;
        uint                     _numNo;
        VoteStatus               _status;
    }

    struct ValidatorThresholdProposal {
        uint                     _proposedValue;
        mapping(address => bool) _votes;
        uint                     _numYes;
        uint                     _numNo;
        VoteStatus               _status;
    }

    // Validator Address => whether they are a Validator
    mapping(address => bool) public _validators;
    // Validator Address => ValidatorProposal
    mapping(address => ValidatorProposal) public _validatorProposals;

    event ValidatorProposalCreated(address indexed proposedAddress, ValidatorActionType indexed validatorActionType);
    event ValidatorProposalVote(address indexed proposedAddress, Vote vote);
    event ValidatorAdded(address indexed validatorAddress);
    event ValidatorRemoved(address indexed validatorAddress);
    event ValidatorThresholdProposalCreated(uint indexed proposedValue);
    event ValidatorThresholdProposalVote(Vote vote);
    event ValidatorThresholdChanged(uint indexed newThreshold);

    modifier _onlyValidators() {
        require(_validators[msg.sender], "sender is not a validator");
        _;
    }

    constructor (address[] memory initialValidators, uint initialValidatorThreshold) public {
        for (uint i; i < initialValidators.length; i++) {
            _addValidator(initialValidators[i]);
        }

        _validatorThreshold = initialValidatorThreshold;
    }

    function isValidator(address validatorAddress) public returns (bool) {
        return _validators[validatorAddress];
    }

    function getValidatorThreshold() public view returns (uint) {
        return _validatorThreshold;
    }

    function getTotalValidators() public returns (uint) {
        return _totalValidators;
    }

    function getCurrentValidatorThresholdProposal() public view returns (
        uint, uint, uint, string memory) {
        return (
        _currentValidatorThresholdProposal._proposedValue,
        _currentValidatorThresholdProposal._numYes,
        _currentValidatorThresholdProposal._numNo,
        _voteStatusStrings[uint(_currentValidatorThresholdProposal._status)]);
    }

    function getValidatorProposal(address proposedAddress) public view returns (
        address, string memory, uint, uint, string memory) {
        ValidatorProposal memory validatorProposal = _validatorProposals[proposedAddress];
        return (
        validatorProposal._proposedAddress,
        _validatorActionTypeStrings[uint(validatorProposal._action)],
        validatorProposal._numYes,
        validatorProposal._numNo,
        _voteStatusStrings[uint(validatorProposal._status)]);
    }

    function createValidatorProposal(address proposedAddress, ValidatorActionType action) public _onlyValidators {
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

        if (_validatorThreshold <= 1) {
            _validatorProposals[proposedAddress]._status = VoteStatus.Inactive;
            if (action == ValidatorActionType.Add) {
                _addValidator(proposedAddress);
            } else {
                _removeValidator(proposedAddress);
            }
        }
        // Record vote
        _validatorProposals[proposedAddress]._votes[msg.sender] = true;
        emit ValidatorProposalCreated(proposedAddress, action);
    }

    function voteValidatorProposal(address proposedAddress, Vote vote) public _onlyValidators {
        require(_validatorProposals[proposedAddress]._status == VoteStatus.Active, "there is no active proposal for this address");
        require(!_validatorProposals[proposedAddress]._votes[msg.sender], "validator has already voted");
        require(uint(vote) <= 1, "vote out of the vote enum range");

        // Cast vote
        if (vote == Vote.Yes) {
            _validatorProposals[proposedAddress]._numYes++;
        } else {
            _validatorProposals[proposedAddress]._numNo++;
        }

        // Record vote
        _validatorProposals[proposedAddress]._votes[msg.sender] = true;
        emit ValidatorProposalVote(proposedAddress, vote);

        // Todo: Edge case if validator threshold changes?
        // Todo: For a proposal to pass does the number of yes votes just need to be higher than the threshold, or does it also have to be greater than the number of no votes?
        if (_validatorProposals[proposedAddress]._numYes >= _validatorThreshold) {
            if (_validatorProposals[proposedAddress]._action == ValidatorActionType.Add) {
                _addValidator(proposedAddress);
            } else {
                _removeValidator(proposedAddress);
            }

            _validatorProposals[proposedAddress]._status = VoteStatus.Inactive;
        } else if (_totalValidators.sub(_validatorProposals[proposedAddress]._numNo) < _validatorThreshold) {
            _validatorProposals[proposedAddress]._status = VoteStatus.Inactive;
        }
    }

    function createValidatorThresholdProposal(uint proposedValue) public _onlyValidators {
        require(_currentValidatorThresholdProposal._status == VoteStatus.Inactive, "a proposal is currently active");
        require(proposedValue <= _totalValidators, "proposed value cannot be greater than the total number of validators");

        _currentValidatorThresholdProposal = ValidatorThresholdProposal({
            _proposedValue: proposedValue,
            _numYes: 1, // Creator always votes in favour
            _numNo: 0,
            _status: VoteStatus.Active
            });

        if (_validatorThreshold <= 1) {
            _validatorThreshold = _currentValidatorThresholdProposal._proposedValue;
            _currentValidatorThresholdProposal._status = VoteStatus.Inactive;
            emit ValidatorThresholdChanged(proposedValue);
        }
        // Record vote
        _currentValidatorThresholdProposal._votes[msg.sender] = true;
        emit ValidatorThresholdProposalCreated(proposedValue);
    }

    function voteValidatorThresholdProposal(Vote vote) public _onlyValidators {
        require(_currentValidatorThresholdProposal._status == VoteStatus.Active, "no proposal is currently active");
        require(!_currentValidatorThresholdProposal._votes[msg.sender], "validator has already voted");
        require(uint(vote) <= 1, "vote out of the vote enum range");

        // Cast vote
        if (vote == Vote.Yes) {
            _currentValidatorThresholdProposal._numYes++;
        } else {
            _currentValidatorThresholdProposal._numNo++;
        }

        _currentValidatorThresholdProposal._votes[msg.sender] = true;
        emit ValidatorThresholdProposalVote(vote);

        // Todo: Edge case if validator threshold changes?
        // Todo: For a proposal to pass does the number of yes votes just need to be higher than the threshold, or does it also have to be greater than the number of no votes?
        if (_currentValidatorThresholdProposal._numYes >= _validatorThreshold) {
            _validatorThreshold = _currentValidatorThresholdProposal._proposedValue;
            _currentValidatorThresholdProposal._status = VoteStatus.Inactive;
            emit ValidatorThresholdChanged(_currentValidatorThresholdProposal._proposedValue);
        } else if (_totalValidators.sub(_currentValidatorThresholdProposal._numNo) < _validatorThreshold) {
            _currentValidatorThresholdProposal._status = VoteStatus.Inactive;
        }
    }

    function _addValidator(address addr) internal {
        _validators[addr] = true;
        _totalValidators++;
        emit ValidatorAdded(addr);
    }

    function _removeValidator(address addr) internal {
        _validators[addr] = false;
        _totalValidators--;
        emit ValidatorRemoved(addr);
    }
}

pragma solidity 0.6.4;

import "./interfaces/IRelayer.sol";
import "./helpers/SafeMath.sol";

contract Relayer is IRelayer {
    using SafeMath for uint;

    uint public _relayerThreshold;
    uint public _totalRelayers;
    RelayerThresholdProposal private _currentRelayerThresholdProposal;
    // RelayerActionType and _relayerActionTypeStrings must be kept
    // the same length and order to function properly
    string[] _relayerActionTypeStrings = ["remove", "add"];
    // VoteStatus and _voteStatusStrings must be kept
    // the same length and order to function properly
    string[] _voteStatusStrings = ["inactive", "active"];

    struct RelayerProposal {
        address                  _proposedAddress;
        RelayerActionType      _action;
        mapping(address => bool) _votes;
        uint                     _numYes;
        uint                     _numNo;
        VoteStatus               _status;
    }

    struct RelayerThresholdProposal {
        uint                     _proposedValue;
        mapping(address => bool) _votes;
        uint                     _numYes;
        uint                     _numNo;
        VoteStatus               _status;
    }

    // Relayer Address => whether they are a Relayer
    mapping(address => bool) public _relayers;
    // Relayer Address => RelayerProposal
    mapping(address => RelayerProposal) public _relayerProposals;

    event RelayerProposalCreated(address indexed proposedAddress, RelayerActionType indexed relayerActionType);
    event RelayerProposalVote(address indexed proposedAddress, Vote vote);
    event RelayerAdded(address indexed relayerAddress);
    event RelayerRemoved(address indexed relayerAddress);
    event RelayerThresholdProposalCreated(uint indexed proposedValue);
    event RelayerThresholdProposalVote(Vote vote);
    event RelayerThresholdChanged(uint indexed newThreshold);

    modifier _onlyRelayers() {
        require(_relayers[msg.sender], "sender is not a relayer");
        _;
    }

    constructor (address[] memory initialRelayers, uint initialRelayerThreshold) public {
        for (uint i; i < initialRelayers.length; i++) {
            _addRelayer(initialRelayers[i]);
        }

        _relayerThreshold = initialRelayerThreshold;
    }

    function isRelayer(address relayerAddress) public override returns (bool) {
        return _relayers[relayerAddress];
    }

    function getRelayerThreshold() public view returns (uint) {
        return _relayerThreshold;
    }

    function getTotalRelayers() public override returns (uint) {
        return _totalRelayers;
    }

    function getCurrentRelayerThresholdProposal() public view returns (
        uint, uint, uint, string memory) {
        return (
        _currentRelayerThresholdProposal._proposedValue,
        _currentRelayerThresholdProposal._numYes,
        _currentRelayerThresholdProposal._numNo,
        _voteStatusStrings[uint(_currentRelayerThresholdProposal._status)]);
    }

    function getRelayerProposal(address proposedAddress) public view returns (
        address, string memory, uint, uint, string memory) {
        RelayerProposal memory relayerProposal = _relayerProposals[proposedAddress];
        return (
        relayerProposal._proposedAddress,
        _relayerActionTypeStrings[uint(relayerProposal._action)],
        relayerProposal._numYes,
        relayerProposal._numNo,
        _voteStatusStrings[uint(relayerProposal._status)]);
    }

    function createRelayerProposal(address proposedAddress, RelayerActionType action) public override _onlyRelayers {
        require(uint(action) <= 1, "action out of the vote enum range");
        require(action == RelayerActionType.Remove && _relayers[proposedAddress] == true, "address is not a relayer");
        require(action == RelayerActionType.Add && _relayers[proposedAddress] == false, "address is currently a relayer");
        require(_relayerProposals[proposedAddress]._status == VoteStatus.Inactive, "there is already an active proposal for this address");

        _relayerProposals[proposedAddress] = RelayerProposal({
            _proposedAddress: proposedAddress,
            _action: action,
            _numYes: 1, // Creator always votes in favour
            _numNo: 0,
            _status: VoteStatus.Active
            });

        if (_relayerThreshold <= 1) {
            _relayerProposals[proposedAddress]._status = VoteStatus.Inactive;
            if (action == RelayerActionType.Add) {
                _addRelayer(proposedAddress);
            } else {
                _removeRelayer(proposedAddress);
            }
        }
        // Record vote
        _relayerProposals[proposedAddress]._votes[msg.sender] = true;
        emit RelayerProposalCreated(proposedAddress, action);
    }

    function voteRelayerProposal(address proposedAddress, Vote vote) public override _onlyRelayers {
        require(_relayerProposals[proposedAddress]._status == VoteStatus.Active, "there is no active proposal for this address");
        require(!_relayerProposals[proposedAddress]._votes[msg.sender], "relayer has already voted");
        require(uint(vote) <= 1, "vote out of the vote enum range");

        // Cast vote
        if (vote == Vote.Yes) {
            _relayerProposals[proposedAddress]._numYes++;
        } else {
            _relayerProposals[proposedAddress]._numNo++;
        }

        // Record vote
        _relayerProposals[proposedAddress]._votes[msg.sender] = true;
        emit RelayerProposalVote(proposedAddress, vote);

        // Todo: Edge case if relayer threshold changes?
        // Todo: For a proposal to pass does the number of yes votes just need to be higher than the threshold, or does it also have to be greater than the number of no votes?
        if (_relayerProposals[proposedAddress]._numYes >= _relayerThreshold) {
            if (_relayerProposals[proposedAddress]._action == RelayerActionType.Add) {
                _addRelayer(proposedAddress);
            } else {
                _removeRelayer(proposedAddress);
            }

            _relayerProposals[proposedAddress]._status = VoteStatus.Inactive;
        } else if (_totalRelayers.sub(_relayerProposals[proposedAddress]._numNo) < _relayerThreshold) {
            _relayerProposals[proposedAddress]._status = VoteStatus.Inactive;
        }
    }

    function createRelayerThresholdProposal(uint proposedValue) public override _onlyRelayers {
        require(_currentRelayerThresholdProposal._status == VoteStatus.Inactive, "a proposal is currently active");
        require(proposedValue <= _totalRelayers, "proposed value cannot be greater than the total number of relayers");

        _currentRelayerThresholdProposal = RelayerThresholdProposal({
            _proposedValue: proposedValue,
            _numYes: 1, // Creator always votes in favour
            _numNo: 0,
            _status: VoteStatus.Active
            });

        if (_relayerThreshold <= 1) {
            _relayerThreshold = _currentRelayerThresholdProposal._proposedValue;
            _currentRelayerThresholdProposal._status = VoteStatus.Inactive;
            emit RelayerThresholdChanged(proposedValue);
        }
        // Record vote
        _currentRelayerThresholdProposal._votes[msg.sender] = true;
        emit RelayerThresholdProposalCreated(proposedValue);
    }

    function voteRelayerThresholdProposal(Vote vote) public override _onlyRelayers {
        require(_currentRelayerThresholdProposal._status == VoteStatus.Active, "no proposal is currently active");
        require(!_currentRelayerThresholdProposal._votes[msg.sender], "relayer has already voted");
        require(uint(vote) <= 1, "vote out of the vote enum range");

        // Cast vote
        if (vote == Vote.Yes) {
            _currentRelayerThresholdProposal._numYes++;
        } else {
            _currentRelayerThresholdProposal._numNo++;
        }

        _currentRelayerThresholdProposal._votes[msg.sender] = true;
        emit RelayerThresholdProposalVote(vote);

        // Todo: Edge case if relayer threshold changes?
        // Todo: For a proposal to pass does the number of yes votes just need to be higher than the threshold, or does it also have to be greater than the number of no votes?
        if (_currentRelayerThresholdProposal._numYes >= _relayerThreshold) {
            _relayerThreshold = _currentRelayerThresholdProposal._proposedValue;
            _currentRelayerThresholdProposal._status = VoteStatus.Inactive;
            emit RelayerThresholdChanged(_currentRelayerThresholdProposal._proposedValue);
        } else if (_totalRelayers.sub(_currentRelayerThresholdProposal._numNo) < _relayerThreshold) {
            _currentRelayerThresholdProposal._status = VoteStatus.Inactive;
        }
    }

    function _addRelayer(address addr) internal {
        _relayers[addr] = true;
        _totalRelayers++;
        emit RelayerAdded(addr);
    }

    function _removeRelayer(address addr) internal {
        _relayers[addr] = false;
        _totalRelayers--;
        emit RelayerRemoved(addr);
    }
}

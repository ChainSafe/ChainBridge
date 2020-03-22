pragma solidity 0.6.1;
pragma experimental ABIEncoderV2;

import "./helpers/SafeMath.sol";
import "./interfaces/IValidator.sol";
import "./interfaces/IERC20Handler.sol";
import "./interfaces/IERC721Handler.sol";
import "./interfaces/IDepositHandler.sol";

contract Bridge {
    using SafeMath for uint;

    IValidator                 public _validatorContract;
    uint256                    public _validatorThreshold;
    ValidatorThresholdProposal public _currentValidatorThresholdProposal;

    enum Vote {No, Yes}
    enum ValidatorThresholdProposalStatus {Inactive, Active}
    enum DepositProposalStatus {Inactive, Active, Denied, Passed, Transferred}

    struct DepositRecord {
        bytes   _data;
    }
    struct DepositProposal {
        bytes32                  _dataHash;
        mapping(address => bool) _votes;
        uint8                    _numYes;
        uint8                    _numNo;
        DepositProposalStatus    _status;
    }
    struct ValidatorThresholdProposal {
        uint256                          _proposedValue;
        mapping(address => bool)         _votes;
        uint8                            _numYes;
        uint8                            _numNo;
        ValidatorThresholdProposalStatus _status;
    }

    // handler address => number of deposits
    mapping(address => uint256) public _depositCounts;
    // handler address => depositID => DepositRecord
    mapping(address => mapping(uint256 => DepositRecord)) public _depositRecords;
    // originChainID => handler address => depositID => depositProposal
    mapping(uint256 => mapping(address => mapping(uint256 => DepositProposal))) public _depositProposals;

    event ValidatorThresholdProposalCreated(uint indexed proposedValue);
    event ValidatorThresholdProposalVote(Vote vote);
    event ValidatorThresholdChanged(uint indexed newThreshold);
    event Deposit(address indexed handlerAddress, uint256 indexed depositID);
    event DepositProposalCreated(uint indexed originChainID, address indexed originChainHandlerAddress, uint indexed depositID, bytes32 dataHash);
    event DepositProposalVote(uint indexed originChainID, uint indexed depositID, Vote indexed vote, DepositProposalStatus status);
    event DepositProposalFinalized(uint indexed originChainID, uint indexed depositID);

    modifier _onlyValidators() {
        IValidator validatorContract = IValidator(_validatorContract);
        require(validatorContract.isValidator(msg.sender), "sender is not a validator");
        _;
    }

    constructor (address validatorContract, uint initialValidatorThreshold) public {
        _validatorContract = IValidator(validatorContract);
        _validatorThreshold = initialValidatorThreshold;
    }

    function getValidatorThreshold() public view returns (uint256) {
        return _validatorThreshold;
    }

    function getCurrentValidatorThresholdProposal() public view returns (ValidatorThresholdProposal memory) {
        return _currentValidatorThresholdProposal;
    }

    function getDepositCount(uint originChainID) public view returns (uint256) {
        return _depositCounts[originChainID];
    }

    function getDepositRecord(address handlerAddress, uint256 depositID) public view returns (DepositRecord memory) {
        return _depositRecords[handlerAddress][depositID];
    }

    function getDepositProposal(
        uint256 originChainID, address originChainHandlerAddress, uint256 depositID) public view returns (DepositProposal memory) {
        return _depositProposals[handlerAddress][originChainHandlerAddress][depositID];
    }

    function hasVoted(uint originChainID, uint depositID, address validatorAddress) public view returns (bool) {
        return _depositProposals[originChainID][depositID]._votes[validatorAddress];
    }

    function deposit(
        address      handlerAddress,
        bytes memory data
    ) public {
        uint depositID = ++_depositCounts[handlerAddress];

        _depositRecords[handlerAddress][depositID] = DepositRecord(
            data
        );

        IDepositHandler depositHandler = IDepositHandler(handlerAddress);
        depositHandler.deposit(depositID, data);

        emit Deposit(handlerAddress, depositID);
    }

    function createDepositProposal(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositID,
        bytes32 dataHash
    ) public _onlyValidators {
        require(
            _depositProposals[originChainID][originChainHandlerAddress][depositID]._status == DepositProposalStatus.Inactive ||
            _depositProposals[originChainID][originChainHandlerAddress][depositID]._status == DepositProposalStatus.Denied,
            "proposal is either currently active or has already been passed/transferred"
        );

        // If _depositThreshold is set to 1, then auto finalize
        if (_validatorThreshold <= 1) {
            _depositProposals[originChainID][originChainHandlerAddress][depositID] = DepositProposal({
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: DepositProposalStatus.Passed
            });
        } else {
            _depositProposals[originChainID][originChainHandlerAddress][depositID] = DepositProposal({
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: DepositProposalStatus.Active
            });
        }

        // Creator always votes in favour
        _depositProposals[originChainID][originChainHandlerAddress][depositID]._votes[msg.sender] = true;

        emit DepositProposalCreated(originChainID, originChainHandlerAddress, depositID, dataHash);
    }

    function voteDepositProposal(
        uint256 originChainID,
        uint256 depositID,
        Vote vote
    ) public _onlyValidators {
        DepositProposal storage depositProposal = _depositProposals[originChainID][depositID];

        require(depositProposal._status != DepositProposalStatus.Inactive, "proposal is not active");
        require(depositProposal._status == DepositProposalStatus.Active, "proposal has been finalized");
        require(!depositProposal._votes[msg.sender], "validator has already voted");
        require(uint(vote) <= 1, "invalid vote");

        if (vote == Vote.Yes) {
            depositProposal._numYes++;
        } else {
            depositProposal._numNo++;
        }

        depositProposal._votes[msg.sender] = true;

        // Todo: Edge case if validator threshold changes?
        if (depositProposal._numYes >= _validatorThreshold) {
            depositProposal._status = DepositProposalStatus.Passed;
            emit DepositProposalFinalized(originChainID, depositID);
        } else if (_validatorContract.getTotalValidators().sub(depositProposal._numNo) < _validatorThreshold) {
            depositProposal._status = DepositProposalStatus.Denied;
            emit DepositProposalFinalized(originChainID, depositID);
        }

        emit DepositProposalVote(originChainID, depositID, vote, depositProposal._status);
    }

    function executeDepositProposal(uint originChainID, uint depositID, address destinationChainHandlerAddress, bytes memory data) public {
        DepositProposal storage depositProposal = _depositProposals[originChainID][depositID];

        require(depositProposal._status != DepositProposalStatus.Inactive, "proposal is not active");
        require(depositProposal._status == DepositProposalStatus.Passed, "proposal was not passed or has already been transferred");
        require(keccak256(data) == depositProposal._dataHash, "provided data does not match proposal's data hash");

        IDepositHandler depositHandler = IDepositHandler(destinationChainHandlerAddress);
        depositHandler.executeDeposit(data);

        depositProposal._status = DepositProposalStatus.Transferred;
    }

    function createValidatorThresholdProposal(uint proposedValue) public _onlyValidators {
        require(_currentValidatorThresholdProposal._status == ValidatorThresholdProposalStatus.Inactive, "a proposal is currently active");
        require(proposedValue <= _validatorContract.getTotalValidators(), "proposed value cannot be greater than the total number of validators");

        _currentValidatorThresholdProposal = ValidatorThresholdProposal({
            _proposedValue: proposedValue,
            _numYes: 1, // Creator always votes in favour
            _numNo: 0,
            _status: ValidatorThresholdProposalStatus.Active
            });

        if (_validatorThreshold <= 1) {
            _validatorThreshold = _currentValidatorThresholdProposal._proposedValue;
            _currentValidatorThresholdProposal._status = ValidatorThresholdProposalStatus.Inactive;
            emit ValidatorThresholdChanged(proposedValue);
        }
        // Record vote
        _currentValidatorThresholdProposal._votes[msg.sender] = true;
        emit ValidatorThresholdProposalCreated(proposedValue);
    }

    function voteValidatorThresholdProposal(Vote vote) public _onlyValidators {
        require(_currentValidatorThresholdProposal._status == ValidatorThresholdProposalStatus.Active, "no proposal is currently active");
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
            _currentValidatorThresholdProposal._status = ValidatorThresholdProposalStatus.Inactive;
            emit ValidatorThresholdChanged(_currentValidatorThresholdProposal._proposedValue);
        } else if (_validatorContract.getTotalValidators().sub(_currentValidatorThresholdProposal._numNo) < _validatorThreshold) {
            _currentValidatorThresholdProposal._status = ValidatorThresholdProposalStatus.Inactive;
        }
    }
}
pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "./helpers/SafeMath.sol";
import "./interfaces/IRelayer.sol";
import "./interfaces/IDepositHandler.sol";

contract Bridge {
    using SafeMath for uint;

    uint256 constant public CHAIN_ID = 0;

    IRelayer                 public _relayerContract;
    uint256                    public _relayerThreshold;
    RelayerThresholdProposal public _currentRelayerThresholdProposal;

    enum Vote {No, Yes}
    enum RelayerThresholdProposalStatus {Inactive, Active}
    enum DepositProposalStatus {Inactive, Active, Denied, Passed, Transferred}

    struct CentrifugeAssetDepositRecord {
        address   _originChainTokenAddress;
        address   _originChainHandlerAddress;
        uint      _destinationChainID;
        address   _destinationChainHandlerAddress;
        address   _destinationRecipientAddress;
        bytes32   _metaDataHash;
    }

    struct DepositProposal {
        bytes32                  _dataHash;
        mapping(address => bool) _votes;
        uint256                  _numYes;
        uint256                  _numNo;
        DepositProposalStatus    _status;
    }
    struct ReturnDepositProposal {
        bytes32                  _dataHash;
        uint256                  _numYes;
        uint256                  _numNo;
        DepositProposalStatus    _status;
    }

    struct RelayerThresholdProposal {
        uint256                          _proposedValue;
        mapping(address => bool)         _votes;
        uint256                          _numYes;
        uint256                          _numNo;
        RelayerThresholdProposalStatus _status;
    }

    // originChainHandlerAddress address => number of deposits
    mapping(address => uint256) public _depositCounts;
    // originChainHandlerAddress => depositNonce => bytes
    mapping(address => mapping(uint256 => bytes)) public _depositRecords;
    // originChainID => originChainHandlerAddress => depositNonce => depositProposal
    mapping(uint256 => mapping(address => mapping(uint256 => DepositProposal))) public _depositProposals;

    event RelayerThresholdProposalCreated(uint indexed proposedValue);
    event RelayerThresholdProposalVote(Vote vote);
    event RelayerThresholdChanged(uint indexed newThreshold);
    event Deposit(
        uint256 indexed originChainID,
        address indexed originChainHandlerAddress,
        uint256 indexed depositNonce
    );
    event DepositProposalCreated(
        uint256 indexed originChainID,
        address indexed originChainHandlerAddress,
        uint256 indexed depositNonce,
        bytes32         dataHash
    );
    event DepositProposalVote(
        uint256 indexed       originChainID,
        address indexed       originChainHandlerAddress,
        uint256 indexed       depositNonce,
        Vote                  vote,
        DepositProposalStatus status
    );
    event DepositProposalFinalized(
        uint256 indexed originChainID,
        address indexed originChainHandlerAddress,
        uint256 indexed depositNonce
    );
    event DepositProposalExecuted(
        uint256 indexed originChainID,
        address indexed originChainHandlerAddress,
        uint256 indexed depositNonce
    );

    modifier _onlyRelayers() {
        IRelayer relayerContract = IRelayer(_relayerContract);
        require(relayerContract.isRelayer(msg.sender), "sender must be a relayer");
        _;
    }

    constructor (address relayerContract, uint initialRelayerThreshold) public {
        _relayerContract = IRelayer(relayerContract);
        _relayerThreshold = initialRelayerThreshold;
    }

    function getDepositProposal(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositNonce
    ) public view returns (ReturnDepositProposal memory) {
        DepositProposal memory depositProposal = _depositProposals[originChainID][originChainHandlerAddress][depositNonce];
        return ReturnDepositProposal(
            depositProposal._dataHash,
            depositProposal._numYes,
            depositProposal._numNo,
            depositProposal._status
        );
    }

    function hasVoted(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositNonce,
        address relayerAddress
    ) public view returns (bool) {
        return _depositProposals[originChainID][originChainHandlerAddress][depositNonce]._votes[relayerAddress];
    }

    function deposit(
        address      originChainHandlerAddress,
        bytes memory data
    ) public {
        uint depositNonce = ++_depositCounts[originChainHandlerAddress];
        _depositRecords[originChainHandlerAddress][depositNonce] = data;

        IDepositHandler depositHandler = IDepositHandler(originChainHandlerAddress);
        depositHandler.deposit(depositNonce, data);

        emit Deposit(CHAIN_ID, originChainHandlerAddress, depositNonce);
    }

    function createDepositProposal(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositNonce,
        bytes32 dataHash
    ) public _onlyRelayers {
        require(
            _depositProposals[originChainID][originChainHandlerAddress][depositNonce]._status == DepositProposalStatus.Inactive ||
            _depositProposals[originChainID][originChainHandlerAddress][depositNonce]._status == DepositProposalStatus.Denied,
            "proposal is either currently active or has already been passed/transferred"
        );

        // If _depositThreshold is set to 1, then auto finalize
        if (_relayerThreshold <= 1) {
            _depositProposals[originChainID][originChainHandlerAddress][depositNonce] = DepositProposal({
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: DepositProposalStatus.Passed
            });
        } else {
            _depositProposals[originChainID][originChainHandlerAddress][depositNonce] = DepositProposal({
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: DepositProposalStatus.Active
            });
        }

        // Creator always votes in favour
        _depositProposals[originChainID][originChainHandlerAddress][depositNonce]._votes[msg.sender] = true;

        emit DepositProposalCreated(originChainID, originChainHandlerAddress, depositNonce, dataHash);
    }

    function voteDepositProposal(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositNonce,
        Vote    vote
    ) public _onlyRelayers {
        DepositProposal storage depositProposal = _depositProposals[originChainID][originChainHandlerAddress][depositNonce];

        require(depositProposal._status != DepositProposalStatus.Inactive, "proposal is not active");
        require(depositProposal._status == DepositProposalStatus.Active, "proposal has been finalized");
        require(!depositProposal._votes[msg.sender], "relayer has already voted");
        require(uint(vote) <= 1, "invalid vote");

        if (vote == Vote.Yes) {
            depositProposal._numYes++;
        } else {
            depositProposal._numNo++;
        }

        depositProposal._votes[msg.sender] = true;

        // Todo: Edge case if relayer threshold changes?
        if (depositProposal._numYes >= _relayerThreshold) {
            depositProposal._status = DepositProposalStatus.Passed;
            emit DepositProposalFinalized(originChainID, originChainHandlerAddress, depositNonce);
        } else if (_relayerContract.getTotalRelayers().sub(depositProposal._numNo) < _relayerThreshold) {
            depositProposal._status = DepositProposalStatus.Denied;
            emit DepositProposalFinalized(originChainID, originChainHandlerAddress, depositNonce);
        }

        emit DepositProposalVote(originChainID, originChainHandlerAddress, depositNonce, vote, depositProposal._status);
    }

    function executeDepositProposal(
        uint256      originChainID,
        address      originChainHandlerAddress,
        uint256      depositNonce,
        address      destinationChainHandlerAddress,
        bytes memory data
    ) public {
        DepositProposal storage depositProposal = _depositProposals[originChainID][originChainHandlerAddress][depositNonce];

        require(depositProposal._status != DepositProposalStatus.Inactive, "proposal is not active");
        require(depositProposal._status == DepositProposalStatus.Passed, "proposal was not passed or has already been transferred");
        require(keccak256(data) == depositProposal._dataHash, "provided data does not match proposal's data hash");

        IDepositHandler depositHandler = IDepositHandler(destinationChainHandlerAddress);
        depositHandler.executeDeposit(data);

        depositProposal._status = DepositProposalStatus.Transferred;
        emit DepositProposalExecuted(originChainID, originChainHandlerAddress, depositNonce);
    }

    function createRelayerThresholdProposal(uint proposedValue) public _onlyRelayers {
        require(_currentRelayerThresholdProposal._status == RelayerThresholdProposalStatus.Inactive, "a proposal is currently active");
        require(proposedValue <= _relayerContract.getTotalRelayers(), "proposed value cannot be greater than the total number of relayers");

        _currentRelayerThresholdProposal = RelayerThresholdProposal({
            _proposedValue: proposedValue,
            _numYes: 1, // Creator always votes in favour
            _numNo: 0,
            _status: RelayerThresholdProposalStatus.Active
            });

        if (_relayerThreshold <= 1) {
            _relayerThreshold = _currentRelayerThresholdProposal._proposedValue;
            _currentRelayerThresholdProposal._status = RelayerThresholdProposalStatus.Inactive;
            emit RelayerThresholdChanged(proposedValue);
        }
        // Record vote
        _currentRelayerThresholdProposal._votes[msg.sender] = true;
        emit RelayerThresholdProposalCreated(proposedValue);
    }

    function voteRelayerThresholdProposal(Vote vote) public _onlyRelayers {
        require(_currentRelayerThresholdProposal._status == RelayerThresholdProposalStatus.Active, "no proposal is currently active");
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
            _currentRelayerThresholdProposal._status = RelayerThresholdProposalStatus.Inactive;
            emit RelayerThresholdChanged(_currentRelayerThresholdProposal._proposedValue);
        } else if (_relayerContract.getTotalRelayers().sub(_currentRelayerThresholdProposal._numNo) < _relayerThreshold) {
            _currentRelayerThresholdProposal._status = RelayerThresholdProposalStatus.Inactive;
        }
    }
}
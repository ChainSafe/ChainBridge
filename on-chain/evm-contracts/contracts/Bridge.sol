pragma solidity 0.6.4;
<<<<<<< HEAD
pragma experimental ABIEncoderV2;

import "./helpers/SafeMath.sol";
import "./interfaces/IValidator.sol";
=======

import "./helpers/SafeMath.sol";
import "./interfaces/IRelayer.sol";
import "./interfaces/IERC20Handler.sol";
import "./interfaces/IERC721Handler.sol";
>>>>>>> master
import "./interfaces/IDepositHandler.sol";

contract Bridge {
    using SafeMath for uint;

<<<<<<< HEAD
    uint256 constant public CHAIN_ID = 0;

    IValidator                 public _validatorContract;
    uint256                    public _validatorThreshold;
    ValidatorThresholdProposal public _currentValidatorThresholdProposal;

    enum Vote {No, Yes}
    enum ValidatorThresholdProposalStatus {Inactive, Active}
=======
    IRelayer public _relayerContract;
    uint public _relayerThreshold;
    RelayerThresholdProposal public _currentRelayerThresholdProposal;

    enum Vote {No, Yes}

    // RelayerThresholdProposalStatus and _relayerThresholdProposalStatusStrings must be kept
    // the same length and order to function properly
    enum RelayerThresholdProposalStatus {Inactive, Active}
    string[] _relayerThresholdProposalStatusStrings = ["inactive", "active"];

    // DepositProposalStatus and _depositProposalStatusStrings must be kept
    // the same length and order to function properly
>>>>>>> master
    enum DepositProposalStatus {Inactive, Active, Denied, Passed, Transferred}

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
<<<<<<< HEAD
    struct ValidatorThresholdProposal {
        uint256                          _proposedValue;
        mapping(address => bool)         _votes;
        uint256                          _numYes;
        uint256                          _numNo;
        ValidatorThresholdProposalStatus _status;
    }

    // originChainHandlerAddress address => number of deposits
    mapping(address => uint256) public _depositCounts;
    // originChainHandlerAddress => depositID => bytes
    mapping(address => mapping(uint256 => bytes)) public _depositRecords;
    // originChainID => originChainHandlerAddress => depositID => depositProposal
    mapping(uint256 => mapping(address => mapping(uint256 => DepositProposal))) public _depositProposals;

    event ValidatorThresholdProposalCreated(uint indexed proposedValue);
    event ValidatorThresholdProposalVote(Vote vote);
    event ValidatorThresholdChanged(uint indexed newThreshold);
    event Deposit(
        uint256 indexed originChainID,
        address indexed originChainHandlerAddress,
        uint256 indexed depositID
    );
    event DepositProposalCreated(
        uint256 indexed originChainID,
        address indexed originChainHandlerAddress,
        uint256 indexed depositID,
        bytes32         dataHash
    );
    event DepositProposalVote(
        uint256 indexed       originChainID,
        address indexed       originChainHandlerAddress,
        uint256 indexed       depositID,
        Vote                  vote,
        DepositProposalStatus status
    );
    event DepositProposalFinalized(
        uint256 indexed originChainID,
        address indexed originChainHandlerAddress,
        uint256 indexed depositID
    );
    event DepositProposalExecuted(
        uint256 indexed originChainID,
        address indexed originChainHandlerAddress,
        uint256 indexed depositID
    );

    modifier _onlyValidators() {
        IValidator validatorContract = IValidator(_validatorContract);
        require(validatorContract.isValidator(msg.sender), "sender is not a validator");
=======

    struct RelayerThresholdProposal {
        uint                             _proposedValue;
        mapping(address => bool)         _votes;
        uint                             _numYes;
        uint                             _numNo;
        RelayerThresholdProposalStatus _status;
    }

    // chainID => number of deposits
    mapping(uint => uint) public _depositCounts;
    // chainID => depositID => GenericDepositRecord
    mapping(uint => mapping(uint => GenericDepositRecord)) public _genericDepositRecords;
    // chainID => depositID => ERC20DepositRecord
    mapping(uint => mapping(uint => ERC20DepositRecord)) public _erc20DepositRecords;
    // chainID => depositID => ERC721DepositRecord
    mapping(uint => mapping(uint => ERC721DepositRecord)) public _erc721DepositRecords;
    // ChainId => DepositID => Proposal
    mapping(uint => mapping(uint => DepositProposal)) public _depositProposals;

    event GenericDeposited(uint indexed depositID);
    event ERC20Deposited(uint indexed depositID);
    event ERC721Deposited(uint indexed depositID);
    event DepositProposalCreated(uint indexed originChainID, uint indexed depositID, bytes32 indexed dataHash);
    event DepositProposalVote(uint indexed originChainID, uint indexed depositID, Vote indexed vote, DepositProposalStatus status);
    event DepositProposalFinalized(uint indexed originChainID, uint indexed depositID);
    event RelayerThresholdProposalCreated(uint indexed proposedValue);
    event RelayerThresholdProposalVote(Vote vote);
    event RelayerThresholdChanged(uint indexed newThreshold);

    modifier _onlyRelayers() {
        IRelayer relayerContract = IRelayer(_relayerContract);
        require(relayerContract.isRelayer(msg.sender));
>>>>>>> master
        _;
    }

    constructor (address relayerContract, uint initialRelayerThreshold) public {
        _relayerContract = IRelayer(relayerContract);
        _relayerThreshold = initialRelayerThreshold;
    }

<<<<<<< HEAD
    function getDepositProposal(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositID
    ) public view returns (ReturnDepositProposal memory) {
        DepositProposal memory depositProposal = _depositProposals[originChainID][originChainHandlerAddress][depositID];
        return ReturnDepositProposal(
            depositProposal._dataHash,
            depositProposal._numYes,
            depositProposal._numNo,
            depositProposal._status
=======
    function getRelayerThreshold() public view returns (uint) {
        return _relayerThreshold;
    }

    function getDepositCount(uint originChainID) public view returns (uint) {
        return _depositCounts[originChainID];
    }

    function getGenericDepositRecord(uint originChainID, uint depositID) public view returns (
        address, address, uint, address, address, bytes memory) {
        GenericDepositRecord memory genericDepositRecord = _genericDepositRecords[originChainID][depositID];
        return (
            genericDepositRecord._originChainTokenAddress,
            genericDepositRecord._originChainHandlerAddress,
            genericDepositRecord._destinationChainID,
            genericDepositRecord._destinationChainHandlerAddress,
            genericDepositRecord._destinationRecipientAddress,
            genericDepositRecord._data);
    }

    function getERC20DepositRecord(uint originChainID, uint depositID) public view returns (
        address, address, uint, address, address, uint) {
        ERC20DepositRecord memory erc20DepositRecord = _erc20DepositRecords[originChainID][depositID];
        return (
            erc20DepositRecord._originChainTokenAddress,
            erc20DepositRecord._originChainHandlerAddress,
            erc20DepositRecord._destinationChainID,
            erc20DepositRecord._destinationChainHandlerAddress,
            erc20DepositRecord._destinationRecipientAddress,
            erc20DepositRecord._amount);
    }

    function getERC721DepositRecord(uint originChainID, uint depositID) public view returns (
        address, address, uint, address, address, uint, bytes memory) {
        ERC721DepositRecord memory erc721DepositRecord = _erc721DepositRecords[originChainID][depositID];
        return (
            erc721DepositRecord._originChainTokenAddress,
            erc721DepositRecord._originChainHandlerAddress,
            erc721DepositRecord._destinationChainID,
            erc721DepositRecord._destinationChainHandlerAddress,
            erc721DepositRecord._destinationRecipientAddress,
            erc721DepositRecord._tokenID,
            erc721DepositRecord._data);
    }

    function getCurrentRelayerThresholdProposal() public view returns (
        uint, uint, uint, string memory) {
        return (
            _currentRelayerThresholdProposal._proposedValue,
            _currentRelayerThresholdProposal._numYes,
            _currentRelayerThresholdProposal._numNo,
            _relayerThresholdProposalStatusStrings[uint(_currentRelayerThresholdProposal._status)]);
    }

    function getDepositProposal(uint originChainID, uint depositID) public view returns (
        uint, uint, bytes32, uint, uint, string memory) {
        DepositProposal memory depositProposal = _depositProposals[originChainID][depositID];
        return (
            depositProposal._originChainID,
            depositProposal._depositID,
            depositProposal._dataHash,
            depositProposal._numYes,
            depositProposal._numNo,
            _depositProposalStatusStrings[uint(depositProposal._status)]);
    }

    function hasVoted(uint originChainID, uint depositID, address relayerAddress) public view returns (bool) {
        return _depositProposals[originChainID][depositID]._votes[relayerAddress];
    }

    function depositGeneric(
        uint         destinationChainID,
        address      destinationRecipientAddress,
        bytes memory data
    ) public {
        uint depositID = ++_depositCounts[destinationChainID];

        _genericDepositRecords[destinationChainID][depositID] = GenericDepositRecord(
            address(0),
            address(0),
            destinationChainID,
            address(0),
            destinationRecipientAddress,
            data
        );

        emit GenericDeposited(depositID);
    }

    function depositGeneric(
        address      originChainContractAddress,
        address      originChainHandlerAddress,
        uint         destinationChainID,
        address      destinationChainHandlerAddress,
        address      destinationRecipientAddress,
        bytes memory data
    ) public {
        uint depositID = ++_depositCounts[destinationChainID];

        _genericDepositRecords[destinationChainID][depositID] = GenericDepositRecord(
            originChainContractAddress,
            originChainHandlerAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            data
>>>>>>> master
        );
    }

    function hasVoted(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositID,
        address validatorAddress
    ) public view returns (bool) {
        return _depositProposals[originChainID][originChainHandlerAddress][depositID]._votes[validatorAddress];
    }

    function deposit(
        address      originChainHandlerAddress,
        bytes memory data
    ) public {
        uint depositID = ++_depositCounts[originChainHandlerAddress];
        _depositRecords[originChainHandlerAddress][depositID] = data;

        IDepositHandler depositHandler = IDepositHandler(originChainHandlerAddress);
        depositHandler.deposit(depositID, data);

        emit Deposit(CHAIN_ID, originChainHandlerAddress, depositID);
    }

<<<<<<< HEAD
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
=======
    function createDepositProposal(uint originChainID, uint depositID, bytes32 dataHash) public _onlyRelayers {
        require(_depositProposals[originChainID][depositID]._status == DepositProposalStatus.Inactive ||
        _depositProposals[originChainID][depositID]._status == DepositProposalStatus.Denied, "this proposal is either currently active or has already been passed/transferred");

        // If _depositThreshold is set to 1, then auto finalize
        if (_relayerThreshold <= 1) {
            _depositProposals[originChainID][depositID] = DepositProposal({
                _originChainID: originChainID,
                _depositID: depositID,
>>>>>>> master
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

<<<<<<< HEAD
    function voteDepositProposal(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositID,
        Vote    vote
    ) public _onlyValidators {
        DepositProposal storage depositProposal = _depositProposals[originChainID][originChainHandlerAddress][depositID];
=======
    function voteDepositProposal(uint originChainID, uint depositID, Vote vote) public _onlyRelayers {
        DepositProposal storage depositProposal = _depositProposals[originChainID][depositID];
>>>>>>> master

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
<<<<<<< HEAD
            emit DepositProposalFinalized(originChainID, originChainHandlerAddress, depositID);
        } else if (_validatorContract.getTotalValidators().sub(depositProposal._numNo) < _validatorThreshold) {
=======
            emit DepositProposalFinalized(originChainID, depositID);
        } else if (_relayerContract.getTotalRelayers().sub(depositProposal._numNo) < _relayerThreshold) {
>>>>>>> master
            depositProposal._status = DepositProposalStatus.Denied;
            emit DepositProposalFinalized(originChainID, originChainHandlerAddress, depositID);
        }

        emit DepositProposalVote(originChainID, originChainHandlerAddress, depositID, vote, depositProposal._status);
    }

    function executeDepositProposal(
        uint256      originChainID,
        address      originChainHandlerAddress,
        uint256      depositID,
        address      destinationChainHandlerAddress,
        bytes memory data
    ) public {
        DepositProposal storage depositProposal = _depositProposals[originChainID][originChainHandlerAddress][depositID];

        require(depositProposal._status != DepositProposalStatus.Inactive, "proposal is not active");
        require(depositProposal._status == DepositProposalStatus.Passed, "proposal was not passed or has already been transferred");
        require(keccak256(data) == depositProposal._dataHash, "provided data does not match proposal's data hash");

        IDepositHandler depositHandler = IDepositHandler(destinationChainHandlerAddress);
        depositHandler.executeDeposit(data);

        depositProposal._status = DepositProposalStatus.Transferred;
        emit DepositProposalExecuted(originChainID, originChainHandlerAddress, depositID);
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
pragma solidity ^0.5.12;

import "./helpers/SafeMath.sol";
import "./interfaces/IValidator.sol";
import "./interfaces/IERC20Handler.sol";
import "./interfaces/IERC721Handler.sol";
import "./interfaces/IDepositHandler.sol";

contract Bridge {
    using SafeMath for uint;

    IValidator public _validatorContract;
    uint public _validatorThreshold;
    ValidatorThresholdProposal public _currentValidatorThresholdProposal;

    enum Vote {No, Yes}
    enum ValidatorThresholdProposalStatus {Inactive, Active}
    enum DepositProposalStatus {Inactive, Active, Denied, Passed, Transferred}

    struct GenricDepositRecord {
        address _originChainTokenAddress;
        address _originChainHandlerAddress;
        uint    _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        bytes   _data;
    }

    struct ERC20DepositRecord {
        address _originChainTokenAddress;
        address _originChainHandlerAddress;
        uint    _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        uint    _amount;
    }

    struct ERC721DepositRecord {
        address _originChainTokenAddress;
        address _originChainHandlerAddress;
        uint    _destinationChainID;
        address _destinationChainHandlerAddress;
        address _destinationRecipientAddress;
        uint    _tokenID;
        bytes   _data;
    }

    struct DepositProposal {
        uint _originChainID;
        uint _depositID;
        bytes32 _dataHash;
        mapping(address => bool) _votes;
        uint _numYes;
        uint _numNo;
        DepositProposalStatus _status;
    }

    struct ValidatorThresholdProposal {
        uint                             _proposedValue;
        mapping(address => bool)         _votes;
        uint                             _numYes;
        uint                             _numNo;
        ValidatorThresholdProposalStatus _status;
    }

    // chainID => number of deposits
    mapping(uint => uint) public _depositCounts;
    // chainID => depositID => GenricDepositRecord
    mapping(uint => mapping(uint => GenricDepositRecord)) public _genericDepositRecords;
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

    modifier _onlyValidators() {
        IValidator validatorContract = IValidator(_validatorContract);
        require(validatorContract.isValidator(msg.sender));
        _;
    }

    constructor (address validatorContract) public {
        _validatorContract = IValidator(validatorContract);
    }

    function hasVoted(uint originChainID, uint depositID, address validatorAddress) public view returns (bool) {
        return _depositProposals[originChainID][depositID]._votes[validatorAddress];
    }

    function depositGeneric(
        uint         destinationChainID,
        address      destinationRecipientAddress,
        bytes memory data
    ) public {
        uint depositID = _depositCounts[destinationChainID]++;

        _genericDepositRecords[destinationChainID][depositID] = GenricDepositRecord(
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
        uint depositID = _depositCounts[destinationChainID]++;

        _genericDepositRecords[destinationChainID][depositID] = GenricDepositRecord(
            originChainContractAddress,
            originChainHandlerAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            data
        );

        emit GenericDeposited(depositID);
    }

    function depositERC20(
        address originChainTokenAddress,
        address originChainHandlerAddress,
        uint    destinationChainID,
        address destinationChainHandlerAddress,
        address destinationRecipientAddress,
        uint    amount
    ) public {
        IERC20Handler erc20Handler = IERC20Handler(originChainHandlerAddress);
        erc20Handler.depositERC20(originChainTokenAddress, msg.sender, amount);

        uint depositID = _depositCounts[destinationChainID]++;

        _erc20DepositRecords[destinationChainID][depositID] = ERC20DepositRecord(
            originChainTokenAddress,
            originChainHandlerAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            amount
        );

        emit ERC20Deposited(depositID);
    }

    function depositERC721(
        address      originChainTokenAddress,
        address      originChainHandlerAddress,
        uint         destinationChainID,
        address      destinationChainHandlerAddress,
        address      destinationRecipientAddress,
        uint         tokenID,
        bytes memory data
    ) public {
        IERC721Handler erc721Handler = IERC721Handler(originChainHandlerAddress);
        erc721Handler.depositERC721(originChainTokenAddress, msg.sender, tokenID);

        uint depositID = _depositCounts[destinationChainID]++;

        _erc721DepositRecords[destinationChainID][depositID] = ERC721DepositRecord(
            originChainTokenAddress,
            originChainHandlerAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            tokenID,
            data
        );

        emit ERC721Deposited(depositID);
    }

    function createDepositProposal(uint originChainID, uint depositID, bytes32 dataHash) public _onlyValidators {
        require(_depositProposals[originChainID][depositID]._status == DepositProposalStatus.Inactive ||
        _depositProposals[originChainID][depositID]._status == DepositProposalStatus.Denied, "this proposal is either currently active or has already been passed/transferred");

        // If _depositThreshold is set to 1, then auto finalize
        if (_validatorThreshold <= 1) {
            _depositProposals[originChainID][depositID] = DepositProposal({
                _originChainID: originChainID,
                _depositID: depositID,
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: DepositProposalStatus.Passed
                });
        } else {
            _depositProposals[originChainID][depositID] = DepositProposal({
                _originChainID: originChainID,
                _depositID: depositID,
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: DepositProposalStatus.Active
                });
        }

        // Creator always votes in favour
        _depositProposals[originChainID][depositID]._votes[msg.sender] = true;

        emit DepositProposalCreated(originChainID, depositID, dataHash);
    }

    function voteDepositProposal(uint originChainID, uint depositID, Vote vote) public _onlyValidators {
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
        }
        // Record vote
        _currentValidatorThresholdProposal._votes[msg.sender] = true;
    }

    // function voteThresholdProposal(Vote vote) external;
}
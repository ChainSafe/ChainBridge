pragma solidity ^0.5.12;

import "./helpers/SafeMath.sol";
import "./interfaces/IValidator.sol";
import "./interfaces/IERC20Handler.sol";
import "./interfaces/IERC721Handler.sol";
import "./interfaces/IDepositHandler.sol";

contract Bridge {
    using SafeMath for uint;

    address public _validatorContract;
    uint public _depositThreshold;

    enum Vote {No, Yes}
    enum ProposalStatus {Inactive, Active, Denied, Passed, Transferred}
    enum ThresholdType {Validator, Deposit}

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
        ProposalStatus _status;
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
    event DepositProposalVote(uint indexed originChainID, uint indexed depositID, Vote indexed vote, ProposalStatus status);
    event DepositProposalFinalized(uint indexed originChainID, uint indexed depositID);

    modifier _onlyValidators() {
        IValidator validatorContract = IValidator(_validatorContract);
        require(validatorContract.isValidator(msg.sender));
        _;
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
        require(_depositProposals[originChainID][depositID]._status == ProposalStatus.Inactive, "this proposal already exists");

        // If _depositThreshold is set to 1, then auto finalize
        if (_depositThreshold == 1) {
            _depositProposals[originChainID][depositID] = DepositProposal({
                _originChainID: originChainID,
                _depositID: depositID,
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: ProposalStatus.Passed
                });
        } else {
            _depositProposals[originChainID][depositID] = DepositProposal({
                _originChainID: originChainID,
                _depositID: depositID,
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: ProposalStatus.Active
                });
        }

        // Creator always votes in favour
        _depositProposals[originChainID][depositID]._votes[msg.sender] = true;

        emit DepositProposalCreated(originChainID, depositID, dataHash);
    }

    function voteDepositProposal(uint originChainID, uint depositID, Vote vote) public _onlyValidators {
        DepositProposal storage depositProposal = _depositProposals[originChainID][depositID];

        require(depositProposal._status != ProposalStatus.Inactive, "proposal is not active");
        require(depositProposal._status != ProposalStatus.Active, "proposal has been finalized");
        require(!depositProposal._votes[msg.sender], "validator has already voted");
        require(uint(vote) <= 1, "invalid vote");

        if (vote == Vote.Yes) {
            depositProposal._numYes++;
        } else {
            depositProposal._numNo++;
        }

        depositProposal._votes[msg.sender] = true;

        IValidator validatorContract = IValidator(_validatorContract);

        // Todo: Edge case if validator threshold changes?
        if (depositProposal._numYes >= _depositThreshold) {
            depositProposal._status = ProposalStatus.Passed;
            emit DepositProposalFinalized(originChainID, depositID);
        } else if (validatorContract.getTotalValidators().sub(depositProposal._numNo) < _depositThreshold) {
            depositProposal._status = ProposalStatus.Denied;
            emit DepositProposalFinalized(originChainID, depositID);
        }

        emit DepositProposalVote(originChainID, depositID, vote, depositProposal._status);
    }

    function executeDepositProposal(uint originChainID, uint depositID, address destinationChainHandlerAddress, bytes memory data) public {
        DepositProposal storage depositProposal = _depositProposals[originChainID][depositID];

        require(depositProposal._status == ProposalStatus.Passed, "proposal was not passed");
        require(keccak256(data) == depositProposal._dataHash, "provided data does not match proposal's data hash");

        IDepositHandler depositHandler = IDepositHandler(destinationChainHandlerAddress);
        depositHandler.executeDeposit(data);

        depositProposal._status = ProposalStatus.Transferred;
    }

//    function createThresholdProposal(uint value) external;
//
//    function voteThresholdProposal(Vote vote) external;
}
pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "./helpers/SafeMath.sol";
import "./interfaces/IRelayer.sol";
<<<<<<< HEAD
=======
import "./interfaces/IERC20Handler.sol";
import "./interfaces/IERC721Handler.sol";
import "./interfaces/ICentrifugeAssetHandler.sol";
>>>>>>> master
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
<<<<<<< HEAD
        bytes32                  _dataHash;
=======
        uint _destinationChainID;
        uint _depositNonce;
        bytes32 _dataHash;
>>>>>>> master
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

<<<<<<< HEAD
    // originChainHandlerAddress address => number of deposits
    mapping(address => uint256) public _depositCounts;
    // originChainHandlerAddress => depositID => bytes
    mapping(address => mapping(uint256 => bytes)) public _depositRecords;
    // originChainID => originChainHandlerAddress => depositID => depositProposal
    mapping(uint256 => mapping(address => mapping(uint256 => DepositProposal))) public _depositProposals;

=======
    // chainID => number of deposits
    mapping(uint => uint) public _depositCounts;
    // chainID => depositNonce => GenericDepositRecord
    mapping(uint => mapping(uint => GenericDepositRecord)) public _genericDepositRecords;
    // chainID => depositNonce => ERC20DepositRecord
    mapping(uint => mapping(uint => ERC20DepositRecord)) public _erc20DepositRecords;
    // chainID => depositNonce => ERC721DepositRecord
    mapping(uint => mapping(uint => ERC721DepositRecord)) public _erc721DepositRecords;
    // chainID => depositNonce => CentrifugeAssetDepositRecord
    mapping(uint => mapping(uint => CentrifugeAssetDepositRecord)) public _centrifugeDepositRecords;
    // ChainId => DepositNonce => Proposal
    mapping(uint => mapping(uint => DepositProposal)) public _depositProposals;

    event GenericDeposited(uint indexed destinationChainID, uint indexed depositNonce);
    event ERC20Deposited(uint indexed destinationChainID, uint indexed depositNonce);
    event ERC721Deposited(uint indexed destinationChainID, uint indexed depositNonce);
    event CentrifugeAssetDeposited(uint indexed destinationChainID, uint indexed depositNonce);
    event DepositProposalCreated(uint indexed destinationChainID, uint indexed depositNonce, bytes32 indexed dataHash);
    event DepositProposalVote(uint indexed destinationChainID, uint indexed depositNonce, Vote indexed vote, DepositProposalStatus status);
    event DepositProposalFinalized(uint indexed destinationChainID, uint indexed depositNonce);
>>>>>>> master
    event RelayerThresholdProposalCreated(uint indexed proposedValue);
    event RelayerThresholdProposalVote(Vote vote);
    event RelayerThresholdChanged(uint indexed newThreshold);
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

    modifier _onlyRelayers() {
        IRelayer relayerContract = IRelayer(_relayerContract);
<<<<<<< HEAD
        require(relayerContract.isRelayer(msg.sender), "sender is not a relayer");
=======
        require(relayerContract.isRelayer(msg.sender), "sender must be a relayer");
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
        );
    }

    function hasVoted(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositID,
        address relayerAddress
    ) public view returns (bool) {
        return _depositProposals[originChainID][originChainHandlerAddress][depositID]._votes[relayerAddress];
=======
    function getRelayerThreshold() public view returns (uint) {
        return _relayerThreshold;
    }

    function getDepositCount(uint destinationChainID) public view returns (uint) {
        return _depositCounts[destinationChainID];
    }

    function getGenericDepositRecord(uint destinationChainID, uint depositNonce) public view returns (
        address, address, uint, address, address, bytes memory) {
        GenericDepositRecord memory genericDepositRecord = _genericDepositRecords[destinationChainID][depositNonce];
        return (
            genericDepositRecord._originChainTokenAddress,
            genericDepositRecord._originChainHandlerAddress,
            genericDepositRecord._destinationChainID,
            genericDepositRecord._destinationChainHandlerAddress,
            genericDepositRecord._destinationRecipientAddress,
            genericDepositRecord._data);
    }

    function getERC20DepositRecord(uint destinationChainID, uint depositNonce) public view returns (
        address, address, uint, address, address, uint) {
        ERC20DepositRecord memory erc20DepositRecord = _erc20DepositRecords[destinationChainID][depositNonce];
        return (
            erc20DepositRecord._originChainTokenAddress,
            erc20DepositRecord._originChainHandlerAddress,
            erc20DepositRecord._destinationChainID,
            erc20DepositRecord._destinationChainHandlerAddress,
            erc20DepositRecord._destinationRecipientAddress,
            erc20DepositRecord._amount);
    }

    function getERC721DepositRecord(uint destinationChainID, uint depositNonce) public view returns (
        address, address, uint, address, address, uint, bytes memory) {
        ERC721DepositRecord memory erc721DepositRecord = _erc721DepositRecords[destinationChainID][depositNonce];
        return (
            erc721DepositRecord._originChainTokenAddress,
            erc721DepositRecord._originChainHandlerAddress,
            erc721DepositRecord._destinationChainID,
            erc721DepositRecord._destinationChainHandlerAddress,
            erc721DepositRecord._destinationRecipientAddress,
            erc721DepositRecord._tokenID,
            erc721DepositRecord._data);
    }

    function getCentrifugeAssetDepositRecord(uint destinationChainID, uint depositNonce) public view returns (
        address, address, uint, address, address, bytes32) {
        CentrifugeAssetDepositRecord memory centrifugeDepositRecord = _centrifugeDepositRecords[destinationChainID][depositNonce];
        return (
            centrifugeDepositRecord._originChainTokenAddress,
            centrifugeDepositRecord._originChainHandlerAddress,
            centrifugeDepositRecord._destinationChainID,
            centrifugeDepositRecord._destinationChainHandlerAddress,
            centrifugeDepositRecord._destinationRecipientAddress,
            centrifugeDepositRecord._metaDataHash);
    }

    function getCurrentRelayerThresholdProposal() public view returns (
        uint, uint, uint, string memory) {
        return (
            _currentRelayerThresholdProposal._proposedValue,
            _currentRelayerThresholdProposal._numYes,
            _currentRelayerThresholdProposal._numNo,
            _relayerThresholdProposalStatusStrings[uint(_currentRelayerThresholdProposal._status)]);
    }

    function getDepositProposal(uint destinationChainID, uint depositNonce) public view returns (
        uint, uint, bytes32, uint, uint, string memory) {
        DepositProposal memory depositProposal = _depositProposals[destinationChainID][depositNonce];
        return (
            depositProposal._destinationChainID,
            depositProposal._depositNonce,
            depositProposal._dataHash,
            depositProposal._numYes,
            depositProposal._numNo,
            _depositProposalStatusStrings[uint(depositProposal._status)]);
    }

    function hasVoted(uint destinationChainID, uint depositNonce, address relayerAddress) public view returns (bool) {
        return _depositProposals[destinationChainID][depositNonce]._votes[relayerAddress];
    }

    function depositGeneric(
        uint         destinationChainID,
        address      destinationRecipientAddress,
        bytes memory data
    ) public {
        uint depositNonce = ++_depositCounts[destinationChainID];

        _genericDepositRecords[destinationChainID][depositNonce] = GenericDepositRecord(
            address(0),
            address(0),
            destinationChainID,
            address(0),
            destinationRecipientAddress,
            data
        );

        emit GenericDeposited(destinationChainID, depositNonce);
    }

    function depositGeneric(
        address      originChainContractAddress,
        address      originChainHandlerAddress,
        uint         destinationChainID,
        address      destinationChainHandlerAddress,
        address      destinationRecipientAddress,
        bytes memory data
    ) public {
        uint depositNonce = ++_depositCounts[destinationChainID];

        _genericDepositRecords[destinationChainID][depositNonce] = GenericDepositRecord(
            originChainContractAddress,
            originChainHandlerAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            data
        );

        emit GenericDeposited(destinationChainID, depositNonce);
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

        uint depositNonce = ++_depositCounts[destinationChainID];

        _erc20DepositRecords[destinationChainID][depositNonce] = ERC20DepositRecord(
            originChainTokenAddress,
            originChainHandlerAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            amount
        );

        emit ERC20Deposited(destinationChainID, depositNonce);
>>>>>>> master
    }

    function deposit(
        address      originChainHandlerAddress,
        bytes memory data
    ) public {
<<<<<<< HEAD
        uint depositID = ++_depositCounts[originChainHandlerAddress];
        _depositRecords[originChainHandlerAddress][depositID] = data;

        IDepositHandler depositHandler = IDepositHandler(originChainHandlerAddress);
        depositHandler.deposit(depositID, data);

        emit Deposit(CHAIN_ID, originChainHandlerAddress, depositID);
    }

    function createDepositProposal(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositID,
        bytes32 dataHash
    ) public _onlyRelayers {
        require(
            _depositProposals[originChainID][originChainHandlerAddress][depositID]._status == DepositProposalStatus.Inactive ||
            _depositProposals[originChainID][originChainHandlerAddress][depositID]._status == DepositProposalStatus.Denied,
            "proposal is either currently active or has already been passed/transferred"
        );

        // If _depositThreshold is set to 1, then auto finalize
        if (_relayerThreshold <= 1) {
            _depositProposals[originChainID][originChainHandlerAddress][depositID] = DepositProposal({
=======
        IERC721Handler erc721Handler = IERC721Handler(originChainHandlerAddress);
        erc721Handler.depositERC721(originChainTokenAddress, msg.sender, tokenID);

        uint depositNonce = ++_depositCounts[destinationChainID];

        _erc721DepositRecords[destinationChainID][depositNonce] = ERC721DepositRecord(
            originChainTokenAddress,
            originChainHandlerAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            tokenID,
            data
        );

        emit ERC721Deposited(destinationChainID, depositNonce);
    }

    function depositCentrifugeAsset(
        address originChainContractAddress,
        address originChainHandlerAddress,
        uint    destinationChainID,
        address destinationChainHandlerAddress,
        address destinationRecipientAddress,
        bytes32 metaDataHash
    ) public {
        ICentrifugeAssetHandler centrifugeAssetHandler = ICentrifugeAssetHandler(originChainHandlerAddress);
        centrifugeAssetHandler.depositAsset(metaDataHash);

        uint depositNonce = ++_depositCounts[destinationChainID];

        _centrifugeDepositRecords[destinationChainID][depositNonce] = CentrifugeAssetDepositRecord(
            originChainContractAddress,
            originChainHandlerAddress,
            destinationChainID,
            destinationChainHandlerAddress,
            destinationRecipientAddress,
            metaDataHash
        );

        emit CentrifugeAssetDeposited(destinationChainID, depositNonce);
    }

    function createDepositProposal(uint destinationChainID, uint depositNonce, bytes32 dataHash) public _onlyRelayers {
        require(_depositProposals[destinationChainID][depositNonce]._status == DepositProposalStatus.Inactive ||
        _depositProposals[destinationChainID][depositNonce]._status == DepositProposalStatus.Denied, "this proposal is either currently active or has already been passed/transferred");

        // If _depositThreshold is set to 1, then auto finalize
        if (_relayerThreshold <= 1) {
            _depositProposals[destinationChainID][depositNonce] = DepositProposal({
                _destinationChainID: destinationChainID,
                _depositNonce: depositNonce,
>>>>>>> master
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: DepositProposalStatus.Passed
            });
        } else {
<<<<<<< HEAD
            _depositProposals[originChainID][originChainHandlerAddress][depositID] = DepositProposal({
=======
            _depositProposals[destinationChainID][depositNonce] = DepositProposal({
                _destinationChainID: destinationChainID,
                _depositNonce: depositNonce,
>>>>>>> master
                _dataHash: dataHash,
                _numYes: 1, // Creator always votes in favour
                _numNo: 0,
                _status: DepositProposalStatus.Active
            });
        }

        // Creator always votes in favour
<<<<<<< HEAD
        _depositProposals[originChainID][originChainHandlerAddress][depositID]._votes[msg.sender] = true;

        emit DepositProposalCreated(originChainID, originChainHandlerAddress, depositID, dataHash);
    }

    function voteDepositProposal(
        uint256 originChainID,
        address originChainHandlerAddress,
        uint256 depositID,
        Vote    vote
    ) public _onlyRelayers {
        DepositProposal storage depositProposal = _depositProposals[originChainID][originChainHandlerAddress][depositID];
=======
        _depositProposals[destinationChainID][depositNonce]._votes[msg.sender] = true;

        emit DepositProposalCreated(destinationChainID, depositNonce, dataHash);
    }

    function voteDepositProposal(uint destinationChainID, uint depositNonce, Vote vote) public _onlyRelayers {
        DepositProposal storage depositProposal = _depositProposals[destinationChainID][depositNonce];
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
        } else if (_relayerContract.getTotalRelayers().sub(depositProposal._numNo) < _relayerThreshold) {
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
=======
            emit DepositProposalFinalized(destinationChainID, depositNonce);
        } else if (_relayerContract.getTotalRelayers().sub(depositProposal._numNo) < _relayerThreshold) {
            depositProposal._status = DepositProposalStatus.Denied;
            emit DepositProposalFinalized(destinationChainID, depositNonce);
        }

        emit DepositProposalVote(destinationChainID, depositNonce, vote, depositProposal._status);
    }

    function executeDepositProposal(uint destinationChainID, uint depositNonce, address destinationChainHandlerAddress, bytes memory data) public {
        DepositProposal storage depositProposal = _depositProposals[destinationChainID][depositNonce];
>>>>>>> master

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
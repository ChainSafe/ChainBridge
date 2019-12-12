pragma solidity 0.5.11;

contract Home {

    uint public voteThreshold;

    struct ChainStruct {
        // The total count of processed tx from this chain`
        uint count;
    }
    enum Vote {Null, Yes, No}

    // Used by validators to vote on deposits
    // A validator should submit a 32 byte keccak hash of the deposit data
    // The hash will be verified when finalizing the deposit
    struct Proposal {
        // 32 byte hash of the deposit object
        bytes32 hash;
        // Incremented based on the origin chain
        uint id;
        // Chain that the tx originated from
        uint originChain;
        // Keeps track if a user has voted
        mapping(address => bool) votes;
        // Number of votes yes
        uint numYes;
        // Number of votes no
        uint numNo;
	    // If proposal has passed
	    bool finalized;
    }

    // List of validators
    mapping(address => bool) Validators;

    // Keep track of blockchains
    // ChainId => ChainStruct
    mapping(uint => ChainStruct) Chains;

    // keep track of all proposed deposits per origin chain
    // ChainId => DepositId => Proposal
    mapping(uint => mapping(uint => Proposal)) Proposals;

    // Ensure user is a validator
    modifier _isValidator() {
        require(Validators[msg.sender], "Sender is not a validator.");
        _;
    }

    /**
     * @param _addrs - Bridge validator addresses
     * @param _threshold - The number of votes required for a vote to pass
     */
    constructor ( address[] memory _addrs, uint _threshold) public {
        // set the validators
        for (uint i = 0; i<_addrs.length; i++) {
          Validators[_addrs[i]] = true;
        }
    	// Set the validator threshold
        voteThreshold = _threshold;
    }

    /**
     * @notice Validators propose to make a deposit, this isn't final and requires the validators to reach majority consensus
     * @param _hash - 32 bytes hash of the Deposit data
     * @param _id - The deposit id generated from the origin chain
     * @param _originChain - The chain id from which the deposit was originally made
     */
    // TODO; if the proposal has already been made should we vote? If they're submitting we can assume they're in favour?
    function submitProposal(bytes32 _hash, uint _depositId, uint _originChain) public _isValidator {
        // Ensure this proposal hasn't already been made
	    require(Proposals[_originChain][_depositId].id != _id, "Proposal already exists");

        // Create Proposal
        Proposals[_originChain][_depositId] = Proposal({
            hash: _hash,
            id: _depositId,
            originChain: _originChain,
            numYes: 0,
            numNo: 0,
            finalized: false
        });
        // The creator always votes in favour
        Proposals[_originChain][_depositId].votes[msg.sender] = true;
        Proposals[_originChain][_depositId].numYes++;
    }

    /**
     * @notice - Used to cast a vote on a given proposal
     * @param _originChainId - The id of the origin chain for a given proposal
     * @param _depositId - The id assigned to a deposit, generated on the origin chain
     * @param _vote - uint from 0-2 representing the casted vote
     */
    function vote(uint _originChainId, uint _depositId, Vote _vote) public _isValidator {
        require(!Proposals[_originChainId][_depositId].finalized, "Proposal has already been finalized!");
        require(!Proposals[_originChainId][_depositId].votes[msg.sender], "User has already voted!");
        require(uint(_vote) > 0 && uint(_vote) <= 2, "Invalid vote!");
        // Add validator signoff
        if (_vote == Vote.Yes) {
            Proposals[_originChainId][_depositId].numYes++;
        } else if (_vote == Vote.No) {
            Proposals[_originChainId][_depositId].numNo--;
        }

        // Mark that the user voted
        Proposals[_originChainId][_depositId].votes[msg.sender] = true;

        // Check if the threshold has been met
        if (Proposals[_originChainId][_depositId].numYes + Proposals[_originChainId][_depositId].numNo >= voteThreshold) {
            Proposal[_originChainId][_depositId].finalized = true;
        }
    }

    /**
     * Executes a deposit, anyone can execute this as long as they pass in the correct _data
     */
    function executeDeposit(bytes _data, uint _originChainId, uint _proposalId) public {
        proposal = Proposals[_originChainId][_proposalId];
        require(proposal.numYes >= voteThreshold, "Vote has not passed!"); // Check that voted passed
        // Ensure that the incoming data is the same as the hashed data from the proposal
        require(keccak256(_data) == proposal.hash, "Incorrect data supplied for hash");

        // TODO unpack the _data
    }

    // Todo add MultiSig checks
    function addValidator(address _addr) public {
        Validators[_addr] = true;
    }

    // Todo add MultiSig checks
    function removeValidator(address _addr) public {
        Validators[_addr] = false;
    }
}
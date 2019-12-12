pragma solidity 0.5.12;

contract Home {

    uint public voteDepositThreshold;
    uint public voteValidatorThreshold;

    struct ChainStruct {
        // The total count of processed tx from this chain`
        uint count;
    }

    enum Vote {Yes, No}
    enum ValidatorVoteType {Add, Remove}

    // Used by validators to vote on deposits
    // A validator should submit a 32 byte keccak hash of the deposit data
    // The hash will be verified when finalizing the deposit
    struct DepositProposal {
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

    // Proposal to add/remove a bridge validator
    struct ValidatorProposal {
        // Address of the proposed validator
        address validator;
        // validator action
        ValidatorVoteType action;
        // Keeps track if a user has voted
        mapping(address => bool) votes;
        // Number of votes in favour
        uint numYes;
        // Number of votes against
        uint numNo;
        // Total votes
        uint totalVotes;
    }

    // List of validators
    mapping(address => bool) Validators;
    uint TotalValidatos
    // Validator proposals
    // Address = validator to add/remove
    mapping(address => ValidatorProposal) ValidatorProposals;

    // keep track of all proposed deposits per origin chain
    // ChainId => DepositId => Proposal
    mapping(uint => mapping(uint => DepositProposal)) DepositProposals;

    // Ensure user is a validator
    modifier _isValidator() {
        require(Validators[msg.sender], "Sender is not a validator.");
        _;
    }

    /**
     * @param _addrs - Bridge validator addresses
     * @param _threshold - The number of votes required for a vote to pass
     */
    constructor ( address[] memory _addrs, uint _depositThreshold, uint _validatorThreshold) public {
        // set the validators
        for (uint i = 0; i<_addrs.length; i++) {
          Validators[_addrs[i]] = true;
        }
        // Set total validators
        TotalValidators = addrs.length;

    	// Set the thresholds
        voteDepositThreshold = _depositThreshold;
        voteValidatorThreshold = _validatorThreshold;
    }

    /**
     * Validators propose to make a deposit, this isn't final and requires the validators to reach majority consensus
     * @param _hash - 32 bytes hash of the Deposit data
     * @param _id - The deposit id generated from the origin chain
     * @param _originChain - The chain id from which the deposit was originally made
     */
    // TODO; if the proposal has already been made should we vote? If they're submitting we can assume they're in favour?
    function createDepositProposal(bytes32 _hash, uint _depositId, uint _originChain) public _isValidator {
        // Ensure this proposal hasn't already been made
	    require(DepositProposals[_originChain][_depositId].id != _depositId, "Proposal already exists");

        // Create Proposal
        DepositProposals[_originChain][_depositId] = DepositProposal({
            hash: _hash,
            id: _depositId,
            originChain: _originChain,
            numYes: 1, // The creator always votes in favour
            numNo: 0,
            finalized: false
        });
        // The creator always votes in favour
        DepositProposals[_originChain][_depositId].votes[msg.sender] = true;
    }

    /**
     * Used to cast a vote on a given proposal
     * @param _originChainId - The id of the origin chain for a given proposal
     * @param _depositId - The id assigned to a deposit, generated on the origin chain
     * @param _vote - uint from 0-2 representing the casted vote
     */
    function voteDepositProposal(uint _originChainId, uint _depositId, Vote _vote) public _isValidator {
        require(!DepositProposals[_originChainId][_depositId].finalized, "Proposal has already been finalized!");
        require(!DepositProposals[_originChainId][_depositId].votes[msg.sender], "User has already voted!");
        require(uint(_vote) <= 1, "Invalid vote!");

        // Add vote signoff
        if (_vote == Vote.Yes) {
            DepositProposals[_originChainId][_depositId].numYes++;
        } else if (_vote == Vote.No) {
            DepositProposals[_originChainId][_depositId].numNo++;
        }

        // Mark that the validator voted
        DepositProposals[_originChainId][_depositId].votes[msg.sender] = true;

        // Check if the threshold has been met
        // Todo: Edge case if validator threshold changes?
        if (DepositProposals[_originChainId][_depositId].numYes >= voteDepositThreshold ||
            Validators.length - DepositProposals[_originChainId][_depositId].numNo >= voteDepositThreshold) {
            DepositProposals[_originChainId][_depositId].finalized = true;
        }
    }

    /**
     * Executes a deposit, anyone can execute this as long as they pass in the correct _data
     * @param _data - Deposit data to be unpacked
     * @param _originChainId - The chain id representing where the deposit originated from
     * @param _depositId - The id generated from the origin chain
     */
    function executeDeposit(bytes memory _data, uint _originChainId, uint _depositId) public {
        DepositProposal storage proposal = DepositProposals[_originChainId][_depositId];
        require(proposal.numYes >= voteDepositThreshold, "Vote has not passed!"); // Check that voted passed
        // Ensure that the incoming data is the same as the hashed data from the proposal
        require(keccak256(_data) == proposal.hash, "Incorrect data supplied for hash");

        // TODO unpack the _data
    }

    /**
     * Creates a new proposal to add or remove a validator
     * @param _addr - Address of the validator to be added or removed
     * @param _action - Action to either remove or add validator
     */
    function createValidatorProposal(address _addr,  ValidatorVoteType _action) public _isValidator {
        require(_action <= 1, "Action out of the vote enum range!");
        require(_action == ValidatorVoteType.Remove && Validators[_addr], "Validator is not active!");
        require(_action == Validators.Add && !Validators[_addr], "Validator is already active!");
        require(ValidatorProposals[_addr].numYes + ValidatorProposals[_addr].numNo > 0, "There is already an active proposal!");

        ValidatorProposals[_addr] = ValidatorProposal({
            validator: _addr,
            action: _action,
            numYes: 1, // Creator must vote in favour
            numNo: 0,
            totalVots: 1 // Creator must vote (below)
        });
    }

    /**
     * Casts vote to add or remove a validator, if the vote succeeds it will perform the action
     * @param _addr - Address of the validator to be added or removed
     * @param _vote - Vote to either remove or add validator
     */
    function voteValidatorProposal(address _addr, Vote _vote) public _isValidator {
        require(!ValidatorProposals[_addr], "Vote has ended");
        require(_vote <= 1, "Vote out of the vote enum range!");

        // Cast vote
        if (_vote == Vote.Yes) {
            ValidatorProposals[_addr].numYes++;
        } else {
            ValidatorProposals[_addr].numNo++;
        }

        // Incrememnt vote
        ValidatorProposals[_addr].totalVotes++;

        // Record vote
        ValidatorProposals[_addr].votes[msg.sender] = true;

        // Check if vote has met the threshold
        if (ValidatorProposals[_addr].numYes >= voteValidatorThreshold ||
            Validators.length - ValidatorProposals[_addr].numNo >= voteValidatorThreshold) {

            // Vote succeeded, perform action
            if (ValidatorProposals[_addr].numYes > ValidatorProposals[_addr].numNo) {
                if (ValidatorProposals[_addr] == ValidatorVoteType.Add) {
                    // Add validator
                    Validators[_addr] = true;
                    TotalValidators++;
                } else {
                    // Remove validator
                    Validators[_addr] = false;
                    TotalValidators--;
                }
            }
        }
    }
}
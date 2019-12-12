pragma solidity 0.5.11;

contract Home {

    uint public voteThreshold;

    struct ChainStruct {
        // The total count of processed tx from this chain`
        uint count;
    }

    enum Vote {Yes, No}
    enum ValidatorVote {Add, Remove}

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
        ValidatorVote action;
        // Keeps track if a user has voted
        mapping(address => bool) votes;
        // Number of votes in favour
        uint numYes;
        // Number of votes against
        uint numNo;
        // If the proposal ended
        bool finalized;
    }

    // List of validators
    mapping(address => bool) Validators;

    // Validator proposals
    mapping(address => ValidatorProposal) ValidatorProposals;

    // keep track of all proposed deposits per origin chain
    // ChainId => DepositId => Proposal
    mapping(uint => mapping(uint => DepositProposal)) Proposals;

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
     * Validators propose to make a deposit, this isn't final and requires the validators to reach majority consensus
     * @param _hash - 32 bytes hash of the Deposit data
     * @param _id - The deposit id generated from the origin chain
     * @param _originChain - The chain id from which the deposit was originally made
     */
    // TODO; if the proposal has already been made should we vote? If they're submitting we can assume they're in favour?
    function createDepositProposal(bytes32 _hash, uint _depositId, uint _originChain) public _isValidator {
        // Ensure this proposal hasn't already been made
	    require(Proposals[_originChain][_depositId].id != _id, "Proposal already exists");

        // Create Proposal
        Proposals[_originChain][_depositId] = DepositProposal({
            hash: _hash,
            id: _depositId,
            originChain: _originChain,
            numYes: 1, // The creator always votes in favour
            numNo: 0,
            finalized: false
        });
        // The creator always votes in favour
        Proposals[_originChain][_depositId].votes[msg.sender] = true;
    }

    /**
     * Used to cast a vote on a given proposal
     * @param _originChainId - The id of the origin chain for a given proposal
     * @param _depositId - The id assigned to a deposit, generated on the origin chain
     * @param _vote - uint from 0-2 representing the casted vote
     */
    function voteDepositProposal(uint _originChainId, uint _depositId, Vote _vote) public _isValidator {
        require(!Proposals[_originChainId][_depositId].finalized, "Proposal has already been finalized!");
        require(!Proposals[_originChainId][_depositId].votes[msg.sender], "User has already voted!");
        require(uint(_vote) <= 1, "Invalid vote!");

        // Add vote signoff
        if (_vote == Vote.Yes) {
            Proposals[_originChainId][_depositId].numYes++;
        } else if (_vote == Vote.No) {
            Proposals[_originChainId][_depositId].numNo++;
        }

        // Mark that the validator voted
        Proposals[_originChainId][_depositId].votes[msg.sender] = true;

        // Check if the threshold has been met
        if (Proposals[_originChainId][_depositId].numYes + Proposals[_originChainId][_depositId].numNo >= voteThreshold) {
            Proposals[_originChainId][_depositId].finalized = true;
        }
    }

    /**
     * Executes a deposit, anyone can execute this as long as they pass in the correct _data
     * @param _data - Deposit data to be unpacked
     * @param _originChainId - The chain id representing where the deposit originated from
     * @param _depositId - The id generated from the origin chain
     */
    function executeDeposit(bytes _data, uint _originChainId, uint _depositId) public {
        proposal = Proposals[_originChainId][_depositId];
        require(proposal.numYes >= voteThreshold, "Vote has not passed!"); // Check that voted passed
        // Ensure that the incoming data is the same as the hashed data from the proposal
        require(keccak256(_data) == proposal.hash, "Incorrect data supplied for hash");

        // TODO unpack the _data
    }

    /**
     * Creates a new proposal to add or remove a validator
     * @param _addr - Address of the validator to be added or removed
     * @param _action - Action to either remove or add validator
     */
    function createValidatorProposal(address _addr,  ValidatorVote _action) public _isValidator {
        require(_action >= 1, "Action out of the vote enum range!");
        if (_action == ValidatorVote.Remove) {
            // A validator must be present to remove them
            require(Validators[_addr], "Validator is not active!");
        } else if (_action == Validator.Add) {
            // A validator must not be present to add them
            require(!Validators[_addr], "Validator is already active!");
        }
        // Ensure there are no active proposals
        require(ValidatorProposals[_addr].finalized, "There is already an active proposal!");

        ValidatorProposals[_addr] = ValidatorProposal({
            validator: _addr,
            action: _action,
            numYes: 0,
            numNo: 0,
            finalized: false
        });

        // Cast vote based on action
        if (_action == ValidatorVote.Add) {
            ValidatorProposals[_addr].numYes++;
        } else if (_action == ValidatorVote.Remove) {
            ValidatorProposals[_addr].numNo++;
        }
    }

    /**
     * Casts vote to add or remove a validator, if the vote succeeds it will perform the action
     * @param _addr - Address of the validator to be added or removed
     * @param _action - Action to either remove or add validator
     */
    function voteValidatorProposal(address _addr, ValidatorVote _vote) public _isValidator {
        require(!ValidatorProposals[_addr].finalized, "Vote has ended");
        require(_action >= 1, "Action out of the vote enum range!");

        // Cast vote based on action
        if (_action == ValidatorVote.Add) {
            ValidatorProposals[_addr].numYes++;
        } else if (_action == ValidatorVote.Remove) {
            ValidatorProposals[_addr].numNo++;
        }

        // Record vote
        ValidatorProposals[_addr].votes[msg.sender] = true;

        // Check if vote has met the threshold
        if (ValidatorProposals[_addr].numYes + ValidatorProposals[_addr].numNo >= voteThreshold) {
            ValidatorProposals[_addr].finalized = true;

            // Vote succeeded, perform action
            if (ValidatorProposals[_addr].numYes > ValidatorProposals[_addr].numNo) {
                if (ValidatorProposals[_addr] == ValidatorVote.Add) {
                    // Add validator
                    Validators[_addr] = true;
                } else {
                    // Remove validator
                    Validators[_addr] = false;
                }
            }
        }
    }
}
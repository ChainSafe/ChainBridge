pragma solidity 0.5.13;

contract Home {

    uint public validatorThreshold;
    
    address public admin;

    struct ChainStruct {
        // The total count of processed tx from this chain`
        uint count;
    }
    
    // Used by validators to vote on deposits
    // A validator should submit a 32 byte keccak hash of the deposit data
    // The hash will be verified when finalizing the deposit
    struct Proposal {
        // 32 byte hash of the deposit object
        bytes32 hash;
        // Incremented based on the origin chain
        uint id;
        // Chain that the tx originated from
        uint fromChain;
        // n-m signatures
        mapping(address => bool) votes;
	// If proposal has passed
	bool finalized;
    }

    // After a succesul vote, a validator will submit a Deposit that will be verified against the 32 byte hash
    // found in Proposal.hash
    struct Deposit {
        // TODO change :: Arbitrary data 
        bytes data;
        // Incremented based on the origin chain
        uint id;
        // Chain that the tx originated from
        uint fromChain;
    }

    // List of validators
    mapping(address => bool) Validators;

    // Keep track of blockchains
    // ChainId => ChainStruct
    mapping(uint => ChainStruct) Chains;

    // keep track of all proposed deposits per origin chain
    // ChainId => DepositId => Proposal
    mapping(uint => mapping(uint => Proposal)) Proposals;    

    // Keep track of all finalized deposits per origin chain
    // ChainId => DepositId => Deposit
    mapping(uint => mapping(uint => Deposit)) Deposits;

    // Ensure user is a validator
    modifier _isValidator() {
        require(Validators[msg.sender] || msg.sender == admin, "Sender is not a validator.");
        _;
    }

    constructor (address[] _addrs, uint _threshold) {
        // set the validators
        for (uint i=0; i<_addrs.length; i++) {
          Validators[_addrs[i]] = true;
        }
        
	// Set the validator threshold
        validatorThreshold = _threshold;
    }

    // Validators propose to make a deposit, this isn't final and requires the validators to reach majority consensus
    // TODO; if the deposit has already been made should we vote? 
    function submitProposal(bytes32 _hash, uint _id, uint _fromChain) public _isValidator {
        // Ensure this deposit hasn't already been made
	require(Deposits[_from][_id].id !== _id);

        // Create deposit struct
        deposit = Proposal({
            hash: _hash,
            id: _id,
            fromChain: _fromChain,
        })

	// Whoever makes the deposit casts the first vote
	deposit.votes[msg.sender] = true; 

        // Assign deposit
        // Todo can this be overridden?
        Deposits[_from][_id] = deposit;
    }

    // Votes on a proposed deposit
    function vote(uint _chainId, uint id) public _isValidator {

        // Add validator signoff
        Deposits[_chainId][_id].signatures.push(msg.sender);

        // Check if deposit has already been finalized
        if (!Deposits[_chainId][_id].finalized) {

            // Check if the threshold has been met
            if (Deposits[_chainId][_id].signatures.length >= validatorThreshold) {
                /**
                 * TODO
                 * - Check type of tx
                 * - Handle the type accordingly
                 * - eg: ERC20.transfer(_to, _value);
                 */
                 Deposits[_chainId][_id].finalized = true;
            }
        }
    }

    // Todo add MultiSig checks
    function addValidator(address _addr) {
        Validators[_addr] = true;
    }

    // Todo add MultiSig checks
    function removeValidator(address _addr) {
        Validators[_addr] = false;
    }
}

//TODO finish
contract MultiSig {

    // MultiSig signers
    mapping(address => bool) signers;

    // All votes
    mapping(uint => Vote) votes;
    
    modifier _isSigner {
        require(signers[msg.sender], "Sender is not a signer!");
        _;
    }

    constructor(address[] _signers) {
        for (uint i=0; _signers.length; i++) {
            signers[_signers[i]] = true;
        }
    }

    function vote() public _isSigner {}

    function addSigner(address _addr) public _isSigner {}

    function removeSigner(address _addr) public _isSigner {}
}



pragma solidity 0.6.4;

import "./interfaces/IHandler.sol";

/**
 * @title Receiver
 * @dev The receiver is found on the destination chain,
 * it can only be written to by the relayers, and is
 * the location where a relayer would write the deposit
 * too. A receiver doesn't directly handle the asset,
 * rather it sends it to the respective handler.
 */
contract Receiver {

    // These are the required number of YES votes for the respectful proposals
    uint public DepositThreshold;
    uint public RelayerThreshold;

    enum Vote {Yes, No}
    enum RelayerActionType {Add, Remove}
    enum VoteStatus {Inactive, Active, Finalized, Transferred}
    enum ThresholdType {Relayer, Deposit}

    // Used by relayers to vote on deposits
    // A relayer should submit a 32 byte keccak hash of the deposit data
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
	    VoteStatus status;
    }

    // Proposal to add/remove a bridge relayer
    struct RelayerProposal {
        // Address of the proposed relayer
        address relayer;
        // relayer action
        RelayerActionType action;
        // Keeps track if a user has voted
        mapping(address => bool) votes;
        // Number of votes in favour
        uint numYes;
        // Number of votes against
        uint numNo;
        // Status of a vote
        VoteStatus status;
    }

    // Proposal to change threshold
    struct ThresholdProposal {
        // new value to change to
        uint newValue;
        // Type of vote
        ThresholdType proposalType;
        // Keeps track if a user has voted
        mapping(address => bool) votes;
        // Number of votes in favour
        uint numYes;
        // Number of votes against
        uint numNo;
        // Status of a vote
        VoteStatus status;
    }

    // List of relayers
    mapping(address => bool) public Relayers;
    uint public TotalRelayers;

    // Current threshold proposals
    // ThresholdType => ThersholdProposal
    mapping(uint => ThresholdProposal) public ThresholdProposals;

    // Relayer proposals
    // Address = relayer to add/remove
    mapping(address => RelayerProposal) public RelayerProposals;

    // keep track of all proposed deposits per origin chain
    // ChainId => DepositId => Proposal
    mapping(uint => mapping(uint => DepositProposal)) public DepositProposals;

    // Ensure user is a relayer
    modifier _isRelayer() {
        require(Relayers[msg.sender], "Sender is not a relayer.");
        _;
    }

    event DepositProposalCreated(bytes32 _hash, uint _depositId, uint _originChain, VoteStatus _voteStatus);
    event DepositProposalVote(uint _originChainId, uint _depositId, Vote _vote, VoteStatus _voteStatus);
    event DepositExecuted(uint _originChainId, uint _depositId, address _to);

    /**
     * @param _addrs - Bridge relayer addresses
     * @param _depositThreshold - The number of votes required for a deposit vote to pass
     * @param _relayerThreshold - The number of votes required for a relayer vote to pass
     */
    constructor (address[] memory _addrs, uint _depositThreshold, uint _relayerThreshold) public {
        require(_depositThreshold > 0, "Deposit threshold must be greater than 0!");
        require(_relayerThreshold > 0, "Relayer threshold must be greater than 0!");
        // set the relayers
        for (uint i = 0; i<_addrs.length; i++) {
          Relayers[_addrs[i]] = true;
        }
        // Set total relayers
        TotalRelayers = _addrs.length;

    	// Set the thresholds
        DepositThreshold = _depositThreshold;
        RelayerThreshold = _relayerThreshold;
    }

    /**
     * Relayers propose to make a deposit, this isn't final and requires the relayers to reach majority consensus
     * @param _hash - 32 bytes hash of the Deposit data
     * @param _depositId - The deposit id generated from the origin chain
     * @param _originChain - The chain id from which the deposit was originally made
     */
    function createDepositProposal(bytes32 _hash, uint _depositId, uint _originChain) public _isRelayer {
        // Ensure this proposal hasn't already been made
	    require(DepositProposals[_originChain][_depositId].status == VoteStatus.Inactive, "A proposal already exists!");

        if (DepositThreshold == 1) {
            // If the threshold is set to 1 auto finalize
            DepositProposals[_originChain][_depositId] = DepositProposal({
                hash: _hash,
                id: _depositId,
                originChain: _originChain,
                numYes: 1, // The creator always votes in favour
                numNo: 0,
                status: VoteStatus.Finalized
            });
        } else {
            // Create Proposal
            DepositProposals[_originChain][_depositId] = DepositProposal({
                hash: _hash,
                id: _depositId,
                originChain: _originChain,
                numYes: 1, // The creator always votes in favour
                numNo: 0,
                status: VoteStatus.Active
            });
        }
        // The creator always votes in favour
        DepositProposals[_originChain][_depositId].votes[msg.sender] = true;
        // Trigger event
        emit DepositProposalCreated(_hash, _depositId, _originChain, DepositProposals[_originChain][_depositId].status);
    }

    /**
     * Used to cast a vote on a given proposal
     * @param _originChainId - The id of the origin chain for a given proposal
     * @param _depositId - The id assigned to a deposit, generated on the origin chain
     * @param _vote - uint from 0-1 representing the casted vote
     */
    function voteDepositProposal(uint _originChainId, uint _depositId, Vote _vote) public _isRelayer {
        require(DepositProposals[_originChainId][_depositId].status >= VoteStatus.Active, "There is no active proposal!");
        require(DepositProposals[_originChainId][_depositId].status < VoteStatus.Finalized, "Proposal has already been finalized!");
        require(!DepositProposals[_originChainId][_depositId].votes[msg.sender], "Relayer has already voted!");
        require(uint(_vote) <= 1, "Invalid vote!");

        // Add vote signoff
        if (_vote == Vote.Yes) {
            DepositProposals[_originChainId][_depositId].numYes++;
        } else if (_vote == Vote.No) {
            DepositProposals[_originChainId][_depositId].numNo++;
        }

        // Mark that the relayer voted
        DepositProposals[_originChainId][_depositId].votes[msg.sender] = true;

        // Check if the threshold has been met
        // Todo: Edge case if relayer threshold changes?
        if (DepositProposals[_originChainId][_depositId].numYes >= DepositThreshold ||
            TotalRelayers - DepositProposals[_originChainId][_depositId].numNo < DepositThreshold) {
            DepositProposals[_originChainId][_depositId].status = VoteStatus.Finalized;
        }
        // Triger event
        emit DepositProposalVote(_originChainId, _depositId, _vote, DepositProposals[_originChainId][_depositId].status);
    }

    /**
     * Executes a deposit, anyone can execute this as long as they pass in the correct _data
     * @param _originChainId - The chain id representing where the deposit originated from
     * @param _depositId - The id generated from the origin chain
     * @param _to - The receiver contract address
     * @param _data - Deposit data to be unpacked
     */
    function executeDeposit(uint _originChainId, uint _depositId, address _to, bytes memory _data) public {
        DepositProposal storage proposal = DepositProposals[_originChainId][_depositId];
        require(proposal.status == VoteStatus.Finalized, "Vote has not finalized!");
        require(proposal.numYes >= DepositThreshold, "Vote has not passed!"); // Check that voted passed
        // Ensure that the incoming data is the same as the hashed data from the proposal
        require(keccak256(_data) == proposal.hash, "Incorrect data supplied for hash");

        // TODO use generic receiver
        IHandler handler = IHandler(_to);
        handler.executeTransfer(_originChainId, _data);

        // Mark deposit Transferred
        DepositProposals[_originChainId][_depositId].status = VoteStatus.Transferred;

        // Trigger event
        emit DepositExecuted(_originChainId, _depositId, _to);
    }

    /**
     * Creates a new proposal to add or remove a relayer
     * Note: There is no `Finalized` state for Relayer proposals they are either active or inacive
     * @param _addr - Address of the relayer to be added or removed
     * @param _action - Action to either remove or add relayer
     */
    function createRelayerProposal(address _addr,  RelayerActionType _action) public _isRelayer {
        require(uint(_action) <= 1, "Action out of the vote enum range!");
        require(!(_action == RelayerActionType.Remove && Relayers[_addr] == false), "Relayer is not active!");
        require(!(_action == RelayerActionType.Add && Relayers[_addr] == true), "Relayer is already active!");
        require(RelayerProposals[_addr].status == VoteStatus.Inactive, "There is already an active proposal!");

        if (RelayerThreshold == 1) {
            // If the threshold is set to 1, auto complete
            RelayerProposals[_addr] = RelayerProposal({
                relayer: _addr,
                action: _action,
                numYes: 1, // Creator must vote in favour
                numNo: 0,
                status: VoteStatus.Inactive
            });
            if (RelayerProposals[_addr].action == RelayerActionType.Add) {
                // Add relayer
                Relayers[_addr] = true;
                TotalRelayers++;
            } else {
                // Remove relayer
                Relayers[_addr] = false;
                TotalRelayers--;
            }
        } else {
            RelayerProposals[_addr] = RelayerProposal({
                relayer: _addr,
                action: _action,
                numYes: 1, // Creator must vote in favour
                numNo: 0,
                status: VoteStatus.Active
            });
        }
        // Record vote
        RelayerProposals[_addr].votes[msg.sender] = true;
    }

    /**
     * Casts vote to add or remove a relayer, if the vote succeeds it will perform the action
     * @param _addr - Address of the relayer to be added or removed
     * @param _vote - Vote to either remove or add relayer
     */
    function voteRelayerProposal(address _addr, Vote _vote) public _isRelayer {
        require(RelayerProposals[_addr].status != VoteStatus.Inactive, "There is no active proposal!");
        require(!RelayerProposals[_addr].votes[msg.sender], "Relayer has already voted!");
        require(uint(_vote) <= 1, "Vote out of the vote enum range!");

        // Cast vote
        if (_vote == Vote.Yes) {
            RelayerProposals[_addr].numYes++;
        } else {
            RelayerProposals[_addr].numNo++;
        }

        // Record vote
        RelayerProposals[_addr].votes[msg.sender] = true;

        // Check if vote has met the threshold
        if (RelayerProposals[_addr].numYes >= RelayerThreshold ||
            TotalRelayers - RelayerProposals[_addr].numNo < RelayerThreshold) {

            // Vote succeeded, perform action
            if (RelayerProposals[_addr].numYes > RelayerProposals[_addr].numNo) {
                if (RelayerProposals[_addr].action == RelayerActionType.Add) {
                    // Add relayer
                    Relayers[_addr] = true;
                    TotalRelayers++;
                    RelayerProposals[_addr].status = VoteStatus.Inactive;
                } else {
                    // Remove relayer
                    Relayers[_addr] = false;
                    TotalRelayers--;
                    RelayerProposals[_addr].status = VoteStatus.Inactive;
                }
            }
        }
    }

    /**
     * A generic function that allows for the creation of any threshold change defined in ThresholdType.
     * @param _value - The proposed new value for the given threshold
     * @param _type - The threshold type that the proposal is for
     */
    function createThresholdProposal(uint _value, ThresholdType _type) public _isRelayer {
        uint key = uint(_type);

        require(_value <= TotalRelayers, "Total value must be lower than total Relayers!");
        require(ThresholdProposals[key].status != VoteStatus.Active, "A proposal is active!");

        ThresholdProposals[key] = ThresholdProposal({
            newValue: _value,
            proposalType: _type,
            numYes: 1,
            numNo: 0,
            status: VoteStatus.Active
        });
        // Add vote
        ThresholdProposals[key].votes[msg.sender] = true;
    }

    /**
     * Vote on a given threshold vote.
     * @param _vote - The vote that a relayer is making
     * @param _type - The type of threshold that is being proposed
     */
    function voteThresholdProposal(Vote _vote, ThresholdType _type) public _isRelayer {
        uint key = uint(_type);

        require(ThresholdProposals[key].status == VoteStatus.Active, "There is no active proposal!");
        require(!ThresholdProposals[key].votes[msg.sender], "Relayer has already voted!");

        if (_vote == Vote.Yes) {
            ThresholdProposals[key].numYes++;
        } else if (_vote == Vote.No) {
            ThresholdProposals[key].numNo++;
        }
        ThresholdProposals[key].votes[msg.sender] = true;

        if (ThresholdProposals[key].numYes >= RelayerThreshold ||
            TotalRelayers - ThresholdProposals[key].numNo < RelayerThreshold) {
            // Finalize the vote
            ThresholdProposals[key].status = VoteStatus.Finalized;

            if (ThresholdProposals[key].numYes > ThresholdProposals[key].numNo) {
                if (_type == ThresholdType.Relayer) {
                    RelayerThreshold = ThresholdProposals[key].newValue;
                } else if (_type == ThresholdType.Deposit) {
                    DepositThreshold = ThresholdProposals[key].newValue;
                }
            }
        }
    }

    /**
     * Gets the given deposit proposal
     * @param _chainId - The chain id from where the deposit originated from
     * @param _depositId - The deposit id that was generated on the origin chain
     */
    function getDepositProposal(uint _chainId, uint _depositId) public view returns (
        bytes32 hash,
        uint id,
        uint originChain,
        uint numYes,
        uint numNo,
        VoteStatus status
    ) {
        DepositProposal memory d = DepositProposals[_chainId][_depositId];
        return (
            d.hash,
            d.id,
            d.originChain,
            d.numYes,
            d.numNo,
            d.status
        );
    }
}
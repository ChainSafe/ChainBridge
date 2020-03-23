pragma solidity 0.6.4;

interface IRelayer {
    enum Vote {No, Yes}

    // RelayerActionType and _relayerActionTypeStrings must be kept
    // the same length and order to function properly
    enum RelayerActionType {Remove, Add}

    // VoteStatus and _voteStatusStrings must be kept
    // the same length and order to function properly
    enum VoteStatus {Inactive, Active}

    function isRelayer(address relayerAddress) external returns (bool);
    function getTotalRelayers() external returns (uint);
    function createRelayerProposal(address proposedAddress, RelayerActionType action) external;
    function voteRelayerProposal(address proposedAddress, Vote vote) external;
    function createRelayerThresholdProposal(uint proposedValue) external;
    function voteRelayerThresholdProposal(Vote vote) external;
}

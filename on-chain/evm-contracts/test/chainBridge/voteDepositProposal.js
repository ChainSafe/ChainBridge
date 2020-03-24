/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const TruffleAssert = require('truffle-assertions');
const Ethers = require('ethers');

const RelayerContract = artifacts.require("Relayer");
const BridgeContract = artifacts.require("Bridge");
const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ERC20HandlerContract = artifacts.require("ERC20Handler");

contract('Bridge - [voteDepositProposal with relayerThreshold > 1]', async (accounts) => {
    // const minterAndRelayer = accounts[0];
    const originChainRelayerAddress = accounts[1];
    const originChainRelayer2Address = accounts[4];
    const originChainRelayer3Address = accounts[5];
    const depositerAddress = accounts[2];
    const destinationChainRecipientAddress = accounts[3];
    const originChainID = 0;
    const destinationChainID = 0;
    const originChainInitialTokenAmount = 100;
    const depositAmount = 10;
    const expectedDepositNonce = 1;
    const relayerThreshold = 2;

    let RelayerInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
        RelayerInstance = await RelayerContract.new([
            originChainRelayerAddress,
            originChainRelayer2Address,
            originChainRelayer3Address], relayerThreshold);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, relayerThreshold);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        data = '0x' +
            Ethers.utils.hexZeroPad(DestinationERC20MintableInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(DestinationERC20HandlerInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(Ethers.utils.hexlify(depositAmount), 32).substr(2);
        dataHash = Ethers.utils.keccak256(data);

        await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress });
    });

    it('[sanity] depositProposal should be created with expected values', async () => {
        const expectedDepositProposal = {
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 1 // active
        };

        const depositProposal = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositNonce);
        for (const expectedProperty of Object.keys(expectedDepositProposal)) {
            // Testing all expected object properties
            assert.property(depositProposal, expectedProperty, `property: ${expectedProperty} does not exist in depositRecord`);

            // Testing all expected object values
            const depositProposalValue = depositProposal[expectedProperty].toNumber !== undefined ?
                depositProposal[expectedProperty].toNumber() : depositProposal[expectedProperty];
            assert.strictEqual(
                depositProposalValue, expectedDepositProposal[expectedProperty],
                `property: ${expectedProperty}'s value: ${depositProposalValue} does not match expected value: ${expectedDepositProposal[expectedProperty]}`)
        }
    });

    it('should vote on depositProposal successfully', async () => {
        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));
    });

    it('should revert because depositerAddress is not a relayer', async () => {
        await TruffleAssert.reverts(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            1,
            { from: depositerAddress }
        ));
    });

    it("depositProposal shouldn't be voted on if it has an Inactive status", async () => {
        await TruffleAssert.passes(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            42, // non existent depositID
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    xit("depositProposal shouldn't be voted on if it has a Passed status", async () => {
        
    });

    xit("depositProposal shouldn't be voted on if it has a Denied status", async () => {
        
    });

    xit("depositProposal shouldn't be voted on if it has a Transferred status", async () => {
        
    });

    it("relayer shouldn't be able to vote on a depositProposal more than once", async () => {
        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));

        await TruffleAssert.reverts(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));
    });

    it("relayer shouldn't be able to provide an invalid vote", async () => {
        await TruffleAssert.reverts(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            42, // invalid vote, out of range
            { from: originChainRelayer2Address }
        ));
    });

    it("Relayer's vote should be recorded correctly - yes vote", async () => {
        const depositProposalBeforeSecondVote = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositNonce);
        assert.strictEqual(depositProposalBeforeSecondVote._numYes.toNumber(), 1);
        assert.strictEqual(depositProposalBeforeSecondVote._numNo.toNumber(), 0);

        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));

        const depositProposalAfterSecondVote = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositNonce);
        assert.strictEqual(depositProposalAfterSecondVote._numYes.toNumber(), 2);
        assert.strictEqual(depositProposalAfterSecondVote._numNo.toNumber(), 0);
    });

    it("Relayer's vote should be recorded correctly - no vote", async () => {
        const depositProposalBeforeSecondVote = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositNonce);
        assert.strictEqual(depositProposalBeforeSecondVote._numYes.toNumber(), 1);
        assert.strictEqual(depositProposalBeforeSecondVote._numNo.toNumber(), 0);

        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            0, // vote against
            { from: originChainRelayer2Address }
        ));

        const depositProposalAfterSecondVote = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositNonce);
        assert.strictEqual(depositProposalAfterSecondVote._numYes.toNumber(), 1);
        assert.strictEqual(depositProposalAfterSecondVote._numNo.toNumber(), 1);
    });

    it("Relayer's address should be marked as voted for proposal", async () => {
        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));

        const hasVoted = await BridgeInstance.hasVoted(
            originChainID, OriginERC20HandlerInstance.address,
            expectedDepositNonce, originChainRelayer2Address);
        assert.isTrue(hasVoted);
    });

    it('Proposal status should be updated to passed after numYes >= relayerThreshold', async () => {
        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));

        const depositProposal = await BridgeInstance._depositProposals(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositNonce);
        assert.strictEqual(depositProposal._status.toNumber(), 3);
    });

    it('DepositProposalFinalized event should be emitted when proposal status updated to passed after numYes >= relayerThreshold', async () => {
        const voteTx = await BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            1, // vote in favor
            { from: originChainRelayer2Address }
        );

        TruffleAssert.eventEmitted(voteTx, 'DepositProposalFinalized', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.depositNonce.toNumber() === expectedDepositNonce
        });
    });

    it('Proposal status should be updated to denied if majority of relayers vote no', async () => {
        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            0, // vote in favor
            { from: originChainRelayer2Address }
        ));

        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            0, // vote in favor
            { from: originChainRelayer3Address }
        ));

        const depositProposal = await BridgeInstance._depositProposals(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositNonce);
        assert.strictEqual(depositProposal._status.toNumber(), 2);
    });

    it('DepositProposalFinalized event fired when proposal status updated to denied if majority of relayers vote no', async () => {
        await TruffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            0, // vote in favor
            { from: originChainRelayer2Address }
        ));

        const voteTx = await BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            0, // vote in favor
            { from: originChainRelayer3Address }
        );

        TruffleAssert.eventEmitted(voteTx, 'DepositProposalFinalized', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.depositNonce.toNumber() === expectedDepositNonce
        });
    });

    it('DepositProposalVote event fired when proposal vote made', async () => {
        const voteTx = await BridgeInstance.voteDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositNonce,
            0, // vote in favor
            { from: originChainRelayer2Address }
        );

        TruffleAssert.eventEmitted(voteTx, 'DepositProposalVote', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.depositNonce.toNumber() === expectedDepositNonce &&
                event.vote.toNumber() === 0 &&
                event.status.toNumber() === 1
        });
    });
});
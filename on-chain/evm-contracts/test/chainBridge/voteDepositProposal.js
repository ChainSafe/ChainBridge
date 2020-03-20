/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const truffleAssert = require('truffle-assertions');
const ethers = require('ethers');

const RelayerContract = artifacts.require("Relayer");
const BridgeContract = artifacts.require("Bridge");
const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ERC20HandlerContract = artifacts.require("ERC20Handler");

contract('Bridge - [voteDepositProposal with relayerThreshold > 1]', async (accounts) => {
    // const minterAndRelayer = accounts[0];
    const originChainRelayerAddress = accounts[1];
    const originChainRelayer2Address = accounts[4];
    const originChainRelayer3Address = accounts[5];
    const originChainDepositerAddress = accounts[2];
    const destinationChainRecipientAddress = accounts[3];
    const originChainID = 0;
    const destinationChainID = 0;
    const originChainInitialTokenAmount = 100;
    const depositAmount = 10;
    const expectedDepositID = 1;

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
            originChainRelayer3Address], 1);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, 2);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        await OriginERC20MintableInstance.mint(originChainDepositerAddress, originChainInitialTokenAmount);
        await OriginERC20MintableInstance.approve(OriginERC20HandlerInstance.address, depositAmount, { from: originChainDepositerAddress });
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        data = '0x' + ethers.utils.hexZeroPad(DestinationERC20MintableInstance.address, 32).substr(2) +
            ethers.utils.hexZeroPad(DestinationERC20HandlerInstance.address, 32).substr(2) +
            ethers.utils.hexZeroPad(ethers.utils.hexlify(depositAmount), 32).substr(2);
        dataHash = ethers.utils.keccak256(data);

        await BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress });
    });

    it('[sanity] ERC20 deposit record is created with expected depositID and values', async () => {
        const expectedDepositRecord = {
            _originChainTokenAddress: OriginERC20MintableInstance.address,
            _originChainHandlerAddress: OriginERC20HandlerInstance.address,
            _destinationChainID: destinationChainID,
            _destinationChainHandlerAddress: DestinationERC20HandlerInstance.address,
            _destinationRecipientAddress: destinationChainRecipientAddress,
            _amount: depositAmount
        };

        const depositRecord = await BridgeInstance._erc20DepositRecords.call(destinationChainID, expectedDepositID);
        for (const expectedProperty of Object.keys(expectedDepositRecord)) {
            // Testing all expected object properties
            assert.property(depositRecord, expectedProperty, `property: ${expectedProperty} does not exist in depositRecord`);

            // Testing all expected object values
            const depositRecordValue = depositRecord[expectedProperty].toNumber !== undefined ?
                depositRecord[expectedProperty].toNumber() : depositRecord[expectedProperty];
            assert.strictEqual(
                depositRecordValue, expectedDepositRecord[expectedProperty],
                `property: ${expectedProperty}'s value: ${depositRecordValue} does not match expected value: ${expectedDepositRecord[expectedProperty]}`)
        }
    });

    it('[sanity] depositProposal is created with expected values', async () => {
        const expectedDepositProposal = {
            _originChainID: originChainID,
            _depositID: expectedDepositID,
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 1 // passed
        };

        const depositProposal = await BridgeInstance._depositProposals.call(originChainID, expectedDepositID);
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

    it('depositProposal can be voted on', async () => {
        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));
    });

    it('Only relayers should be able to vote on a deposit proposal', async () => {
        await truffleAssert.reverts(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            1, // vote in favor
            { from: originChainDepositerAddress }
        ));
    });

    it('Can only vote on a proposal that has active status', async () => {
        await truffleAssert.reverts(BridgeInstance.voteDepositProposal(
            originChainID,
            400, // fake depositID
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));
    });

    it("Can only vote on a proposal if relayer hasn't already voted for it", async () => {
        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));

        await truffleAssert.reverts(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));
    });

    it('Relayer must provide a valid vote', async () => {
        await truffleAssert.fails(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            42, // invalid vote, out of range
            { from: originChainRelayer2Address }
        ), truffleAssert.ErrorType.INVALID_OPCODE);
    });

    it("Relayer's vote should be recorded correctly - yes vote", async () => {
        const depositProposalBeforeSecondVote = await BridgeInstance._depositProposals.call(originChainID, expectedDepositID);
        assert.strictEqual(depositProposalBeforeSecondVote._numYes.toNumber(), 1);
        assert.strictEqual(depositProposalBeforeSecondVote._numNo.toNumber(), 0);

        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));

        const depositProposalAfterSecondVote = await BridgeInstance._depositProposals.call(originChainID, expectedDepositID);
        assert.strictEqual(depositProposalAfterSecondVote._numYes.toNumber(), 2);
        assert.strictEqual(depositProposalAfterSecondVote._numNo.toNumber(), 0);
    });

    it("Relayer's vote should be recorded correctly - no vote", async () => {
        const depositProposalBeforeSecondVote = await BridgeInstance._depositProposals.call(originChainID, expectedDepositID);
        assert.strictEqual(depositProposalBeforeSecondVote._numYes.toNumber(), 1);
        assert.strictEqual(depositProposalBeforeSecondVote._numNo.toNumber(), 0);

        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            0, // vote against
            { from: originChainRelayer2Address }
        ));

        const depositProposalAfterSecondVote = await BridgeInstance._depositProposals.call(originChainID, expectedDepositID);
        assert.strictEqual(depositProposalAfterSecondVote._numYes.toNumber(), 1);
        assert.strictEqual(depositProposalAfterSecondVote._numNo.toNumber(), 1);
    });

    it("Relayer's address should be marked as voted for proposal", async () => {
        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));

        const hasVoted = await BridgeInstance.hasVoted(originChainID, expectedDepositID, originChainRelayer2Address);
        assert.isTrue(hasVoted);
    });

    it('Proposal status should be updated to passed after numYes >= relayerThreshold', async () => {
        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            1, // vote in favor
            { from: originChainRelayer2Address }
        ));

        const depositProposal = await BridgeInstance._depositProposals(originChainID, expectedDepositID);
        assert.strictEqual(depositProposal._status.toNumber(), 3);
    });

    it('DepositProposalFinalized event fired when proposal status updated to passed after numYes >= relayerThreshold', async () => {
        const voteTx = await BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            1, // vote in favor
            { from: originChainRelayer2Address }
        );

        truffleAssert.eventEmitted(voteTx, 'DepositProposalFinalized', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.depositID.toNumber() === expectedDepositID
        });
    });

    it('Proposal status should be updated to denied if majority of relayers vote no', async () => {
        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            0, // vote in favor
            { from: originChainRelayer2Address }
        ));

        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            0, // vote in favor
            { from: originChainRelayer3Address }
        ));

        const depositProposal = await BridgeInstance._depositProposals(originChainID, expectedDepositID);
        assert.strictEqual(depositProposal._status.toNumber(), 2);
    });

    it('DepositProposalFinalized event fired when proposal status updated to denied if majority of relayers vote no', async () => {
        await truffleAssert.passes(BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            0, // vote in favor
            { from: originChainRelayer2Address }
        ));

        const voteTx = await BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            0, // vote in favor
            { from: originChainRelayer3Address }
        );

        truffleAssert.eventEmitted(voteTx, 'DepositProposalFinalized', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.depositID.toNumber() === expectedDepositID
        });
    });

    it('DepositProposalVote event fired when proposal vote made', async () => {
        const voteTx = await BridgeInstance.voteDepositProposal(
            originChainID,
            expectedDepositID,
            0, // vote in favor
            { from: originChainRelayer2Address }
        );

        truffleAssert.eventEmitted(voteTx, 'DepositProposalVote', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.depositID.toNumber() === expectedDepositID &&
                event.vote.toNumber() === 0 &&
                event.status.toNumber() === 1
        });
    });
});
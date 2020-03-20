/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ReceiverContract = artifacts.require("Receiver");
const { Vote, VoteStatus, ThresholdType } = require("../helpers");

contract('Receiver - [Thresholds::Create]', async (accounts) => {
    let ReceiverInstance;

    // Set relayers
    let v1 = accounts[0];
    let v2 = accounts[1];
    let v3 = accounts[3];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2, v3], // bridge relayers
            2,            // depoist threshold
            2             // relayer threshold
        )
    });

    it('create deposit threshold proposal', async () => {
        const { receipt } = await ReceiverInstance.createThresholdProposal(3, ThresholdType.Deposit);
        assert.strictEqual(receipt.status, true);
    });

    it('create relayer threshold proposal', async () => {
        const { receipt } = await ReceiverInstance.createThresholdProposal(3, ThresholdType.Relayer);
        assert.strictEqual(receipt.status, true);
    });
});

contract('Receiver - [Thresholds::Voting]', async (accounts) => {
    let ReceiverInstance;

    // Set relayers
    let v1 = accounts[0];
    let v2 = accounts[1];
    let v3 = accounts[3];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2, v3], // bridge relayers
            2,            // depoist threshold
            2             // relayer threshold
        )
    });

    it('can increment deposit threshold', async () => {
        const before = await ReceiverInstance.DepositThreshold();
        assert.strictEqual(before.toNumber(), 2);
        // Create & vote
        // This should pass because we have 3 relayers
        await ReceiverInstance.createThresholdProposal(3, ThresholdType.Deposit, { from: v1 });
        await ReceiverInstance.voteThresholdProposal(Vote.Yes, ThresholdType.Deposit, { from: v2 });
        // Check success
        const vote = await ReceiverInstance.ThresholdProposals(ThresholdType.Deposit);
        assert.strictEqual(vote.status.toNumber(), VoteStatus.Finalized);
        const after = await ReceiverInstance.DepositThreshold();
        assert.strictEqual(after.toNumber(), 3);
    });

    it('can decrement deposit threshold', async () => {
        const before = await ReceiverInstance.DepositThreshold();
        assert.strictEqual(before.toNumber(), 2);
        // Create & vote
        // This should pass because we have 3 relayers
        await ReceiverInstance.createThresholdProposal(1, ThresholdType.Deposit, { from: v1 });
        await ReceiverInstance.voteThresholdProposal(Vote.Yes, ThresholdType.Deposit, { from: v2 });
        // Check success
        const vote = await ReceiverInstance.ThresholdProposals(ThresholdType.Deposit);
        assert.strictEqual(vote.status.toNumber(), VoteStatus.Finalized);
        const after = await ReceiverInstance.DepositThreshold();
        assert.strictEqual(after.toNumber(), 1);
    });

    it('can increment relayer threshold', async () => {
        const before = await ReceiverInstance.RelayerThreshold();
        assert.strictEqual(before.toNumber(), 2);
        // Create & vote
        // This should pass because we have 3 relayers
        await ReceiverInstance.createThresholdProposal(3, ThresholdType.Relayer, { from: v1 });
        await ReceiverInstance.voteThresholdProposal(Vote.Yes, ThresholdType.Relayer, { from: v2 });
        // Check success
        const vote = await ReceiverInstance.ThresholdProposals(ThresholdType.Relayer);
        assert.strictEqual(vote.status.toNumber(), VoteStatus.Finalized);
        const after = await ReceiverInstance.RelayerThreshold();
        assert.strictEqual(after.toNumber(), 3);
    });

    it('can decrement relayer threshold', async () => {
        const before = await ReceiverInstance.RelayerThreshold();
        assert.strictEqual(before.toNumber(), 2);
        // Create & vote
        // This should pass because we have 3 relayers
        await ReceiverInstance.createThresholdProposal(1, ThresholdType.Relayer, { from: v1 });
        await ReceiverInstance.voteThresholdProposal(Vote.Yes, ThresholdType.Relayer, { from: v2 });
        // Check success
        const vote = await ReceiverInstance.ThresholdProposals(ThresholdType.Relayer);
        assert.strictEqual(vote.status.toNumber(), VoteStatus.Finalized);
        const after = await ReceiverInstance.RelayerThreshold();
        assert.strictEqual(after.toNumber(), 1);
    });
});
const ReceiverContract = artifacts.require("Receiver");
const { Vote, VoteStatus, ThresholdType } = require("../helpers");

contract('Receiver - [Thresholds::Create]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];
    let v3 = accounts[3];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2, v3], // bridge validators
            2,            // depoist threshold
            2             // validator threshold
        )
    });

    it('create deposit threshold proposal', async () => {
        const { receipt } = await ReceiverInstance.createThresholdProposal(3, ThresholdType.Deposit);
        assert.strictEqual(receipt.status, true);
    });

    it('create validator threshold proposal', async () => {
        const { receipt } = await ReceiverInstance.createThresholdProposal(3, ThresholdType.Validator);
        assert.strictEqual(receipt.status, true);
    });
});

contract('Receiver - [Thresholds::Voting]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];
    let v3 = accounts[3];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2, v3], // bridge validators
            2,            // depoist threshold
            2             // validator threshold
        )
    });

    it('can increment deposit threshold', async () => {
        const before = await ReceiverInstance.DepositThreshold();
        assert.strictEqual(before.toNumber(), 2);
        // Create & vote
        // This should pass because we have 3 validators
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
        // This should pass because we have 3 validators
        await ReceiverInstance.createThresholdProposal(1, ThresholdType.Deposit, { from: v1 });
        await ReceiverInstance.voteThresholdProposal(Vote.Yes, ThresholdType.Deposit, { from: v2 });
        // Check success
        const vote = await ReceiverInstance.ThresholdProposals(ThresholdType.Deposit);
        assert.strictEqual(vote.status.toNumber(), VoteStatus.Finalized);
        const after = await ReceiverInstance.DepositThreshold();
        assert.strictEqual(after.toNumber(), 1);
    });

    it('can increment validator threshold', async () => {
        const before = await ReceiverInstance.ValidatorThreshold();
        assert.strictEqual(before.toNumber(), 2);
        // Create & vote
        // This should pass because we have 3 validators
        await ReceiverInstance.createThresholdProposal(3, ThresholdType.Validator, { from: v1 });
        await ReceiverInstance.voteThresholdProposal(Vote.Yes, ThresholdType.Validator, { from: v2 });
        // Check success
        const vote = await ReceiverInstance.ThresholdProposals(ThresholdType.Validator);
        assert.strictEqual(vote.status.toNumber(), VoteStatus.Finalized);
        const after = await ReceiverInstance.ValidatorThreshold();
        assert.strictEqual(after.toNumber(), 3);
    });

    it('can decrement validator threshold', async () => {
        const before = await ReceiverInstance.ValidatorThreshold();
        assert.strictEqual(before.toNumber(), 2);
        // Create & vote
        // This should pass because we have 3 validators
        await ReceiverInstance.createThresholdProposal(1, ThresholdType.Validator, { from: v1 });
        await ReceiverInstance.voteThresholdProposal(Vote.Yes, ThresholdType.Validator, { from: v2 });
        // Check success
        const vote = await ReceiverInstance.ThresholdProposals(ThresholdType.Validator);
        assert.strictEqual(vote.status.toNumber(), VoteStatus.Finalized);
        const after = await ReceiverInstance.ValidatorThreshold();
        assert.strictEqual(after.toNumber(), 1);
    });
});
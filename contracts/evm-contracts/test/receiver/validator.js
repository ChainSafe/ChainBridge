const ReceiverContract = artifacts.require("Receiver");
const { ValidatorActionType, Vote, VoteStatus } = require("../helpers");

contract('Receiver - [Validator::Create::Basic]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('validators can create proposals', async () => {
        let { receipt } = await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        assert.strictEqual(receipt.status, true);
    });

    it('only validators can create proposals', async () => {
        try {
            let { receipt } = await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: accounts[4] });
            // This case shouldn't be hit, fail safe.
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Sender is not a validator");
        }
    });

    it('cannot create proposal if it exists', async () => {
        try {
            await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            let { receipt } = await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            // This case shouldn't be hit, fail safe.
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is already an active proposal!");
        }
    })

    it('cannnot add an exisitng validator', async () => {
        try {
            const { receipt } = await ReceiverInstance.createValidatorProposal(v2, ValidatorActionType.Add, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Validator is already active!");
        }
    });

    it('cannot remove a non validator', async () => {
        try {
            const { receipt } = await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Remove, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Validator is not active!");
        }
    });

    it('cannot create a proposal with inccorect action type', async () => {
        try {
            const { receipt } = await ReceiverInstance.createValidatorProposal(accounts[2], 3, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "invalid opcode");
        }
    });
});

contract('Receiver - [Validator::Create::low-threshold]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            1         // validator threshold
        )
    });

    it('auto finalize for one threshold', async () => {
        await ReceiverInstance.createValidatorProposal(accounts[4], ValidatorActionType.Add, { from: v1 });
        const after = await ReceiverInstance.Validators(accounts[4]);
        assert.strictEqual(after, true);
    });

});

contract('Receiver - [Validator::Create::Advanced]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('can create proposal if proposal has ended', async () => {
        let proposal;
        // check before
        proposal = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Inactive, "intial failed");

        // Vote in
        await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        // Check during vote
        proposal = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Active, "Vote in failed");
        await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });
        // Check after vote
        proposal = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Inactive, "Post vote failed");

        // Vote out
        await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Remove, { from: v1 });
        proposal = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Active, "Second vote failed");
    });
});

contract('Receiver - [Validator::Voting::Basic]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('can vote on a proposal', async () => {
        await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        let before = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(before.numYes.toNumber() + before.numNo.toNumber(), 1);

        await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });
        let after = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(after.numYes.toNumber() + after.numNo.toNumber(), 2);
    });

    it('cannot vote twice on a proposal', async () => {
        try {
            await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            let { receipt } = await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Validator has already voted!");
        }
    });

    it('only validators can vote on proposals', async () => {
        await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        let before = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(before.numYes.toNumber() + before.numNo.toNumber(), 1);

        try {
            await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: accounts[4] });
            // We shouldn't get here
            let after = await ReceiverInstance.ValidatorProposals(accounts[2]);
            assert.strictEqual(after.numYes.toNumber() + after.numNo.toNumber(), 1);
        } catch (e) {
            assert.include(e.message, "Sender is not a validator");
        }
    });

    it('cannot vote if proposal doesnt exists', async () => {
        try {
            let { receipt } = await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is no active proposal!");
        }
    });

    it('cannot vote incorrect value', async () => {
        try {
            await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            let { receipt } = await ReceiverInstance.voteValidatorProposal(accounts[2], 2, { from: v2 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "invalid opcode");
        }
    });
});

contract('Receiver - [Validator::Voting::Advanced]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('can vote in a new validator', async () => {
        const before = await ReceiverInstance.Validators(accounts[2]);
        assert.strictEqual(before, false);

        await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });

        const after = await ReceiverInstance.Validators(accounts[2]);
        assert.strictEqual(after, true);
    });

    it('can vote out a validator', async () => {
        const beforeStatus = await ReceiverInstance.Validators(v1);
        assert.strictEqual(beforeStatus, true);

        // Add third validator
        await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });

        // Vote out validator
        await ReceiverInstance.createValidatorProposal(v1, ValidatorActionType.Remove, { from: accounts[2] });
        await ReceiverInstance.voteValidatorProposal(v1, Vote.Yes, { from: v2 });

        const afterStatus = await ReceiverInstance.Validators(v1);
        assert.strictEqual(afterStatus, false);
    });

    it('cannot vote if proposal ended', async () => {
        // Add third validator
        await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });

        try {
            let { receipt } = await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: accounts[2] });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is no active proposal!");
        }
    })
});
const HomeContract = artifacts.require("Home");
const { ValidatorActionType, Vote, VoteStatus } = require("./helpers");

contract('Receiver - deployment', async (accounts) => {
    let HomeInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    before(async () => {
        HomeInstance = await HomeContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('should set constructor values', async () => {
        let validators = await HomeInstance.TotalValidators.call();
        assert.strictEqual(parseInt(validators, 10), 2);

        let depositThreshold = await HomeInstance.voteDepositThreshold.call();
        assert.strictEqual(parseInt(depositThreshold, 10), 2);

        let validatorThreshold = await HomeInstance.voteValidatorThreshold.call();
        assert.strictEqual(parseInt(validatorThreshold, 10), 2);

        let validator1 = await HomeInstance.Validators(v1);
        assert.strictEqual(validator1, true);

        let validator2 = await HomeInstance.Validators(v2);
        assert.strictEqual(validator2, true);
    });
});

contract('Receiver - [Create] validator proposals', async (accounts) => {
    let HomeInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        HomeInstance = await HomeContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('validators can create proposals', async () => {
        let {receipt} = await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, {from: v1});
        assert.strictEqual(receipt.status, true);
    });

    it('only validators can create proposals', async () => {
        try {
            let { receipt } = await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: accounts[4] });
            // This case shouldn't be hit, fail safe.
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Sender is not a validator");
        }
    });

    it('cannot create proposal if it exists', async () => {
        try {
            await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            let { receipt } = await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            // This case shouldn't be hit, fail safe.
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is already an active proposal!");
        }
    })

    it('cannnot add an exisitng validator', async () => {
        try {
            const { receipt } = await HomeInstance.createValidatorProposal(v2, ValidatorActionType.Add, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Validator is already active!");
        }
    });

    it('cannot remove a non validator', async () => {
        try {
            const { receipt } = await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Remove, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Validator is not active!");
        }
    });

    it('cannot create a proposal with inccorect action type', async () => {
        try {
            const { receipt } = await HomeInstance.createValidatorProposal(accounts[2], 3, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "invalid opcode");
        }
    });

    it('can create proposal if proposal has ended', async () => {
        let proposal;
        // check before
        proposal = await HomeInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Inactive, "intial failed");

        // Vote in
        await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        // Check during vote
        proposal = await HomeInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Active, "Vote in failed");
        await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });
        // Check after vote
        proposal = await HomeInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Inactive, "Post vote failed");

        // Vote out
        await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Remove, { from: v1 });
        proposal = await HomeInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Active, "Second vote failed");
    });
});

contract('Receiver - [Voting] validator proposals', async (accounts) => {
    let HomeInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        HomeInstance = await HomeContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });
    
    it('can vote on a proposal', async () => {
        await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        let before = await HomeInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(before.totalVotes.toNumber(), 1);

        await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, {from: v2});
        let after = await HomeInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(after.totalVotes.toNumber(), 2);
    });

    it('cannot vote twice on a proposal', async () => {
        try {
            await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            let {receipt} = await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Validator has already voted!");
        }
    });

    it('only validators can vote on proposals', async () => {
        await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        let before = await HomeInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(before.totalVotes.toNumber(), 1);

        try {
            await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: accounts[4] });
            // We shouldn't get here
            let after = await HomeInstance.ValidatorProposals(accounts[2]);
            assert.strictEqual(after.totalVotes.toNumber(), 1);
        } catch (e) {
            assert.include(e.message, "Sender is not a validator");
        }
    });

    it('cannot vote if proposal doesnt exists', async () => {
        try {
            let { receipt } = await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is no active proposal!");
        }
    });

    it('cannot vote incorrect value', async () => {
        try {
            await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            let {receipt} = await HomeInstance.voteValidatorProposal(accounts[2], 2, { from: v2 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "invalid opcode");
        }
    });

    it('can vote in a new validator', async () => {
        const before = await HomeInstance.Validators(accounts[2]);
        assert.strictEqual(before, false);
        
        await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });
        
        const after = await HomeInstance.Validators(accounts[2]);
        assert.strictEqual(after, true);
    });

    it('can vote out a validator', async () => {
        const beforeStatus = await HomeInstance.Validators(v1);
        assert.strictEqual(beforeStatus, true);
        
        // Add third validator
        await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });

        // Vote out validator
        await HomeInstance.createValidatorProposal(v1, ValidatorActionType.Remove, { from: accounts[2] });
        await HomeInstance.voteValidatorProposal(v1, Vote.Yes, { from: v2 });

        const afterStatus = await HomeInstance.Validators(v1);
        assert.strictEqual(afterStatus, false);
    });

    it('cannot vote if proposal ended', async () => {
        // Add third validator
        await HomeInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v2 });

        try {
            let {receipt} = await HomeInstance.voteValidatorProposal(accounts[2], Vote.Yes, {from: accounts[2]});
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is no active proposal!");
        }
    })
});
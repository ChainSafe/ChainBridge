const ReceiverContract = artifacts.require("Receiver");
const { ValidatorActionType, Vote, VoteStatus, CreateDepositData } = require("./helpers");

contract('Receiver - deployment', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    before(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('should set constructor values', async () => {
        let validators = await ReceiverInstance.TotalValidators.call();
        assert.strictEqual(parseInt(validators, 10), 2);

        let depositThreshold = await ReceiverInstance.voteDepositThreshold.call();
        assert.strictEqual(parseInt(depositThreshold, 10), 2);

        let validatorThreshold = await ReceiverInstance.voteValidatorThreshold.call();
        assert.strictEqual(parseInt(validatorThreshold, 10), 2);

        let validator1 = await ReceiverInstance.Validators(v1);
        assert.strictEqual(validator1, true);

        let validator2 = await ReceiverInstance.Validators(v2);
        assert.strictEqual(validator2, true);
    });
});

contract('Receiver - [Create::Basic] validator proposals', async (accounts) => {
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
        let {receipt} = await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, {from: v1});
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

contract('Receiver - [Create::Advanced] validator proposals', async (accounts) => {
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

contract('Receiver - [Voting::Basic] validator proposals', async (accounts) => {
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
        assert.strictEqual(before.totalVotes.toNumber(), 1);

        await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, {from: v2});
        let after = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(after.totalVotes.toNumber(), 2);
    });

    it('cannot vote twice on a proposal', async () => {
        try {
            await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
            let {receipt} = await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Validator has already voted!");
        }
    });

    it('only validators can vote on proposals', async () => {
        await ReceiverInstance.createValidatorProposal(accounts[2], ValidatorActionType.Add, { from: v1 });
        let before = await ReceiverInstance.ValidatorProposals(accounts[2]);
        assert.strictEqual(before.totalVotes.toNumber(), 1);

        try {
            await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, { from: accounts[4] });
            // We shouldn't get here
            let after = await ReceiverInstance.ValidatorProposals(accounts[2]);
            assert.strictEqual(after.totalVotes.toNumber(), 1);
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
            let {receipt} = await ReceiverInstance.voteValidatorProposal(accounts[2], 2, { from: v2 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "invalid opcode");
        }
    });
});

contract('Receiver - [Voting::Advanced] validator proposals', async (accounts) => {
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
            let {receipt} = await ReceiverInstance.voteValidatorProposal(accounts[2], Vote.Yes, {from: accounts[2]});
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is no active proposal!");
        }
    })
});

contract('Receiver - [Create::Basic] deposit proposals', async (accounts) => {
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
        let { receipt } = await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        assert.strictEqual(receipt.status, true);
    });

    it('only validators can create proposals', async () => {
        try {
            let { receipt } = await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: accounts[4] });
            // This case shouldn't be hit, fail safe.
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Sender is not a validator");
        }
    });

    it('cannot create proposal if it exists', async () => {
        try {
            await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
            let { receipt } = await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
            // This case shouldn't be hit, fail safe.
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "A proposal already exists!");
        }
    })
});

contract('Receiver - [Create::Basic] deposit proposals', async (accounts) => {
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

    it('cannot create proposal after finalized', async () => {
        await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: v2 });
        try {
            await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });    
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "A proposal already exists!");
        }
    });
});

contract('Receiver - [Voting::Basic] deposit proposals', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];
    let v3 = accounts[3];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2, v3], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it("can vote on proposal", async () => {
        await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        let before = await ReceiverInstance.getDepositProposal.call(0, 0);
        assert.strictEqual(before.numYes.toNumber(), 1);

        let {receipt} = await ReceiverInstance.voteDepositProposal(0,0, Vote.Yes, {from: v2});
        assert.strictEqual(receipt.status, true);

        let after = await ReceiverInstance.getDepositProposal.call(0, 0);
        assert.strictEqual(after.numYes.toNumber(), 2);
    });

    it('cannot vote twice on a proposal', async () => {
        try {
            await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
            let { receipt } = await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Validator has already voted!");
        }
    });

    it('only validators can vote on proposals', async () => {
        await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        let before = await ReceiverInstance.getDepositProposal.call(0, 0);
        assert.strictEqual(before.numYes.toNumber(), 1);

        try {
            await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: accounts[4] });
            // We shouldn't get here
            let after = await ReceiverInstance.getDepositProposal(0, 0);
            assert.strictEqual(after.numYes.toNumber(), 1);
        } catch (e) {
            assert.include(e.message, "Sender is not a validator");
        }
    });

    it('cannot vote if proposal doesnt exists', async () => {
        try {
            let { receipt } = await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is no active proposal!");
        }
    });

    it('cannot vote incorrect value', async () => {
        try {
            await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
            // 5 is invalid, should be Vote.Yes (0) or Vote.No (2)
            let { receipt } = await ReceiverInstance.voteDepositProposal(0, 0, 5, { from: v2 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "invalid opcode");
        }
    });

    it('cannot vote if proposal finalized', async () => {
        try {
            await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
            await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: v2 });
            let { receipt } = await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: v3 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Proposal has already been finalized!");
        }
    });
});

contract('Receiver - [Voting::Advanced] deposit proposals', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];
    let v3 = accounts[3];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2, v3], // bridge validators
            2,        // depoist threshold
            2         // validator threshold
        )
    });

    it('can vote in a deposit', async () => {
        await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: v2 });
        let res = await ReceiverInstance.getDepositProposal.call(0, 0);
        assert.strictEqual(res.status.toNumber(), VoteStatus.Finalized);
    });
});
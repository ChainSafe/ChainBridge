const ReceiverContract = artifacts.require("Receiver");
const { Vote, VoteStatus, CreateDepositData } = require("../helpers");

contract('Receiver - [Deposit::Create::Basic]', async (accounts) => {
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


contract('Receiver - [Deposit::Create::low-threshold]', async (accounts) => {
    let ReceiverInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge validators
            1,        // depoist threshold
            2         // validator threshold
        )
    });

    it('auto finalize for one threshold', async () => {
        await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        let res = await ReceiverInstance.getDepositProposal.call(0, 0);
        assert.strictEqual(res.status.toNumber(), VoteStatus.Finalized);
    });
});

contract('Receiver - [Deposit::Create::Advanced]', async (accounts) => {
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

contract('Receiver - [Deposit::Voting::Basic]', async (accounts) => {
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

    it("can vote on proposal", async () => {
        await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        let before = await ReceiverInstance.getDepositProposal.call(0, 0);
        assert.strictEqual(before.numYes.toNumber(), 1);

        let { receipt } = await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: v2 });
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

    it('vote result success', async () => {
        await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        await ReceiverInstance.voteDepositProposal(0, 0, Vote.Yes, { from: v2 });
        let res = await ReceiverInstance.getDepositProposal.call(0, 0);
        assert.strictEqual(res.status.toNumber(), VoteStatus.Finalized);
        assert(res.numYes.toNumber() > res.numNo.toNumber(), "No was > Yes");
    });

    it('vote result fails', async () => {
        await ReceiverInstance.createDepositProposal(...CreateDepositData(), { from: v1 });
        await ReceiverInstance.voteDepositProposal(0, 0, Vote.No, { from: v2 });
        await ReceiverInstance.voteDepositProposal(0, 0, Vote.No, { from: v3 });

        let res = await ReceiverInstance.getDepositProposal.call(0, 0);
        console.log({res})
        assert.strictEqual(res.status.toNumber(), VoteStatus.Finalized);
        assert(res.numYes.toNumber() <= res.numNo.toNumber(), "Yes was > No");
    });
});

contract('Receiver - [Deposit::Voting::Advanced]', async (accounts) => {
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

    /**
     * Possible tests
     * - vote failed, tries to submit deposit anyway
     */
});
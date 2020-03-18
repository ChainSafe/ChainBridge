/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ReceiverContract = artifacts.require("Receiver");
const { RelayerActionType, Vote, VoteStatus } = require("../helpers");

contract('Receiver - [Relayer::Create::Basic]', async (accounts) => {
    let ReceiverInstance;

    // Set relayers
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge relayers
            2,        // depoist threshold
            2         // relayer threshold
        )
    });

    it('relayers can create proposals', async () => {
        let { receipt } = await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
        assert.strictEqual(receipt.status, true);
    });

    it('only relayers can create proposals', async () => {
        try {
            let { receipt } = await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: accounts[4] });
            // This case shouldn't be hit, fail safe.
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Sender is not a relayer");
        }
    });

    it('cannot create proposal if it exists', async () => {
        try {
            await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
            let { receipt } = await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
            // This case shouldn't be hit, fail safe.
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is already an active proposal!");
        }
    })

    it('cannnot add an exisitng relayer', async () => {
        try {
            const { receipt } = await ReceiverInstance.createRelayerProposal(v2, RelayerActionType.Add, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Relayer is already active!");
        }
    });

    it('cannot remove a non relayer', async () => {
        try {
            const { receipt } = await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Remove, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Relayer is not active!");
        }
    });

    it('cannot create a proposal with inccorect action type', async () => {
        try {
            const { receipt } = await ReceiverInstance.createRelayerProposal(accounts[2], 3, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "invalid opcode");
        }
    });
});

contract('Receiver - [Relayer::Create::low-threshold]', async (accounts) => {
    let ReceiverInstance;

    // Set relayers
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge relayers
            2,        // depoist threshold
            1         // relayer threshold
        )
    });

    it('auto finalize for one threshold', async () => {
        await ReceiverInstance.createRelayerProposal(accounts[4], RelayerActionType.Add, { from: v1 });
        const after = await ReceiverInstance.Relayers(accounts[4]);
        assert.strictEqual(after, true);
    });

});

contract('Receiver - [Relayer::Create::Advanced]', async (accounts) => {
    let ReceiverInstance;

    // Set relayers
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge relayers
            2,        // depoist threshold
            2         // relayer threshold
        )
    });

    it('can create proposal if proposal has ended', async () => {
        let proposal;
        // check before
        proposal = await ReceiverInstance.RelayerProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Inactive, "intial failed");

        // Vote in
        await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
        // Check during vote
        proposal = await ReceiverInstance.RelayerProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Active, "Vote in failed");
        await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: v2 });
        // Check after vote
        proposal = await ReceiverInstance.RelayerProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Inactive, "Post vote failed");

        // Vote out
        await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Remove, { from: v1 });
        proposal = await ReceiverInstance.RelayerProposals(accounts[2]);
        assert.strictEqual(proposal.status.toNumber(), VoteStatus.Active, "Second vote failed");
    });
});

contract('Receiver - [Relayer::Voting::Basic]', async (accounts) => {
    let ReceiverInstance;

    // Set relayers
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge relayers
            2,        // depoist threshold
            2         // relayer threshold
        )
    });

    it('can vote on a proposal', async () => {
        await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
        let before = await ReceiverInstance.RelayerProposals(accounts[2]);
        assert.strictEqual(before.numYes.toNumber() + before.numNo.toNumber(), 1);

        await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: v2 });
        let after = await ReceiverInstance.RelayerProposals(accounts[2]);
        assert.strictEqual(after.numYes.toNumber() + after.numNo.toNumber(), 2);
    });

    it('cannot vote twice on a proposal', async () => {
        try {
            await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
            let { receipt } = await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "Relayer has already voted!");
        }
    });

    it('only relayers can vote on proposals', async () => {
        await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
        let before = await ReceiverInstance.RelayerProposals(accounts[2]);
        assert.strictEqual(before.numYes.toNumber() + before.numNo.toNumber(), 1);

        try {
            await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: accounts[4] });
            // We shouldn't get here
            let after = await ReceiverInstance.RelayerProposals(accounts[2]);
            assert.strictEqual(after.numYes.toNumber() + after.numNo.toNumber(), 1);
        } catch (e) {
            assert.include(e.message, "Sender is not a relayer");
        }
    });

    it('cannot vote if proposal doesnt exists', async () => {
        try {
            let { receipt } = await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: v1 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is no active proposal!");
        }
    });

    it('cannot vote incorrect value', async () => {
        try {
            await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
            let { receipt } = await ReceiverInstance.voteRelayerProposal(accounts[2], 2, { from: v2 });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "invalid opcode");
        }
    });
});

contract('Receiver - [Relayer::Voting::Advanced]', async (accounts) => {
    let ReceiverInstance;

    // Set relayers
    let v1 = accounts[0];
    let v2 = accounts[1];

    beforeEach(async () => {
        ReceiverInstance = await ReceiverContract.new(
            [v1, v2], // bridge relayers
            2,        // depoist threshold
            2         // relayer threshold
        )
    });

    it('can vote in a new relayer', async () => {
        const before = await ReceiverInstance.Relayers(accounts[2]);
        assert.strictEqual(before, false);

        await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
        await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: v2 });

        const after = await ReceiverInstance.Relayers(accounts[2]);
        assert.strictEqual(after, true);
    });

    it('can vote out a relayer', async () => {
        const beforeStatus = await ReceiverInstance.Relayers(v1);
        assert.strictEqual(beforeStatus, true);

        // Add third relayer
        await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
        await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: v2 });

        // Vote out relayer
        await ReceiverInstance.createRelayerProposal(v1, RelayerActionType.Remove, { from: accounts[2] });
        await ReceiverInstance.voteRelayerProposal(v1, Vote.Yes, { from: v2 });

        const afterStatus = await ReceiverInstance.Relayers(v1);
        assert.strictEqual(afterStatus, false);
    });

    it('cannot vote if proposal ended', async () => {
        // Add third relayer
        await ReceiverInstance.createRelayerProposal(accounts[2], RelayerActionType.Add, { from: v1 });
        await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: v2 });

        try {
            let { receipt } = await ReceiverInstance.voteRelayerProposal(accounts[2], Vote.Yes, { from: accounts[2] });
            assert.strictEqual(receipt.status, false);
        } catch (e) {
            assert.include(e.message, "There is no active proposal!");
        }
    })
});
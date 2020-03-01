/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const EmitterContract = artifacts.require("Emitter");
const ERC721Contract = artifacts.require("ERC721Mintable");

contract('Emitter - [deposits::ERC20]', async (accounts) => {
    const minter = accounts[0];
    const receiver = accounts[1];
    const destinationChainID = 0;
    const tokenID = 1;

    let EmitterInstance;
    let ERC721Instance;

    beforeEach(async () => {
        EmitterInstance = await EmitterContract.new();
        ERC721Instance = await ERC721Contract.new();
        await ERC721Instance.mint(minter, tokenID);
        await ERC721Instance.approve(EmitterInstance.address, tokenID)
    });

    it('[sanity] test minter balance', async () => {
        const tokenOwner = await ERC721Instance.ownerOf(tokenID);
        assert.strictEqual(tokenOwner, minter);

        const minterbal = await ERC721Instance.balanceOf(minter);
        assert.strictEqual(minterbal.toNumber(), 1);

        const contractbal = await ERC721Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(contractbal.toNumber(), 0);

        const approvedAddress = await ERC721Instance.getApproved(tokenID);
        assert.strictEqual(approvedAddress, EmitterInstance.address);
    });

    it('deposit', async () => {
        EmitterInstance.depositNFT(destinationChainID, receiver, ERC721Instance.address, tokenID, "0x");

        const safeBalance = await ERC721Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(safeBalance.toNumber(), 1);

        const minterbal = await ERC721Instance.balanceOf(minter);
        assert.strictEqual(minterbal.toNumber(), 0);

        const owner = await ERC721Instance.ownerOf(tokenID);
        assert.strictEqual(owner, EmitterInstance.address);
    });

    it('depositId is correct', async () => {
        EmitterInstance.depositNFT(destinationChainID, receiver, ERC721Instance.address, tokenID, "0x");
        const depositId = await EmitterInstance.DepositCounts.call(destinationChainID);
        assert.strictEqual(depositId.toNumber(), 1);
    });
});
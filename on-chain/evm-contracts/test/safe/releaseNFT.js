/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const EmitterContract = artifacts.require("Emitter");
const ERC721Contract = artifacts.require("ERC721Mintable");

contract('Safe - [release::ERC721]', async (accounts) => {
    const minter = accounts[0];
    const receiver = accounts[1];
    const tokenID = 1;

    let EmitterInstance;
    let ERC721Instance;

    beforeEach(async () => {
        EmitterInstance = await EmitterContract.new();
        ERC721Instance = await ERC721Contract.new();
        await ERC721Instance.mint(minter, tokenID);
        await ERC721Instance.approve(EmitterInstance.address, tokenID);
    });

    it('[sanity] test minter balance', async () => {
        const tokenOwner = await ERC721Instance.ownerOf(tokenID);
        assert.strictEqual(tokenOwner, minter);

        const minterBalance = await ERC721Instance.balanceOf(minter);
        assert.strictEqual(minterBalance.toNumber(), 1);

        const contractBalance = await ERC721Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(contractBalance.toNumber(), 0);

        const approvedAddress = await ERC721Instance.getApproved(tokenID);
        assert.strictEqual(approvedAddress, EmitterInstance.address);
    });

    xit('release', async () => {
        let safeBalance;
        let minterBalance;
        let owner;

        // Staging the deposit
        EmitterInstance.depositNFT(0, receiver, ERC721Instance.address, tokenID, "0x");

        safeBalance = await ERC721Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(safeBalance.toNumber(), 1);

        minterBalance = await ERC721Instance.balanceOf(minter);
        assert.strictEqual(minterBalance.toNumber(), 0);

        owner = await ERC721Instance.ownerOf(tokenID);
        assert.strictEqual(owner, EmitterInstance.address);

        // Releasing deposited token
        EmitterInstance.releaseNFT(ERC721Instance.address, minter, tokenID);

        safeBalance = await ERC721Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(safeBalance.toNumber(), 0);

        minterBalance = await ERC721Instance.balanceOf(minter);
        assert.strictEqual(minterBalance.toNumber(), 1);

        owner = await ERC721Instance.ownerOf(tokenID);
        assert.strictEqual(owner, minter);
    });
});
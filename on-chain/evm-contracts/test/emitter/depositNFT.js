/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const EmitterContract = artifacts.require("Emitter");
const NFTContract = artifacts.require("ERC721Mintable");

contract('Emitter - [deposits::ERC20]', async (accounts) => {
    let EmitterInstance;
    let NFTInstance;
    let EmitterAddress;
    let NFTAddress;

    let minter = accounts[0];
    let receiver = accounts[1];

    beforeEach(async () => {
        EmitterInstance = await EmitterContract.new();
        EmitterAddress = EmitterInstance.address;
        NFTInstance = await NFTContract.new({ from: minter });
        NFTAddress = NFTInstance.address;
        await NFTInstance.mint(minter, 1, { from: minter });
        await NFTInstance.approve(EmitterInstance.address, 1, { from: minter })
    });

    it('[sanity] test minter balance', async () => {
        const tokenOwner = await NFTInstance.ownerOf(1);
        assert.strictEqual(tokenOwner, minter);

        const minterbal = await NFTInstance.balanceOf(minter);
        assert.strictEqual(minterbal.toNumber(), 1);

        const contractbal = await NFTInstance.balanceOf(EmitterAddress);
        assert.strictEqual(contractbal.toNumber(), 0);

        const approvedAddress = await NFTInstance.getApproved(1);
        assert.strictEqual(approvedAddress, EmitterAddress);
    });

    it('deposit', async () => {
        EmitterInstance.depositNFT(0, receiver, NFTAddress, 1, "0x");

        const safeBalance = await NFTInstance.balanceOf(EmitterAddress);
        assert.strictEqual(safeBalance.toNumber(), 1);

        const minterbal = await NFTInstance.balanceOf(minter);
        assert.strictEqual(minterbal.toNumber(), 0);

        const owner = await NFTInstance.ownerOf(1);
        assert.strictEqual(owner, EmitterAddress);
    });
});
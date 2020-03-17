/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const EmitterContract = artifacts.require("Emitter");
const ERC20Contract = artifacts.require("ERC20Mintable");

contract('Emitter - [deposits::ERC20]', async (accounts) => {
    const minter = accounts[0];
    const receiver = accounts[1];
    const destinationChainID = 0;
    const initialTokenAmount = 100;
    const depositTokenAmount = 1;

    let EmitterInstance;
    let ERC20Instance;

    beforeEach(async () => {
        EmitterInstance = await EmitterContract.new();
        ERC20Instance = await ERC20Contract.new();
        await ERC20Instance.mint(minter, initialTokenAmount);
        await ERC20Instance.approve(EmitterInstance.address, initialTokenAmount);
    });

    it('[sanity] test minter balance', async () => {
        const minterBalance = await ERC20Instance.balanceOf(minter);
        assert.strictEqual(minterBalance.toNumber(), initialTokenAmount);
        const allowance = await ERC20Instance.allowance(minter, EmitterInstance.address);
        assert.strictEqual(allowance.toNumber(), initialTokenAmount);
    });

    it('deposit', async () => {
        EmitterInstance.depositGenericErc(destinationChainID, depositTokenAmount, receiver, ERC20Instance.address);
        
        const safeBalance = await ERC20Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(safeBalance.toNumber(), depositTokenAmount);

        const minterBalNew = await ERC20Instance.balanceOf(minter);
        assert.strictEqual(minterBalNew.toNumber(), initialTokenAmount - depositTokenAmount)
    });

    it('token balance is correct', async () => {
        EmitterInstance.depositGenericErc(destinationChainID, depositTokenAmount, receiver, ERC20Instance.address);
        const balance = await EmitterInstance.balances(ERC20Instance.address);
        assert.strictEqual(balance.toNumber(), depositTokenAmount);
    });

    it('depositId is correct', async () => {
        EmitterInstance.depositGenericErc(destinationChainID, depositTokenAmount, receiver, ERC20Instance.address);
        const depositId = await EmitterInstance.DepositCounts.call(destinationChainID);
        assert.strictEqual(depositId.toNumber(), 1);
    });
});
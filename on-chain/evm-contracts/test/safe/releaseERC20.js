const EmitterContract = artifacts.require("Emitter");
const ERC20Contract = artifacts.require("ERC20Mintable");

contract('Safe - [release::ERC20]', async (accounts) => {
    const minter = accounts[0];
    const receiver = accounts[1];
    const initialTokenAmount = 100;
    const depositTokenAmount = 1;

    let EmitterInstance;
    let ERC20Instance;

    beforeEach(async () => {
        EmitterInstance = await EmitterContract.new();
        ERC20Instance = await ERC20Contract.new();
        await ERC20Instance.mint(minter, initialTokenAmount);
        await ERC20Instance.approve(EmitterInstance.address, initialTokenAmount)
    });

    it('[sanity] test minter balance', async () => {
        const minterBalance = await ERC20Instance.balanceOf(minter);
        assert.strictEqual(minterBalance.toNumber(), initialTokenAmount);
        const allowance = await ERC20Instance.allowance(minter, EmitterInstance.address);
        assert.strictEqual(allowance.toNumber(), initialTokenAmount);
    });

    xit('release', async () => {
        let safeBalance;
        let minterBalance;
        let receiverBalance;

        // Staging the deposit
        EmitterInstance.depositGenericErc(0, depositTokenAmount, receiver, ERC20Instance.address);

        safeBalance = await ERC20Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(safeBalance.toNumber(), depositTokenAmount);

        minterBalance = await ERC20Instance.balanceOf(minter);
        assert.strictEqual(minterBalance.toNumber(), initialTokenAmount - depositTokenAmount);

        receiverBalance = await ERC20Instance.balanceOf(receiver);
        assert.strictEqual(receiverBalance.toNumber(), 0);

        // Releasing deposited token
        EmitterInstance.releaseErc(ERC20Instance.address, depositTokenAmount, minter);

        safeBalance = await ERC20Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(safeBalance.toNumber(), 0);

        minterBalance = await ERC20Instance.balanceOf(minter);
        assert.strictEqual(minterBalance.toNumber(), initialTokenAmount);

        receiverBalance = await ERC20Instance.balanceOf(receiver);
        assert.strictEqual(receiverBalance.toNumber(), 0);
    });
});
const EmitterContract = artifacts.require("Emitter");
const ERC20Contract = artifacts.require("ERC20Mintable");

contract('Emitter - [Deployment]', async (accounts) => {
    let EmitterInstance;
    let erc20Instance;

    let minter = accounts[0];
    let receiver = accounts[1];

    before(async () => {
        EmitterInstance = await EmitterContract.new();
        erc20Instance = await ERC20Contract.new({ from: minter });
        await erc20Instance.mint(minter, 100, { from: minter });
        await erc20Instance.approve(EmitterInstance.address, 100, {from: minter})
    });

    it('[sanity] test minter balance', async () => {
        const minterbal = await erc20Instance.balanceOf(minter);
        assert.strictEqual(minterbal.toNumber(), 100);
        const allowance = await erc20Instance.allowance(minter, EmitterInstance.address);
        assert.strictEqual(allowance.toNumber(), 100);
    });

    it('deposit ERC20', async () => {
        EmitterInstance.depositGenericErc(0, 1, receiver, erc20Instance.address);
        
        const safeBalance = await erc20Instance.balanceOf(EmitterInstance.address);
        assert.strictEqual(safeBalance.toNumber(), 1);

        const minterBalNew = await erc20Instance.balanceOf(minter);
        assert.strictEqual(minterBalNew.toNumber(), 99)
    });
});
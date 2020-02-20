const SafeContract = artifacts.require("Safe");

contract('Safe - [Deployment]', async (accounts) => {
    const contractOwner = accounts[0];

    let SafeInstance;

    before(async () => {
        SafeInstance = await SafeContract.new(contractOwner);
    });

    it('should set constructor values', async () => {
        let owner = await SafeInstance.owner.call();
        assert.strictEqual(contractOwner, owner);
    });
});

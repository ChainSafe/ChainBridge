const SafeContract = artifacts.require("Safe");

contract('Safe - [Deployment]', async (accounts) => {
    let SafeInstance;

    before(async () => {
        SafeInstance = await SafeContract.new(
            accounts[0] // contract owner 
        )
    });

    it('should set constructor values', async () => {
        let owner = await SafeInstance.owner.call();
        assert.strictEqual(accounts[0], owner);
    });
});

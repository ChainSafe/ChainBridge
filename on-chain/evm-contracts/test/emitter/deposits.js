const EmitterContract = artifacts.require("Emitter");

contract('Emitter - [Deployment]', async (accounts) => {
    let EmitterInstance;

    // Set validators
    let v1 = accounts[0];
    let v2 = accounts[1];

    before(async () => {
        EmitterInstance = await EmitterContract.new();
    });

    it('', async () => {
    });
});
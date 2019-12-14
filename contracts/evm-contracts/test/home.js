const HomeContract = artifacts.require("Home");

contract('Receiver', async (accounts) => {
    let HomeInstance;

    before(async () => {
        HomeInstance = await HomeContract.new(
            [accounts[0], accounts[1]], // bridge validators
            2,                          // depoist threshold
            2                           // validator threshold
        )
    });

    it('should set constructor values', async () => {
        let validators = await HomeInstance.TotalValidators.call();
        assert.strictEqual(parseInt(validators, 10), 2);

        let depositThreshold = await HomeInstance.voteDepositThreshold.call();
        assert.strictEqual(parseInt(depositThreshold, 10), 2);

        let validatorThreshold = await HomeInstance.voteValidatorThreshold.call();
        assert.strictEqual(parseInt(validatorThreshold, 10), 2);
    });
});
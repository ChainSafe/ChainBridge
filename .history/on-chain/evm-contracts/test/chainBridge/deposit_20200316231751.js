const truffleAssert = require('truffle-assertions');

const ValidatorContract = artifacts.require('Validator'");
const BridgeContract = artifacts.require('Bridge');
const GenericHandlerContract = artifacts.require('')

contract('Bridge - [deposit]', async (accounts) => {
    const validatorThreshold = 0;
    
    let ValidatorInstance;

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
    });

    it('', async () => {

    });
});

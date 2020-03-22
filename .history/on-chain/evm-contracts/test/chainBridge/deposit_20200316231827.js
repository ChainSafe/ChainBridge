const truffleAssert = require('truffle-assertions');

const ValidatorContract = artifacts.require('Validator');
const BridgeContract = artifacts.require('Bridge');
const GenericHandlerContract = artifacts.require('GenericHandler');

contract('Bridge - [deposit]', async (accounts) => {
    const validatorThreshold = 0;
    
    let ValidatorInstance;
    let BridgeInstance;
    let Ge

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
        GenericHandlerInstance = await GenericHandlerContract.new();
    });

    it('', async () => {

    });
});
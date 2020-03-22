const truffleAssert = require('truffle-assertions');

const ValidatorContract = artifacts.require("Validator");
const BridgeContract = artifacts.require("Bridge");

contract('Bridge - [deposit]', async (accounts) => {
    const initialValidatorThreshold = 2;
    
    let ValidatorInstance;

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([], 0);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, in);
    });

    it('', async () => {

    });
});

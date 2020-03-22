const truffleAssert = require('truffle-assertions');

const ValidatorContract = artifacts.require('Validator');
const BridgeContract = artifacts.require('Bridge');
const GenericHandlerContract = artifacts.require('GenericHandler');

contract('Bridge - [deposit]', async (accounts) => {
    const validatorThreshold = 0;
    const destinationChainID = 1;
    const destinationChainRecipientAddress = accounts [1];
    const genericMetaData = 0x736f796c656e745f677265656e5f69735f70656f706c65;
    
    let ValidatorInstance;
    let BridgeInstance;
    let GenericHandlerInstance;
    const depositData;

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
        GenericHandlerInstance = await GenericHandlerContract.new(BridgeInstance.address);
    });

    it('should make a generic deposit successfully', async () => {
        await BridgeInstance.deposit(GenericHandlerInstance.address, )
    });
});

const TruffleAssert = require('truffle-assertions');
const Ethers = require('ethers');

const ValidatorContract = artifacts.require('Validator');
const BridgeContract = artifacts.require('Bridge');
const ERC20HandlerContract = artifacts.require('ERC20Handler');

contract('Bridge - [deposit]', async (accounts) => {
    const validatorThreshold = 0;
    const destinationChainID = 1;
    const destinationChainRecipientAddress = accounts [1];
    
    let ValidatorInstance;
    let BridgeInstance;
    let ERC20HandlerInstance;

    let depositData;

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
        ERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        
    });

    it('should make a generic deposit successfully', async () => {
        TruffleAssert.passes(await BridgeInstance.deposit(
            GenericHandlerInstance.address, depositData
        ));
    });
});

const TruffleAssert = require('truffle-assertions');
const Ethers = require('ethers');

const ValidatorContract = artifacts.require('Validator');
const BridgeContract = artifacts.require('Bridge');
const GenericHandlerContract = artifacts.require('GenericHandler');

contract('Bridge - [deposit]', async (accounts) => {
    const validatorThreshold = 0;
    const destinationChainID = 1;
    const destinationChainRecipientAddress = accounts [1];
    // const genericMetaData = '0x736f796c656e745f677265656e5f69735f70656f706c65';
    
    
    let ValidatorInstance;
    let BridgeInstance;
    let GenericHandlerInstance;
    let depositData;

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
        GenericHandlerInstance = await GenericHandlerContract.new(BridgeInstance.address);

        depositData = '0x' +
            Ethers.utils.hexZeroPad(Ethers.utils.hexlify(destinationChainID), 32).substr(2) +
            Ethers.utils.hexZeroPad(destinationChainRecipientAddress, 32).substr(2) +
            Ethers.utils.hexZeroPad(genericMetaData, 32).substr(2)
    });

    it('should make a generic deposit successfully', async () => {
        TruffleAssert.passes(await BridgeInstance.deposit(
            GenericHandlerInstance.address, depositData
        ));
    });
});

const TruffleAssert = require('truffle-assertions');
const Ethers = require('ethers');


const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ValidatorContract = artifacts.require('Validator');
const BridgeContract = artifacts.require('Bridge');
const ERC20HandlerContract = artifacts.require('ERC20Handler');

contract('Bridge - [deposit]', async (accounts) => {
    const validatorThreshold = 0;
    const destinationChainID = 1;
    const initialTokenAmount = 100;
    const depositer = accounts[1];
    const depositAmount = 10;
    const destinationChainRecipientAddress = accounts[2];
    
    let ERC20MintableInstance;
    let ValidatorInstance;
    let BridgeInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20HandlerInstance;
    let depositData;

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        
        depositData = '0x' +
            Ethers.utils.hexZeroPad()
    });

    it('should make a generic deposit successfully', async () => {
        TruffleAssert.passes(await BridgeInstance.deposit(
            OriginERC20HandlerInstance.address, depositData
        ));
    });
});

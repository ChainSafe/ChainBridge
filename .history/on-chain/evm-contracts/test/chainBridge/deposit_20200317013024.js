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
    const depositerAddress = accounts[1];
    const depositAmount = 10;
    const destinationChainRecipientAddress = accounts[2];
    
    let ERC20MintableInstance;
    let ValidatorInstance;
    let BridgeInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20HandlerInstance;
    let depositData;

    beforeEach(async () => {
        ERC20MintableInstance = await ERC20MintableContract.new();
        ValidatorInstance = await ValidatorContract.new([], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        await ERC20MintableInstance.mint(depositerAddress, initialTokenAmount);
        await ERC20MintableInstance.approve(OriginERC20HandlerInstance.address, depositAmount, { from: depositerAddress });
        
        depositData = '0x' +
            Ethers.utils.hexZeroPad(OriginERC20HandlerInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(Ethers.utils.hexlify(destinationChainID), 32).substr(2) +
            Ethers.utils.hexZeroPad(DestinationERC20HandlerInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(destinationChainRecipientAddress, 32).substr(2) +
            Ethers.utils.hexZeroPad(depositerAddress, 32).substr(2) +
            Ethers.utils.hexZeroPad(Ethers.utils.hexlify(depositAmount), 32).substr(2)
    });

    it('[sanity] depositerAddress should have a balance of initialTokenAmount', async () => {
        const depositerBalance = await ERC20MintableInstance.balanceOf(depositerAddress);
        assert.strictEqual(depositerBalance.toNumber(), initialTokenAmount);
    });

    it("[sanity] OriginERC20HandlerInstance.address should have an allowance of depositAmount from depositerAddress", async () => {
        const handlerAllowance = await ERC20MintableInstance.allowance(originChainDepositerAddress, OriginERC20HandlerInstance.address);
        assert.strictEqual(originChainHandlerAllowance.toNumber(), depositAmount);
    });

    it('should make a generic deposit successfully', async () => {
        TruffleAssert.passes(await BridgeInstance.deposit(
            OriginERC20HandlerInstance.address, depositData
        ));
    });
});

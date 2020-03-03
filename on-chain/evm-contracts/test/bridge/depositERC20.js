/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const truffleAssert = require('truffle-assertions');

const ValidatorContract = artifacts.require("Validator");
const BridgeContract = artifacts.require("Bridge");
const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ERC20HandlerContract = artifacts.require("ERC20Handler");

contract('Bridge - [Deployment]', async (accounts) => {
    // const minter = accounts[0];
    const originChainDepositerAddress = accounts[1];
    const destinationChainRecipientAddress = accounts[2];
    const destinationChainID = 0;
    const originChainInitialTokenAmount = 100;
    const depositAmount = 10;

    let ValidatorInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;

    before(async () => {
        ValidatorInstance = await ValidatorContract.new();
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        await OriginERC20MintableInstance.mint(originChainDepositerAddress, originChainInitialTokenAmount);
        await OriginERC20MintableInstance.approve(OriginERC20HandlerInstance.address, depositAmount, { from: originChainDepositerAddress });
    });

    it("[sanity] test originChainDepositerAddress' balance", async () => {
        const originChainDepositerBalance = await OriginERC20MintableInstance.balanceOf(originChainDepositerAddress);
        assert.strictEqual(originChainDepositerBalance.toNumber(), originChainInitialTokenAmount);
    });

    it("[sanity] test OriginERC20HandlerInstance.address' allowance", async () => {
        const originChainHandlerAllowance = await OriginERC20MintableInstance.allowance(originChainDepositerAddress, OriginERC20HandlerInstance.address);
        assert.strictEqual(originChainHandlerAllowance.toNumber(), depositAmount);
    });

    it('Bridge - [deposits::ERC20]', async () => {
        const expectedDepositID = 0;

        const depositTx = await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        truffleAssert.eventEmitted(depositTx, 'ERC20Deposited', (event) => {
            return event.depositID.toNumber() === expectedDepositID
        });
    });
});
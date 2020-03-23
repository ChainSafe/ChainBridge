/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const truffleAssert = require('truffle-assertions');

const RelayerContract = artifacts.require("Relayer");
const BridgeContract = artifacts.require("Bridge");
const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ERC20HandlerContract = artifacts.require("ERC20Handler");

contract('Bridge - [depositERC20]', async (accounts) => {
    // const minter = accounts[0];
    const originChainDepositerAddress = accounts[1];
    const destinationChainRecipientAddress = accounts[2];
    const destinationChainID = 0;
    const originChainInitialTokenAmount = 100;
    const depositAmount = 10;
    const expectedDepositNonce = 1;

    let RelayerInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let expectedDepositRecord;

    beforeEach(async () => {
        RelayerInstance = await RelayerContract.new([], 0);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, 0);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        await OriginERC20MintableInstance.mint(originChainDepositerAddress, originChainInitialTokenAmount);
        await OriginERC20MintableInstance.approve(OriginERC20HandlerInstance.address, depositAmount, { from: originChainDepositerAddress });

        expectedDepositRecord = {
            _originChainTokenAddress: OriginERC20MintableInstance.address,
            _originChainHandlerAddress: OriginERC20HandlerInstance.address,
            _destinationChainID: destinationChainID,
            _destinationChainHandlerAddress: DestinationERC20HandlerInstance.address,
            _destinationRecipientAddress: destinationChainRecipientAddress,
            _amount: depositAmount
        };
    });

    it("[sanity] test originChainDepositerAddress' balance", async () => {
        const originChainDepositerBalance = await OriginERC20MintableInstance.balanceOf(originChainDepositerAddress);
        assert.strictEqual(originChainDepositerBalance.toNumber(), originChainInitialTokenAmount);
    });

    it("[sanity] test OriginERC20HandlerInstance.address' allowance", async () => {
        const originChainHandlerAllowance = await OriginERC20MintableInstance.allowance(originChainDepositerAddress, OriginERC20HandlerInstance.address);
        assert.strictEqual(originChainHandlerAllowance.toNumber(), depositAmount);
    });

    it('ERC20 deposit can be made', async () => {
        truffleAssert.passes(await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        ));
    });

    it('_depositCounts should be increments from 0 to 1', async () => {
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        const depositCount = await BridgeInstance._depositCounts.call(destinationChainID);
        assert.strictEqual(depositCount.toNumber(), expectedDepositNonce);
    });

    it('getDepositCounts should return correct expectedDepositNonce', async () => {
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        const depositCount = await BridgeInstance.getDepositCount(destinationChainID);
        assert.strictEqual(depositCount.toNumber(), expectedDepositNonce);
    });

    it('ERC20 can be deposited with correct balances', async () => {
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        const originChainDepositerBalance = await OriginERC20MintableInstance.balanceOf(originChainDepositerAddress);
        assert.strictEqual(originChainDepositerBalance.toNumber(), originChainInitialTokenAmount - depositAmount);

        const originChainHandlerBalance = await OriginERC20MintableInstance.balanceOf(OriginERC20HandlerInstance.address);
        assert.strictEqual(originChainHandlerBalance.toNumber(), depositAmount);
    });

    it('ERC20 deposit record is created with expected depositNonce and values', async () => {
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        const depositRecord = await BridgeInstance._erc20DepositRecords.call(destinationChainID, expectedDepositNonce);
        for (const expectedProperty of Object.keys(expectedDepositRecord)) {
            // Testing all expected object properties
            assert.property(depositRecord, expectedProperty, `property: ${expectedProperty} does not exist in depositRecord`);

            // Testing all expected object values
            const depositRecordValue = depositRecord[expectedProperty].toNumber !== undefined ?
                depositRecord[expectedProperty].toNumber() : depositRecord[expectedProperty];
            assert.strictEqual(
                depositRecordValue, expectedDepositRecord[expectedProperty],
                `property: ${expectedProperty}'s value: ${depositRecordValue} does not match expected value: ${expectedDepositRecord[expectedProperty]}`)
        }
    });

    it('ERC20Deposited event is fired with expected value', async () => {
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
            return event.depositNonce.toNumber() === expectedDepositNonce
        });
    });

    it('getERC20DepositRecord should return correct depositNonce with values in expected order', async () => {
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        const depositRecord = await BridgeInstance.getERC20DepositRecord(destinationChainID, expectedDepositNonce);
        const depositRecordValues = Object.values(depositRecord);
        depositRecordValues.forEach((depositRecordValue, index) => {
            depositRecordValues[index] = depositRecordValue.toNumber !== undefined ?
                depositRecordValue.toNumber() : depositRecordValue;
        });
        assert.sameOrderedMembers(depositRecordValues, Object.values(expectedDepositRecord));
    });
});
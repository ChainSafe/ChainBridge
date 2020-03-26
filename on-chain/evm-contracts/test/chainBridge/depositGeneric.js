/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const truffleAssert = require('truffle-assertions');

const RelayerContract = artifacts.require("Relayer");
const BridgeContract = artifacts.require("Bridge");
const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ERC20HandlerContract = artifacts.require("ERC20Handler");

contract('Bridge - [depositGeneric]', async (accounts) => {
    // const minter = accounts[0];
    const originChainDepositerAddress = accounts[1];
    const destinationChainRecipientAddress = accounts[2];
    const destinationChainID = 0;
    const expectedDepositNonce = 1;
    const genericBytes = '0x736f796c656e745f677265656e5f69735f70656f706c65';

    let RelayerInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let expectedDepositRecord_PartialArguments;
    let expectedDepositRecord_AllArguments;

    beforeEach(async () => {
        RelayerInstance = await RelayerContract.new([], 0);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, 0);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        expectedDepositRecord_PartialArguments = {
            _originChainTokenAddress: '0x0000000000000000000000000000000000000000',
            _originChainHandlerAddress: '0x0000000000000000000000000000000000000000',
            _destinationChainID: destinationChainID,
            _destinationChainHandlerAddress: '0x0000000000000000000000000000000000000000',
            _destinationRecipientAddress: destinationChainRecipientAddress,
            _data: genericBytes
        };

        expectedDepositRecord_AllArguments = {
            _originChainTokenAddress: OriginERC20MintableInstance.address,
            _originChainHandlerAddress: OriginERC20HandlerInstance.address,
            _destinationChainID: destinationChainID,
            _destinationChainHandlerAddress: DestinationERC20HandlerInstance.address,
            _destinationRecipientAddress: destinationChainRecipientAddress,
            _data: genericBytes
        };
    });

    it('Generic deposit can be made with partial arguments', async () => {
        truffleAssert.passes(await BridgeInstance.depositGeneric(
            destinationChainID,
            destinationChainRecipientAddress,
            genericBytes
        ));
    });

    it('_depositCounts is incremented correctly after Generic deposit with partial arguments', async () => {
        await BridgeInstance.depositGeneric(
            destinationChainID,
            destinationChainRecipientAddress,
            genericBytes
        );

        const depositCount = await BridgeInstance._depositCounts.call(destinationChainID);
        assert.strictEqual(depositCount.toNumber(), expectedDepositNonce);
    });

    it('Generic deposit with partial arguments is stored correctly', async () => {
        await BridgeInstance.depositGeneric(
            destinationChainID,
            destinationChainRecipientAddress,
            genericBytes
        );

        const depositRecord = await BridgeInstance._genericDepositRecords.call(destinationChainID, expectedDepositNonce);
        for (const expectedProperty of Object.keys(expectedDepositRecord_PartialArguments)) {
            // Testing all expected object properties
            assert.property(depositRecord, expectedProperty, `property: ${expectedProperty} does not exist in depositRecord`);

            // Testing all expected object values
            const depositRecordValue = depositRecord[expectedProperty].toNumber !== undefined ?
                depositRecord[expectedProperty].toNumber() : depositRecord[expectedProperty];
            assert.strictEqual(
                depositRecordValue, expectedDepositRecord_PartialArguments[expectedProperty],
                `property: ${expectedProperty}'s value: ${depositRecordValue} does not match expected value: ${expectedDepositRecord_PartialArguments[expectedProperty]}`)
        }
    });

    it('GenericDeposited event is fired with expected value after Generic deposit with partial arguments', async () => {
        const depositTx = await BridgeInstance.depositGeneric(
            destinationChainID,
            destinationChainRecipientAddress,
            genericBytes
        );

        truffleAssert.eventEmitted(depositTx, 'GenericDeposited', (event) => {
            return event.depositNonce.toNumber() === expectedDepositNonce
        });
    });

    it('getGenericDepositRecord should return correct depositNonce with values in expected order for Generic deposit with partial arguments', async () => {
        await BridgeInstance.depositGeneric(
            destinationChainID,
            destinationChainRecipientAddress,
            genericBytes
        );

        const depositRecord = await BridgeInstance.getGenericDepositRecord(destinationChainID, expectedDepositNonce);
        const depositRecordValues = Object.values(depositRecord);
        depositRecordValues.forEach((depositRecordValue, index) => {
            depositRecordValues[index] = depositRecordValue.toNumber !== undefined ?
                depositRecordValue.toNumber() : depositRecordValue;
        });
        assert.sameOrderedMembers(depositRecordValues, Object.values(expectedDepositRecord_PartialArguments));
    });

    it('Generic deposit can be made with all arguments', async () => {
        truffleAssert.passes(
            await BridgeInstance.methods['depositGeneric(address,address,uint256,address,address,bytes)'](
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            genericBytes,
            { from: originChainDepositerAddress }
        ));
    });

    it('_depositCounts is incremented correctly after Generic deposit with all arguments', async () => {
        await BridgeInstance.methods['depositGeneric(address,address,uint256,address,address,bytes)'](
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            genericBytes,
            { from: originChainDepositerAddress }
        );

        const depositCount = await BridgeInstance._depositCounts.call(destinationChainID);
        assert.strictEqual(depositCount.toNumber(), expectedDepositNonce);
    });

    it('Generic deposit with all arguments is stored correctly', async () => {
        await BridgeInstance.methods['depositGeneric(address,address,uint256,address,address,bytes)'](
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            genericBytes,
            { from: originChainDepositerAddress }
        );

        const depositRecord = await BridgeInstance._genericDepositRecords.call(destinationChainID, expectedDepositNonce);
        for (const expectedProperty of Object.keys(expectedDepositRecord_AllArguments)) {
            // Testing all expected object properties
            assert.property(depositRecord, expectedProperty, `property: ${expectedProperty} does not exist in depositRecord`);

            // Testing all expected object values
            const depositRecordValue = depositRecord[expectedProperty].toNumber !== undefined ?
                depositRecord[expectedProperty].toNumber() : depositRecord[expectedProperty];
            assert.strictEqual(
                depositRecordValue, expectedDepositRecord_AllArguments[expectedProperty],
                `property: ${expectedProperty}'s value: ${depositRecordValue} does not match expected value: ${expectedDepositRecord_AllArguments[expectedProperty]}`)
        }
    });

    it('GenericDeposited event is fired with expected value after Generic deposit with all arguments', async () => {
        // This is how Truffle provides support for overloaded functions
        const depositTx = await BridgeInstance.methods['depositGeneric(address,address,uint256,address,address,bytes)'](
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            genericBytes,
            { from: originChainDepositerAddress }
        );

        truffleAssert.eventEmitted(depositTx, 'GenericDeposited', (event) => {
            return event.depositNonce.toNumber() === expectedDepositNonce
        });
    });

    it('getGenericDepositRecord should return correct depositNonce with values in expected order for Generic deposit with all arguments', async () => {
        // This is how Truffle provides support for overloaded functions
        await BridgeInstance.methods['depositGeneric(address,address,uint256,address,address,bytes)'](
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            genericBytes,
            { from: originChainDepositerAddress }
        );

        const depositRecord = await BridgeInstance.getGenericDepositRecord(destinationChainID, expectedDepositNonce);
        const depositRecordValues = Object.values(depositRecord);
        depositRecordValues.forEach((depositRecordValue, index) => {
            depositRecordValues[index] = depositRecordValue.toNumber !== undefined ?
                depositRecordValue.toNumber() : depositRecordValue;
        });
        assert.sameOrderedMembers(depositRecordValues, Object.values(expectedDepositRecord_AllArguments));
    });
});
/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const truffleAssert = require('truffle-assertions');
const ethers = require('ethers');

const RelayerContract = artifacts.require("Relayer");
const BridgeContract = artifacts.require("Bridge");
const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ERC20HandlerContract = artifacts.require("ERC20Handler");

contract('Bridge - [createDepositProposal with relayerThreshold = 1]', async (accounts) => {
    // const minterAndRelayer = accounts[0];
    const originChainRelayerAddress = accounts[1];
    const originChainDepositerAddress = accounts[2];
    const destinationChainRecipientAddress = accounts[3];
    const originChainID = 0;
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
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
        RelayerInstance = await RelayerContract.new([originChainRelayerAddress], 1);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, 1);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        await OriginERC20MintableInstance.mint(originChainDepositerAddress, originChainInitialTokenAmount);
        await OriginERC20MintableInstance.approve(OriginERC20HandlerInstance.address, depositAmount, { from: originChainDepositerAddress });
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        data = '0x' + ethers.utils.hexZeroPad(DestinationERC20MintableInstance.address, 32).substr(2) +
            ethers.utils.hexZeroPad(DestinationERC20HandlerInstance.address, 32).substr(2) +
            ethers.utils.hexZeroPad(ethers.utils.hexlify(depositAmount), 32).substr(2);
        dataHash = ethers.utils.keccak256(data);
    });

    it('[sanity] ERC20 deposit record is created with expected depositNonce and values', async () => {
        const expectedDepositRecord = {
            _originChainTokenAddress: OriginERC20MintableInstance.address,
            _originChainHandlerAddress: OriginERC20HandlerInstance.address,
            _destinationChainID: destinationChainID,
            _destinationChainHandlerAddress: DestinationERC20HandlerInstance.address,
            _destinationRecipientAddress: destinationChainRecipientAddress,
            _amount: depositAmount
        };

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

    it('depositProposal can be created', async () => {
        truffleAssert.passes(await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    it('Only relayers should be able to create a deposit proposal', async () => {
        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainDepositerAddress }
        ));
    });

    it('depositProposal is created with expected values', async () => {
        const expectedDepositProposal = {
            _destinationChainID: destinationChainID,
            _depositNonce: expectedDepositNonce,
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 3 // passed
        };

        await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        );

        const depositProposal = await BridgeInstance._depositProposals.call(destinationChainID, expectedDepositNonce);
        for (const expectedProperty of Object.keys(expectedDepositProposal)) {
            // Testing all expected object properties
            assert.property(depositProposal, expectedProperty, `property: ${expectedProperty} does not exist in depositRecord`);

            // Testing all expected object values
            const depositProposalValue = depositProposal[expectedProperty].toNumber !== undefined ?
                depositProposal[expectedProperty].toNumber() : depositProposal[expectedProperty];
            assert.strictEqual(
                depositProposalValue, expectedDepositProposal[expectedProperty],
                `property: ${expectedProperty}'s value: ${depositProposalValue} does not match expected value: ${expectedDepositProposal[expectedProperty]}`)
        }
    });

    it("depositProposal shouldn't be created if it doesn't have Inactive or Denied status", async () => {
        await truffleAssert.passes(BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        ));

        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    it('originChainRelayerAddress should be marked as vote for proposal', async () => {
        await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        );
        const hasVoted = await BridgeInstance.hasVoted(destinationChainID, expectedDepositNonce, originChainRelayerAddress);
        assert.isTrue(hasVoted);
    });

    it('DepositProposalCreated event is fired with expected value', async () => {
        const proposalTx = await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        );

        truffleAssert.eventEmitted(proposalTx, 'DepositProposalCreated', (event) => {
            return event.destinationChainID.toNumber() === destinationChainID &&
                event.depositNonce.toNumber() === expectedDepositNonce &&
                event.dataHash === dataHash
        });
    });

    it('getDepositProposal should return correct values in expected order', async () => {
        const expectedDepositProposal = {
            _destinationChainID: destinationChainID,
            _depositNonce: expectedDepositNonce,
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 'passed'
        };

        await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        );

        const depositProposal = await BridgeInstance.getDepositProposal(destinationChainID, expectedDepositNonce);
        const depositProposalValues = Object.values(depositProposal);
        depositProposalValues.forEach((depositRecordValue, index) => {
            depositProposalValues[index] = depositRecordValue.toNumber !== undefined ?
                depositRecordValue.toNumber() : depositRecordValue;
        });
        assert.sameOrderedMembers(depositProposalValues, Object.values(expectedDepositProposal));
    });
});

contract('Bridge - [createDepositProposal with relayerThreshold > 1]', async (accounts) => {
    // const minterAndRelayer = accounts[0];
    const originChainRelayerAddress = accounts[1];
    const originChainDepositerAddress = accounts[2];
    const destinationChainRecipientAddress = accounts[3];
    const originChainID = 0;
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
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
        RelayerInstance = await RelayerContract.new([originChainRelayerAddress], 1);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, 2);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        await OriginERC20MintableInstance.mint(originChainDepositerAddress, originChainInitialTokenAmount);
        await OriginERC20MintableInstance.approve(OriginERC20HandlerInstance.address, depositAmount, { from: originChainDepositerAddress });
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: originChainDepositerAddress }
        );

        data = '0x' + ethers.utils.hexZeroPad(DestinationERC20MintableInstance.address, 32).substr(2) +
            ethers.utils.hexZeroPad(DestinationERC20HandlerInstance.address, 32).substr(2) +
            ethers.utils.hexZeroPad(ethers.utils.hexlify(depositAmount), 32).substr(2);
        dataHash = ethers.utils.keccak256(data);
    });

    it('[sanity] ERC20 deposit record is created with expected depositNonce and values', async () => {
        const expectedDepositRecord = {
            _originChainTokenAddress: OriginERC20MintableInstance.address,
            _originChainHandlerAddress: OriginERC20HandlerInstance.address,
            _destinationChainID: destinationChainID,
            _destinationChainHandlerAddress: DestinationERC20HandlerInstance.address,
            _destinationRecipientAddress: destinationChainRecipientAddress,
            _amount: depositAmount
        };

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

    it('depositProposal can be created', async () => {
        truffleAssert.passes(await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    it('Only relayers should be able to create a deposit proposal', async () => {
        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainDepositerAddress }
        ));
    });

    it('depositProposal is created with expected values', async () => {
        const expectedDepositProposal = {
            _destinationChainID: destinationChainID,
            _depositNonce: expectedDepositNonce,
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 1 // passed
        };

        await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        );

        const depositProposal = await BridgeInstance._depositProposals.call(destinationChainID, expectedDepositNonce);
        for (const expectedProperty of Object.keys(expectedDepositProposal)) {
            // Testing all expected object properties
            assert.property(depositProposal, expectedProperty, `property: ${expectedProperty} does not exist in depositRecord`);

            // Testing all expected object values
            const depositProposalValue = depositProposal[expectedProperty].toNumber !== undefined ?
                depositProposal[expectedProperty].toNumber() : depositProposal[expectedProperty];
            assert.strictEqual(
                depositProposalValue, expectedDepositProposal[expectedProperty],
                `property: ${expectedProperty}'s value: ${depositProposalValue} does not match expected value: ${expectedDepositProposal[expectedProperty]}`)
        }
    });

    it("depositProposal shouldn't be created if it doesn't have Inactive or Denied status", async () => {
        await truffleAssert.passes(BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        ));

        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    it('originChainRelayerAddress should be marked as voted for proposal', async () => {
        await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        );
        const hasVoted = await BridgeInstance.hasVoted(destinationChainID, expectedDepositNonce, originChainRelayerAddress);
        assert.isTrue(hasVoted);
    });

    it('DepositProposalCreated event is fired with expected value', async () => {
        const proposalTx = await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        );

        truffleAssert.eventEmitted(proposalTx, 'DepositProposalCreated', (event) => {
            return event.destinationChainID.toNumber() === destinationChainID &&
                event.depositNonce.toNumber() === expectedDepositNonce &&
                event.dataHash === dataHash
        });
    });

    it('getDepositProposal should return correct values in expected order', async () => {
        const expectedDepositProposal = {
            _destinationChainID: destinationChainID,
            _depositNonce: expectedDepositNonce,
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 'active'
        };

        await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
            dataHash,
            { from: originChainRelayerAddress }
        );

        const depositProposal = await BridgeInstance.getDepositProposal(destinationChainID, expectedDepositNonce);
        const depositProposalValues = Object.values(depositProposal);
        depositProposalValues.forEach((depositRecordValue, index) => {
            depositProposalValues[index] = depositRecordValue.toNumber !== undefined ?
                depositRecordValue.toNumber() : depositRecordValue;
        });
        assert.sameOrderedMembers(depositProposalValues, Object.values(expectedDepositProposal));
    });
});
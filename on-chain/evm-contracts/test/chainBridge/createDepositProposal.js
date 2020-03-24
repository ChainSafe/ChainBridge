/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const TruffleAssert = require('truffle-assertions');
const Ethers = require('ethers');

const RelayerContract = artifacts.require("Relayer");
const BridgeContract = artifacts.require("Bridge");
const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ERC20HandlerContract = artifacts.require("ERC20Handler");

contract('Bridge - [createDepositProposal with relayerThreshold = 1]', async (accounts) => {
    const originChainRelayerAddress = accounts[1];
    const depositerAddress = accounts[2];
    const originChainID = 0;
    const depositAmount = 10;
<<<<<<< HEAD
    const expectedDepositID = 1;
    const relayerThreshold = 1;
=======
    const expectedDepositNonce = 1;
>>>>>>> master

    let RelayerInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
        RelayerInstance = await RelayerContract.new([originChainRelayerAddress], relayerThreshold);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, relayerThreshold);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        data = '0x' +
            Ethers.utils.hexZeroPad(DestinationERC20MintableInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(DestinationERC20HandlerInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(Ethers.utils.hexlify(depositAmount), 32).substr(2);
        dataHash = Ethers.utils.keccak256(data);
    });

<<<<<<< HEAD
    it('should create depositProposal successfully', async () => {
        TruffleAssert.passes(await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });
=======
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
>>>>>>> master

    it('should revert because depositerAddress is not a relayer', async () => {
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: depositerAddress }
        ));
    });

<<<<<<< HEAD
    it("depositProposal shouldn't be created if it has an Active status", async () => {
        await TruffleAssert.passes(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
=======
    it('depositProposal can be created', async () => {
        truffleAssert.passes(await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
>>>>>>> master
            dataHash,
            { from: originChainRelayerAddress }
        ));

<<<<<<< HEAD
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
=======
    it('Only relayers should be able to create a deposit proposal', async () => {
        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
>>>>>>> master
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    xit("depositProposal shouldn't be created if it has a Passed status", async () => {
        
    });

    xit("depositProposal shouldn't be created if it has a Transferred status", async () => {
        
    });

    it('depositProposal should be created with expected values', async () => {
        const expectedDepositProposal = {
<<<<<<< HEAD
=======
            _destinationChainID: destinationChainID,
            _depositNonce: expectedDepositNonce,
>>>>>>> master
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 3 // passed
        };

        await BridgeInstance.createDepositProposal(
<<<<<<< HEAD
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
=======
            destinationChainID,
            expectedDepositNonce,
>>>>>>> master
            dataHash,
            { from: originChainRelayerAddress }
        );

<<<<<<< HEAD
        const depositProposal = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositID);
=======
        const depositProposal = await BridgeInstance._depositProposals.call(destinationChainID, expectedDepositNonce);
>>>>>>> master
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

<<<<<<< HEAD
    it('originChainRelayerAddress should be marked as voted for proposal', async () => {
        await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );
        const hasVoted = await BridgeInstance.hasVoted(
            originChainID, OriginERC20HandlerInstance.address,
            expectedDepositID, originChainRelayerAddress);
=======
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
>>>>>>> master
        assert.isTrue(hasVoted);
    });

    it('DepositProposalCreated event  should be emitted with expected values', async () => {
        const proposalTx = await BridgeInstance.createDepositProposal(
<<<<<<< HEAD
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
=======
            destinationChainID,
            expectedDepositNonce,
>>>>>>> master
            dataHash,
            { from: originChainRelayerAddress }
        );

<<<<<<< HEAD
        TruffleAssert.eventEmitted(proposalTx, 'DepositProposalCreated', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.originChainHandlerAddress === OriginERC20HandlerInstance.address &&
                event.depositID.toNumber() === expectedDepositID &&
                event.dataHash === dataHash
        });
    });
=======
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
>>>>>>> master
});

contract('Bridge - [createDepositProposal with relayerThreshold > 1]', async (accounts) => {
    // const minterAndRelayer = accounts[0];
    const originChainRelayerAddress = accounts[1];
    const depositerAddress = accounts[2];
    const destinationChainRecipientAddress = accounts[3];
    const originChainID = 0;
    const destinationChainID = 0;
    const originChainInitialTokenAmount = 100;
    const depositAmount = 10;
<<<<<<< HEAD
    const expectedDepositID = 1;
    const relayerThreshold = 2;
=======
    const expectedDepositNonce = 1;
>>>>>>> master

    let RelayerInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
        RelayerInstance = await RelayerContract.new([originChainRelayerAddress], relayerThreshold);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, relayerThreshold);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        data = '0x' +
            Ethers.utils.hexZeroPad(DestinationERC20MintableInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(DestinationERC20HandlerInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(Ethers.utils.hexlify(depositAmount), 32).substr(2);
        dataHash = Ethers.utils.keccak256(data);
    });

<<<<<<< HEAD
    it('should create depositProposal successfully', async () => {
        TruffleAssert.passes(await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });
=======
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
>>>>>>> master

    it('should revert because depositerAddress is not a relayer', async () => {
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: depositerAddress }
        ));
    });

<<<<<<< HEAD
    it("depositProposal shouldn't be created if it has an Active status", async () => {
        await TruffleAssert.passes(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
=======
    it('depositProposal can be created', async () => {
        truffleAssert.passes(await BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
>>>>>>> master
            dataHash,
            { from: originChainRelayerAddress }
        ));

<<<<<<< HEAD
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
=======
    it('Only relayers should be able to create a deposit proposal', async () => {
        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
            destinationChainID,
            expectedDepositNonce,
>>>>>>> master
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    xit("depositProposal shouldn't be created if it has a Passed status", async () => {
        
    });

    xit("depositProposal shouldn't be created if it has a Transferred status", async () => {
        
    });

    it('depositProposal should be created with expected values', async () => {
        const expectedDepositProposal = {
<<<<<<< HEAD
=======
            _destinationChainID: destinationChainID,
            _depositNonce: expectedDepositNonce,
>>>>>>> master
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 1 // active
        };

        await BridgeInstance.createDepositProposal(
<<<<<<< HEAD
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
=======
            destinationChainID,
            expectedDepositNonce,
>>>>>>> master
            dataHash,
            { from: originChainRelayerAddress }
        );

<<<<<<< HEAD
        const depositProposal = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositID);
=======
        const depositProposal = await BridgeInstance._depositProposals.call(destinationChainID, expectedDepositNonce);
>>>>>>> master
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

<<<<<<< HEAD
    it('originChainRelayerAddress should be marked as voted for proposal', async () => {
        await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );
        const hasVoted = await BridgeInstance.hasVoted(
            originChainID, OriginERC20HandlerInstance.address,
            expectedDepositID, originChainRelayerAddress);
=======
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
>>>>>>> master
        assert.isTrue(hasVoted);
    });

    it('DepositProposalCreated event  should be emitted with expected values', async () => {
        const proposalTx = await BridgeInstance.createDepositProposal(
<<<<<<< HEAD
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
=======
            destinationChainID,
            expectedDepositNonce,
>>>>>>> master
            dataHash,
            { from: originChainRelayerAddress }
        );

<<<<<<< HEAD
        TruffleAssert.eventEmitted(proposalTx, 'DepositProposalCreated', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.originChainHandlerAddress === OriginERC20HandlerInstance.address &&
                event.depositID.toNumber() === expectedDepositID &&
                event.dataHash === dataHash
        });
    });
=======
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
>>>>>>> master
});
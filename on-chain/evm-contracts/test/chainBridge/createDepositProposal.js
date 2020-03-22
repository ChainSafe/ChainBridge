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

<<<<<<< HEAD
contract('Bridge - [createDepositProposal with validatorThreshold = 1]', async (accounts) => {
    const originChainValidatorAddress = accounts[1];
    const depositerAddress = accounts[2];
=======
contract('Bridge - [createDepositProposal with relayerThreshold = 1]', async (accounts) => {
    // const minterAndRelayer = accounts[0];
    const originChainRelayerAddress = accounts[1];
    const originChainDepositerAddress = accounts[2];
    const destinationChainRecipientAddress = accounts[3];
>>>>>>> master
    const originChainID = 0;
    const depositAmount = 10;
    const expectedDepositID = 1;
    const validatorThreshold = 1;

    let RelayerInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
<<<<<<< HEAD
        ValidatorInstance = await ValidatorContract.new([originChainValidatorAddress], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
=======
        RelayerInstance = await RelayerContract.new([originChainRelayerAddress], 1);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, 1);
>>>>>>> master
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

    it('should create depositProposal successfully', async () => {
        TruffleAssert.passes(await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        ));
    });

    it('should revert because depositerAddress is not a validator', async () => {
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: depositerAddress }
        ));
    });

    it("depositProposal shouldn't be created if it has an Active status", async () => {
        await TruffleAssert.passes(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        ));

<<<<<<< HEAD
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
=======
    it('Only relayers should be able to create a deposit proposal', async () => {
        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
>>>>>>> master
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        ));
    });

    xit("depositProposal shouldn't be created if it has a Passed status", async () => {
        
    });

    xit("depositProposal shouldn't be created if it has a Transferred status", async () => {
        
    });

    it('depositProposal should be created with expected values', async () => {
        const expectedDepositProposal = {
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 3 // passed
        };

        await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );

        const depositProposal = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositID);
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
    it('originChainValidatorAddress should be marked as voted for proposal', async () => {
=======
    it("depositProposal shouldn't be created if it doesn't have Inactive or Denied status", async () => {
        await truffleAssert.passes(BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        ));

        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    it('originChainRelayerAddress should be marked as vote for proposal', async () => {
>>>>>>> master
        await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );
<<<<<<< HEAD
        const hasVoted = await BridgeInstance.hasVoted(
            originChainID, OriginERC20HandlerInstance.address,
            expectedDepositID, originChainValidatorAddress);
=======
        const hasVoted = await BridgeInstance.hasVoted(originChainID, expectedDepositID, originChainRelayerAddress);
>>>>>>> master
        assert.isTrue(hasVoted);
    });

    it('DepositProposalCreated event  should be emitted with expected values', async () => {
        const proposalTx = await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );

        TruffleAssert.eventEmitted(proposalTx, 'DepositProposalCreated', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.originChainHandlerAddress === OriginERC20HandlerInstance.address &&
                event.depositID.toNumber() === expectedDepositID &&
                event.dataHash === dataHash
        });
    });
<<<<<<< HEAD
});

contract('Bridge - [createDepositProposal with validatorThreshold > 1]', async (accounts) => {
    // const minterAndValidator = accounts[0];
    const originChainValidatorAddress = accounts[1];
    const depositerAddress = accounts[2];
=======

    it('getDepositProposal should return correct values in expected order', async () => {
        const expectedDepositProposal = {
            _originChainID: originChainID,
            _depositID: expectedDepositID,
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 'passed'
        };

        await BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );

        const depositProposal = await BridgeInstance.getDepositProposal(destinationChainID, expectedDepositID);
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
>>>>>>> master
    const destinationChainRecipientAddress = accounts[3];
    const originChainID = 0;
    const destinationChainID = 0;
    const originChainInitialTokenAmount = 100;
    const depositAmount = 10;
    const expectedDepositID = 1;
    const validatorThreshold = 2;

    let RelayerInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
<<<<<<< HEAD
        ValidatorInstance = await ValidatorContract.new([originChainValidatorAddress], validatorThreshold);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, validatorThreshold);
=======
        RelayerInstance = await RelayerContract.new([originChainRelayerAddress], 1);
        BridgeInstance = await BridgeContract.new(RelayerInstance.address, 2);
>>>>>>> master
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

    it('should create depositProposal successfully', async () => {
        TruffleAssert.passes(await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        ));
    });

    it('should revert because depositerAddress is not a validator', async () => {
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: depositerAddress }
        ));
    });

    it("depositProposal shouldn't be created if it has an Active status", async () => {
        await TruffleAssert.passes(BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        ));

<<<<<<< HEAD
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
=======
    it('Only relayers should be able to create a deposit proposal', async () => {
        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
>>>>>>> master
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        ));
    });

    xit("depositProposal shouldn't be created if it has a Passed status", async () => {
        
    });

    xit("depositProposal shouldn't be created if it has a Transferred status", async () => {
        
    });

    it('depositProposal should be created with expected values', async () => {
        const expectedDepositProposal = {
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 1 // active
        };

        await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );

        const depositProposal = await BridgeInstance._depositProposals.call(
            originChainID, OriginERC20HandlerInstance.address, expectedDepositID);
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
    it('originChainValidatorAddress should be marked as voted for proposal', async () => {
=======
    it("depositProposal shouldn't be created if it doesn't have Inactive or Denied status", async () => {
        await truffleAssert.passes(BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        ));

        await truffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        ));
    });

    it('originChainRelayerAddress should be marked as voted for proposal', async () => {
>>>>>>> master
        await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );
<<<<<<< HEAD
        const hasVoted = await BridgeInstance.hasVoted(
            originChainID, OriginERC20HandlerInstance.address,
            expectedDepositID, originChainValidatorAddress);
=======
        const hasVoted = await BridgeInstance.hasVoted(originChainID, expectedDepositID, originChainRelayerAddress);
>>>>>>> master
        assert.isTrue(hasVoted);
    });

    it('DepositProposalCreated event  should be emitted with expected values', async () => {
        const proposalTx = await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );

        TruffleAssert.eventEmitted(proposalTx, 'DepositProposalCreated', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.originChainHandlerAddress === OriginERC20HandlerInstance.address &&
                event.depositID.toNumber() === expectedDepositID &&
                event.dataHash === dataHash
        });
    });
<<<<<<< HEAD
=======

    it('getDepositProposal should return correct values in expected order', async () => {
        const expectedDepositProposal = {
            _originChainID: originChainID,
            _depositID: expectedDepositID,
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 'active'
        };

        await BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainRelayerAddress }
        );

        const depositProposal = await BridgeInstance.getDepositProposal(destinationChainID, expectedDepositID);
        const depositProposalValues = Object.values(depositProposal);
        depositProposalValues.forEach((depositRecordValue, index) => {
            depositProposalValues[index] = depositRecordValue.toNumber !== undefined ?
                depositRecordValue.toNumber() : depositRecordValue;
        });
        assert.sameOrderedMembers(depositProposalValues, Object.values(expectedDepositProposal));
    });
>>>>>>> master
});
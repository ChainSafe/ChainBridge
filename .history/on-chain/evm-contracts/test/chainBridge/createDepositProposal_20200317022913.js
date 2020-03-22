/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const TruffleAssert = require('truffle-assertions');
const Ethers = require('ethers');

const ValidatorContract = artifacts.require("Validator");
const BridgeContract = artifacts.require("Bridge");
const ERC20MintableContract = artifacts.require("ERC20Mintable");
const ERC20HandlerContract = artifacts.require("ERC20Handler");

contract('Bridge - [createDepositProposal with validatorThreshold = 1]', async (accounts) => {
    const originChainValidatorAddress = accounts[1];
    const depositerAddress = accounts[2];
    const originChainID = 0;
    const depositAmount = 10;
    const expectedDepositID = 1;

    let ValidatorInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([originChainValidatorAddress], 1);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, 1);
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
            { from: originChainValidatorAddress }
        ));

        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
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
            { from: originChainValidatorAddress }
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

    it('originChainValidatorAddress should be marked as voted for proposal', async () => {
        await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        );
        const hasVoted = await BridgeInstance.hasVoted(
            originChainID, OriginERC20HandlerInstance.address,
            expectedDepositID, originChainValidatorAddress);
        assert.isTrue(hasVoted);
    });

    it('DepositProposalCreated event  should be emitted with expected values', async () => {
        const proposalTx = await BridgeInstance.createDepositProposal(
            originChainID,
            OriginERC20HandlerInstance.address,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        );

        TruffleAssert.eventEmitted(proposalTx, 'DepositProposalCreated', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.originChainHandlerAddress === OriginERC20HandlerInstance.address &&
                event.depositID.toNumber() === expectedDepositID &&
                event.dataHash === dataHash
        });
    });
});

contract('Bridge - [createDepositProposal with validatorThreshold > 1]', async (accounts) => {
    // const minterAndValidator = accounts[0];
    const originChainValidatorAddress = accounts[1];
    const depositerAddress = accounts[2];
    const destinationChainRecipientAddress = accounts[3];
    const originChainID = 0;
    const destinationChainID = 0;
    const originChainInitialTokenAmount = 100;
    const depositAmount = 10;
    const expectedDepositID = 1;

    let ValidatorInstance;
    let BridgeInstance;
    let OriginERC20MintableInstance;
    let OriginERC20HandlerInstance;
    let DestinationERC20MintableInstance;
    let DestinationERC20HandlerInstance;
    let data = '';
    let dataHash = '';

    beforeEach(async () => {
        ValidatorInstance = await ValidatorContract.new([originChainValidatorAddress], 1);
        BridgeInstance = await BridgeContract.new(ValidatorInstance.address, 2);
        OriginERC20MintableInstance = await ERC20MintableContract.new();
        OriginERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);
        DestinationERC20MintableInstance = await ERC20MintableContract.new();
        DestinationERC20HandlerInstance = await ERC20HandlerContract.new(BridgeInstance.address);

        await OriginERC20MintableInstance.mint(depositerAddress, originChainInitialTokenAmount);
        await OriginERC20MintableInstance.approve(OriginERC20HandlerInstance.address, depositAmount, { from: depositerAddress });
        await BridgeInstance.depositERC20(
            OriginERC20MintableInstance.address,
            OriginERC20HandlerInstance.address,
            destinationChainID,
            DestinationERC20HandlerInstance.address,
            destinationChainRecipientAddress,
            depositAmount,
            { from: depositerAddress }
        );

        data = '0x' + Ethers.utils.hexZeroPad(DestinationERC20MintableInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(DestinationERC20HandlerInstance.address, 32).substr(2) +
            Ethers.utils.hexZeroPad(Ethers.utils.hexlify(depositAmount), 32).substr(2);
        dataHash = Ethers.utils.keccak256(data);
    });

    it('[sanity] ERC20 deposit record is created with expected depositID and values', async () => {
        const expectedDepositRecord = {
            _originChainTokenAddress: OriginERC20MintableInstance.address,
            _originChainHandlerAddress: OriginERC20HandlerInstance.address,
            _destinationChainID: destinationChainID,
            _destinationChainHandlerAddress: DestinationERC20HandlerInstance.address,
            _destinationRecipientAddress: destinationChainRecipientAddress,
            _amount: depositAmount
        };

        const depositRecord = await BridgeInstance._erc20DepositRecords.call(destinationChainID, expectedDepositID);
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
        TruffleAssert.passes(await BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        ));
    });

    it('Only validators should be able to create a deposit proposal', async () => {
        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: depositerAddress }
        ));
    });

    it('depositProposal is created with expected values', async () => {
        const expectedDepositProposal = {
            _originChainID: originChainID,
            _depositID: expectedDepositID,
            _dataHash: dataHash,
            _numYes: 1,
            _numNo: 0,
            _status: 1 // passed
        };

        await BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        );

        const depositProposal = await BridgeInstance._depositProposals.call(originChainID, expectedDepositID);
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
        await TruffleAssert.passes(BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        ));

        await TruffleAssert.reverts(BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        ));
    });

    it('originChainValidatorAddress should be marked as voted for proposal', async () => {
        await BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        );
        const hasVoted = await BridgeInstance.hasVoted(originChainID, expectedDepositID, originChainValidatorAddress);
        assert.isTrue(hasVoted);
    });

    it('DepositProposalCreated event is fired with expected value', async () => {
        const proposalTx = await BridgeInstance.createDepositProposal(
            originChainID,
            expectedDepositID,
            dataHash,
            { from: originChainValidatorAddress }
        );

        TruffleAssert.eventEmitted(proposalTx, 'DepositProposalCreated', (event) => {
            return event.originChainID.toNumber() === originChainID &&
                event.depositID.toNumber() === expectedDepositID &&
                event.dataHash === dataHash
        });
    });
});
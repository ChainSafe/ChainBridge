/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ethers = require('ethers');
const constants = require('./../constants');

const BridgeContract = require("../../../build/contracts/Bridge.json");
const ValidatorContract = require("../../../build/contracts/Validator.json");
const ERC20HandlerContract = require("../../../build/contracts/ERC20Handler.json");
const ERC20MintableContract = require("../../../build/contracts/ERC20Mintable.json");

// old
const TestEmitterContract = require("../../../build/contracts/SimpleEmitter.json");
const CentrifugeContract = require("../../../build/contracts/BridgeAsset.json");
const ERC20Contract = require("../../../build/contracts/ERC20Mintable.json");
const ERC721Contract = require("../../../build/contracts/ERC721Mintable.json");

async function deployValidatorContract(cfg) {
    // Create an instance of a Contract Factory
    let factory = new ethers.ContractFactory(ValidatorContract.abi, ValidatorContract.bytecode, cfg.mainWallet);

    // Set validators
    const validators = constants.validatorAddresses.slice(0, cfg.numValidators);

    // Deploy
    let contract = await factory.deploy(
        validators,
        cfg.validatorThreshold
    );
    console.log("[Validator] Contract address: ", contract.address);
    console.log("[Validator] Transaction Hash: ", contract.deployTransaction.hash);
    await contract.deployed();
}

async function deployBridgeContract(cfg) {
    try {
        // Create an instance of a Contract Factory
        let factory = new ethers.ContractFactory(BridgeContract.abi, BridgeContract.bytecode, cfg.mainWallet);

        // Deploy
        let contract = await factory.deploy(
            constants.VALIDATOR_ADDRESS,
            1
        );

        console.log("[Bridge] Contract address: ", contract.address);
        console.log("[Bridge] Transaction Hash: ", contract.deployTransaction.hash);
        await contract.deployed();
    } catch (e) {
        console.log({ e });
    }
}

async function deployERC20Handler(cfg) {
    const handlerFactory = new ethers.ContractFactory(ERC20HandlerContract.abi, ERC20HandlerContract.bytecode, cfg.mainWallet);
    const erc20MintableFactory = new ethers.ContractFactory(ERC20MintableContract.abi, ERC20MintableContract.bytecode, cfg.mainWallet);
    const handlerContract = await handlerFactory.deploy(constants.BRIDGE_ADDRESS);
    const erc20MintableContract = await erc20MintableFactory.deploy();

    console.log("[ERC20 Handler] Contract address: ", handlerContract.address);
    console.log("[ERC20 Handler] Transaction Hash: ", handlerContract.deployTransaction.hash);

    console.log("[ERC20 Token] Contract address: ", erc20MintableContract.address);
    console.log("[ERC20 Token] Transaction Hash: ", erc20MintableContract.deployTransaction.hash);
}

// Deployment is asynchronous, so we use an async IIFE
async function deployCentrifuge(cfg) {
    try {
        // Create an instance of a Contract Factory
        let factory = new ethers.ContractFactory(CentrifugeContract.abi, CentrifugeContract.bytecode, cfg.mainWallet);
        let contract = await factory.deploy(
            10
        );

        console.log("[Centrifuge] Contract address: ", contract.address);
        console.log("[Centrifuge] Transaction Hash: ", contract.deployTransaction.hash);
        await contract.deployed();
    } catch (e) {
        console.log({ e });
    }
};

// Deployment is asynchronous, so we use an async IIFE
async function deployEmitterTest(cfg) {

    // Create an instance of a Contract Factory
    let factory = new ethers.ContractFactory(TestEmitterContract.abi, TestEmitterContract.bytecode, cfg.mainWallet);
    let contract = await factory.deploy();

    console.log("[Test Emitter] Contract address: ", contract.address);
    console.log("[Test Emitter] Transaction Hash: ", contract.deployTransaction.hash);
    await contract.deployed()
};

module.exports = {
    deployValidatorContract,
    deployBridgeContract,
    deployERC20Handler,
    // old
    // deployCentrifuge,
    // deployEmitterTest,
    // deployERC20,
}
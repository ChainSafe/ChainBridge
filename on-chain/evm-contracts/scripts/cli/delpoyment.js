/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ethers = require('ethers');
const constants = require('./constants');

const ReceiverContract = require("../../build/contracts/Receiver.json");
const EmitterContract = require("../../build/contracts/Emitter.json");
const TestEmitterContract = require("../../build/contracts/SimpleEmitter.json");
const CentrifugeContract = require("../../build/contracts/BridgeAsset.json");
const ERC20Contract = require("../../build/contracts/ERC20Mintable.json");
const ERC721Contract = require("../../build/contracts/ERC721Mintable.json");

async function deployReceiver(cfg) {
    try {
        // Create an instance of a Contract Factory
        let factory = new ethers.ContractFactory(ReceiverContract.abi, ReceiverContract.bytecode, cfg.mainWallet);

        // Set relayers
        const relayers = constants.relayerAddresses.slice(0, cfg.numRelayers);

        // Deploy
        let contract = await factory.deploy(
            relayers,
            cfg.depositThreshold,
            cfg.relayerThreshold
        );

        console.log("[Receiver] Contract address: ", contract.address);
        console.log("[Receiver] Transaction Hash: ", contract.deployTransaction.hash);
        await contract.deployed();

        // Test it worked correctly
        let ReceiverInstance = new ethers.Contract(contract.address, ReceiverContract.abi, cfg.provider);
        // Ensure relayers are set correctly
        // note relayers[] is sorted in order, so we'll check the top arrays
        let failure = false;
        relayers.forEach(async (v, i) => {
            const value = await ReceiverInstance.Relayers(v);
            if (value) {
                const wei = await cfg.provider.getBalance(v);
                const balance = ethers.utils.formatEther(wei);
                console.log(`[Receiver] ${v} (${balance} ETH) is a relayer!`);
            } else {
                console.log(`ERROR: ${v} was not set as a relayer!`);
                failure = true;
            }
        });

        if (failure) {
            process.exit();
        }
    } catch (e) {
        console.log({ e });
    }
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
async function deployEmitter(cfg) {
    try {
        // Create an instance of a Contract Factory
        let factory = new ethers.ContractFactory(EmitterContract.abi, EmitterContract.bytecode, cfg.mainWallet);
        let contract = await factory.deploy();

        console.log("[Emitter] Contract address: ", contract.address);
        console.log("[Emitter] Transaction Hash: ", contract.deployTransaction.hash);
        await contract.deployed();
    } finally {

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

async function deployERC20(cfg) {
    const minterWallet = new ethers.Wallet(constants.relayerPrivKeys[0], cfg.provider);
    let tokenFactory = new ethers.ContractFactory(ERC20Contract.abi, ERC20Contract.bytecode, minterWallet);
    const tokenContract = await tokenFactory.deploy();
    const contract = await tokenContract.deployed();
    // The address the Contract WILL have once mined
    console.log("[Deploy ERC20] Contract address: ", contract.address);

    // The transaction that was sent to the network to deploy the Contract
    console.log("[Deploy ERC20] Transaction Hash: ", contract.deployTransaction.hash);
    await contract.mint(minterWallet.address, 100);
    console.log("[ERC20 Transfer] Minted tokens!");
    console.log("[ERC20 Transfer] Token address:", contract.address);
}

module.exports = {
    deployReceiver,
    deployCentrifuge,
    deployEmitter,
    deployEmitterTest,
    deployERC20,
}
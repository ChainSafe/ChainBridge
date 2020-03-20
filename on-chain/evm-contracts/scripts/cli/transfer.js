/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ethers = require('ethers');
const constants = require('./constants');

const ReceiverContract = require("../../build/contracts/Receiver.json");
const EmitterContract = require("../../build/contracts/Emitter.json");
const TestEmitterContract = require("../../build/contracts/SimpleEmitter.json");
const ERC20Contract = require("../../build/contracts/ERC20Mintable.json");
const ERC721Contract = require("../../build/contracts/ERC721Mintable.json");

async function assetTestTransfer(cfg) {
    try {
        const deployerWallet = new ethers.Wallet(constants.relayerPrivKeys[0], cfg.provider);
        let emitterInstance = new ethers.Contract(constants.TEST_EMITTER_ADDRESS, TestEmitterContract.abi, deployerWallet);
        // Trigger fallback
        const tx = await cfg.mainWallet.sendTransaction({
            to: emitterInstance.address,
            value: ethers.utils.parseEther("0.0")
        });
        console.log("[Deploy Asset] Tx hash: ", tx.hash);
    } catch (e) {
        console.log({ e })
    }
}

async function erc20Transfer(cfg) {
    try {
        const minterWallet = new ethers.Wallet(constants.relayerPrivKeys[0], cfg.provider);

        // Instances
        const erc20Instance = new ethers.Contract(constants.ERC20_ADDRESS, ERC20Contract.abi, minterWallet);
        const emitterInstance = new ethers.Contract(constants.EMITTER_ADDRESS, EmitterContract.abi, minterWallet);
        const receiverInstance = new ethers.Contract(constants.RECEIVER_ADDRESS, ReceiverContract.abi, minterWallet);
        if (cfg.watchMode) {
            watchBalances(cfg, erc20Instance, receiverInstance, emitterInstance.address, minterWallet.address, constants.relayerAddresses[1])
        } else {

            // Mint & Approve tokens
            await erc20Instance.approve(constants.EMITTER_ADDRESS, 1);
            console.log("[ERC20 Transfer] Approved tokens!");

            // Check the balance before the transfer
            const prebal = await emitterInstance.balances(erc20Instance.address);
            console.log("[ERC20 Transfer] Pre token balaance: ", prebal.toNumber());

            // Make the deposit
            await emitterInstance.depositGenericErc(cfg.dest, cfg.value, constants.relayerAddresses[1], erc20Instance.address);
            console.log("[ERC20 Transfer] Created deposit!");

            // Check the balance after the deposit
            const postbal = await emitterInstance.balances(erc20Instance.address);
            console.log("[ERC20 Transfer] Post token balaance: ", postbal.toNumber());
        }
    } catch (e) {
        console.log({ e });
    }
}

async function erc721Transfer(cfg) {
    try {
        console.log("[ERC721 Transfer] EMITTER_ADDRESS:", constants.EMITTER_ADDRESS);
        const minterWallet = new ethers.Wallet(constants.relayerPrivKeys[0], cfg.provider);

        // Create token
        let tokenFactory = new ethers.ContractFactory(ERC721Contract.abi, ERC721Contract.bytecode, minterWallet);
        const tokenContract = await tokenFactory.deploy();
        await tokenContract.deployed();
        console.log("[ERC721 Transfer] Deployed token!")

        // Mint tokens
        let erc721Instance = new ethers.Contract(tokenContract.address, ERC721Contract.abi, minterWallet);
        await erc721Instance.mint(minterWallet.address, 1);
        console.log("[ERC721 Transfer] Minted tokens!");

        // Approve tokens
        await erc721Instance.approve(constants.EMITTER_ADDRESS, 1);
        console.log("[ERC721 Transfer] Approved tokens!");

        // Create emitter instance
        const emitterInstance = new ethers.Contract(constants.EMITTER_ADDRESS, EmitterContract.abi, minterWallet);

        // Check pre balance
        const prebal = await erc721Instance.balanceOf(constants.EMITTER_ADDRESS);
        console.log("[ERC721 Transfer] Pre balance:", prebal.toNumber());

        // Check the owner
        let owner = await erc721Instance.ownerOf(1);
        console.log("[ERC721 Transfer] Owner of token 1:", owner);

        emitterInstance.on("NFTTransfer", (dest, deposit, to, token, data) => {
            console.log(`
            Dest ${dest.toString()}
            Depositd_Id ${deposit.toString()}
            To ${to}
            Token ${token}
            Data ${data.toString()}
        `)
            process.exit()
        })

        // // Perform deposit
        const d = await emitterInstance.depositNFT(cfg.dest, constants.relayerAddresses[1], erc721Instance.address, 1, "0x");
        console.log("[ERC721 Transfer] Created deposit!")
        console.log("[ERC721 Transfer] Deposit Hash", d.hash);
    } catch (e) {
        console.log({ e });
    }
}

async function depositTest(cfg) {
    try {
        console.log("[Deposit Test] RECEIVER ADDRESS:", constants.RECEIVER_ADDRESS);
        const minterWallet = new ethers.Wallet(constants.relayerPrivKeys[0], cfg.provider);

        // Create receiver instance
        const receiverInstance = new ethers.Contract(constants.RECEIVER_ADDRESS, ReceiverContract.abi, minterWallet);

        console.log("signature", receiverInstance.interface.events.DepositProposalCreated.signature)
        console.log("Topic", receiverInstance.interface.events.DepositProposalCreated.topic)

        // // Perform deposit
        await receiverInstance.createDepositProposal(
            ethers.utils.formatBytes32String("HASH"),
            Math.floor(Math.random() * 10000),
            Math.floor(Math.random() * 10000)
        );
    } catch (e) {
        console.log({ e });
    }
}

async function watchBalances(cfg, tokenInstance, receiverInstance, bridge, from, to) {
    let depositCount = 0;
    let bridgeBal = await tokenInstance.balanceOf(bridge)
    let fromBal = await tokenInstance.balanceOf(from)
    let toBal = await tokenInstance.balanceOf(to)
    let proposal = await receiverInstance.getDepositProposal(0, depositCount)
    let curStatus, curOrigin, curHash, threshold;
    receiverInstance.on("DepositProposalCreated", (hash, count, origin, status) => {
        curStatus = status
        curHash = hash
        curOrigin = origin
        depositCount = count.toNumber()
    })
    receiverInstance.on("DepositProposalVote", (origin, count, _, status) => {
        curOrigin = origin
        depositCount = count.toNumber()
        curStatus = status
    })
    displayLog({
        bridgeBal,
        fromBal,
        toBal,
        proposal,
        curHash,
        curOrigin,
        curStatus,
        threshold,
        depositCount,
        port: cfg.port,
    })
    setInterval(async () => {
        // depositCount += 1
        proposal = await receiverInstance.getDepositProposal(0, depositCount)
        bridgeBal = await tokenInstance.balanceOf(bridge)
        fromBal = await tokenInstance.balanceOf(from)
        toBal = await tokenInstance.balanceOf(to)
        receiverX = await receiverInstance.getDepositProposal(0, depositCount)
        threshold = await receiverInstance.DepositThreshold()
        displayLog({
            bridgeBal,
            fromBal,
            toBal,
            proposal,
            curHash,
            curOrigin,
            curStatus,
            threshold,
            depositCount,
            port: cfg.port,
        })
    }, 1000)
}

function displayLog(vals) {
    console.log(`
        Port:             ${vals.port}
        Bridge balance:   ${vals.bridgeBal}
        Sender balance:   ${vals.fromBal}
        Receiver balance: ${vals.toBal}
        Threshold:        ${vals.threshold}
        ===================
        Latest Deposit Info
        ===================
        Deposit Info
        Deposit Count: ${vals.depositCount}
        Orign Chain:   ${vals.curOrigin}
        Vote Status:   ${VoteStatus(vals.curStatus)}
        Deposit Hash:  ${vals.curHash}
        Number Yes:    ${vals.proposal.numYes}
        Number No:     ${vals.proposal.numNo}
        `)
}

function VoteStatus(code) {
    switch (code) {
        case 0:
            return "Inactive";
        case 1:
            return "Active";
        case 2:
            return "Finalized";
        case 3:
            return "Transferred";
    }
}

module.exports = {
    assetTestTransfer,
    erc20Transfer,
    erc721Transfer,
    depositTest,
}
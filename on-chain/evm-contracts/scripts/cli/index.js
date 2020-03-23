/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ethers = require('ethers');
const cli = require('commander');

const deploy = require('./cmd/deployment');
const trasnfer = require('./cmd/transfer');
const constants = require('./constants');

// Capture argument
cli
    .option('--relayers <value>', 'Number of relayers', 2)
    .option('-v, --relayer-threshold <value>', 'Value of relayer threshold', 1)
    .option('-d, --deposit-threshold <value>', 'Value of deposit threshold', 1)
    .option('-p, --port <value>', 'Port of RPC instance', 8545)
    .option('--deposit-erc', "Make an ERC20 deposit", false)
    .option('--deposit-nft', "Make an ERC721 deposit", false)
    .option('--deposit-asset', "Make a test deployment", false)
    .option('--deposit-test', "Make a deposit test", false)
    .option('--test-only', "Skip main contract depoyments, only run tests", false)
    .option('--dest <value>', "destination chain", 1)
    .option('--watchMode', "Watch contracts", false)
    .option('--value <amount>', "Amount to transfer", 1)
    .option('--mint-erc20', "Mint erc20 based on --value", false)
cli.parse(process.argv);

// Connect to the network
cli.url = `http://${process.env.BASE_URL || "localhost"}:${cli.port}`;
cli.provider = new ethers.providers.JsonRpcProvider(cli.url);
// Only support up to 10 in this setup
cli.numRelayers = cli.relayers > 10 ? 10 : cli.relayers;

if (cli.relayerThreshold <= cli.numRelayers) {
    cli.relayerThreshold = cli.numRelayers;
}
if (cli.depositThreshold <= cli.numRelayers) {
    cli.depositThreshold = cli.numRelayers;
}
if (cli.dest) {
    cli.dest = Number(cli.dest);
}
if (cli.value) {
    cli.value = Number(cli.value);
}

// Load the wallet to deploy the contract with
cli.mainWallet = new ethers.Wallet(constants.deployerPrivKey, cli.provider);

// Deployment is asynchronous, so we use an async IIFE
(async function () {
    if (!cli.testOnly) {
        await deploy.deployRelayerContract(cli);
        await deploy.deployBridgeContract(cli);
        await deploy.deployERC20Handler(cli);

        //old
        // await deploy.deployCentrifuge(cli);
        // await deploy.deployEmitterTest(cli);
    }

    if (cli.depositErc) {
        await trasnfer.erc20Transfer(cli);
    } else if (cli.mintErc20) {
        await trasnfer.mintErc20(cli);
    } else if (cli.depositNft) {
        await trasnfer.erc721Transfer(cli);
    } else if (cli.depositTest) {
        await trasnfer.depositTest(cli);
    } else if (cli.depositAsset) {
        await trasnfer.assetTestTransfer(cli);
    }
})();
const ethers = require('ethers');
const ReceiverContract = require("../build/contracts/Receiver.json");
const EmitterContract = require("../build/contracts/Emitter.json");
const TestEmitterContract = require("../build/contracts/SimpleEmitter.json");
const CentrifugeContract = require("../build/contracts/BridgeAsset.json");
const cli = require('commander');

// Capture argument
cli
    .option('--validators <value>', 'Number of validators', 2)
    .option('-v, --validator-threshold <value>', 'Value of validator threshold', 2)
    .option('-d, --deposit-threshold <value>', 'Value of deposit threshold', 1)
    .option('-p, --port <value>', 'Port of RPC instance', 8545)
cli.parse(process.argv);

// Connect to the network
console.log("ENV", process.env.BASE_URL);
let url = `http://${process.env.BASE_URL || "localhost"}:${cli.port}`;
let provider = new ethers.providers.JsonRpcProvider(url);

// Config
let numValidators = 2;
let validatorThreshold = 2;
let depositThreshold = 1;

// CLI args position
// 0 = number of validators
// 1 = validator threshold
// 2 = deposit threshold
if (cli.validators) {
    // Only support up to 10 in this setup
    numValidators = cli.validators > 10 ? 10 : cli.validators;
}
if (cli.validatorThreshold && cli.validatorThreshold <= numValidators) {
    validatorThreshold = cli.validatorThreshold;
}
if (cli.depositThreshold && cli.depositThreshold <= numValidators) {
    depositThreshold = cli.depositThreshold;
}

// Keys generate from: when sound uniform light fee face forum huge impact talent exhaust arrow
const deployerPubKey = "0x34c59fBf82C9e31BA9CBB5faF4fe6df05de18Ad4";
const deployerPrivKey = "0x39a9ea0dce63086c64a80ce045b796bebed2006554e3992d92601515c7b19807";


// TODO Remove these and cycle through mnemonic
const validatorPubkeys = [
    "0x029b67ec8aba36421137e22d874a897f8aa2a47e2d479d772d96ca8c5744b5a95c",
    "0x0256e9cd885e02497bedbd9b780d9ff05fba77fab5c757f35530461681879206ee",
    "0x03237d70c85198221508cefe652bef49e0afaebf738044aabbf86a91f10f5da73b",
    "0x03de552cb89aea6ce30096503c4f4d098d0457d4022258cc287e6a6cf34d0ad70a",
    "0x02da378cfd22df34e4e89df63b4ed279a426992112601ae9fc538055bd55d8a7a5",
]
const validatorPrivKeys = [
    "0x000000000000000000000000000000000000000000000000000000416c696365",
    "0x0000000000000000000000000000000000000000000000000000000000426f62",
    "0x00000000000000000000000000000000000000000000000000436861726c6965",
    "0x0000000000000000000000000000000000000000000000000000000044617665",
    "0x0000000000000000000000000000000000000000000000000000000000457665",
]

// Load the wallet to deploy the contract with
let wallet = new ethers.Wallet(deployerPrivKey, provider);
// Deployment is asynchronous, so we use an async IIFE
(async function () {
    try {
        // Create an instance of a Contract Factory
        let factory = new ethers.ContractFactory(ReceiverContract.abi, ReceiverContract.bytecode, wallet);

        // Set validators
        const validators = validatorPubkeys.slice(0, numValidators);

        // Deploy
        let contract = await factory.deploy(
            validators,
            depositThreshold,
            validatorThreshold
        );
        console.log("here");

        // The address the Contract WILL have once mined
        console.log("[Receiver] Contract address: ", contract.address);

        // The transaction that was sent to the network to deploy the Contract
        console.log("[Receiver] Transaction Hash: ", contract.deployTransaction.hash);
        // The contract is NOT deployed yet; we must wait until it is mined
        await contract.deployed();
        // Done! The contract is deployed.

        // Test it worked correctly
        let ReceiverInstance = new ethers.Contract(contract.address, ReceiverContract.abi, provider);

        // Ensure validators are set correctly
        // note validators[] is sorted in order, so we'll check the top arrays
        let failure = false;
        validators.forEach(async (v, i) => {
            const value = await ReceiverInstance.Validators(v);
            if (value) {
                const wei = await provider.getBalance(v);
                const balance = ethers.utils.formatEther(wei);
                console.log(`${v} (${balance} ETH) is a validator!`);
            } else {
                console.log(`ERROR: ${v} was not set as a validator!`);
                failure = true;
            }
        });

        if (failure) {
            process.exit();
        }

        await deployCentrifuge();
        await deployEmitter();
    } catch (e) {
        console.log({ e });
    }

    await deployCentrifuge();
    await deployEmitter();
    await deployEmitterTest();
})();

// Deployment is asynchronous, so we use an async IIFE
async function deployCentrifuge() {
    try {
        // Create an instance of a Contract Factory
        let factory = new ethers.ContractFactory(CentrifugeContract.abi, CentrifugeContract.bytecode, wallet);

        // Deploy
        let contract = await factory.deploy(
            10
        );

        // The address the Contract WILL have once mined
        console.log("[Centrifuge] Contract address: ", contract.address);

        // The transaction that was sent to the network to deploy the Contract
        console.log("[Centrifuge] Transaction Hash: ", contract.deployTransaction.hash);

        // The contract is NOT deployed yet; we must wait until it is mined
        await contract.deployed();
        // Done! The contract is deployed.

        let CentrifugeInstance = new ethers.Contract(contract.address, CentrifugeContract.abi, provider);
    } catch (e) {
        console.log({ e });
    }
};

// Deployment is asynchronous, so we use an async IIFE
async function deployEmitter() {
    try {
        // Create an instance of a Contract Factory
        let factory = new ethers.ContractFactory(EmitterContract.abi, EmitterContract.bytecode, wallet);

        // Deploy
        let contract = await factory.deploy();

        // The address the Contract WILL have once mined
        console.log("[Emitter] Contract address: ", contract.address);

        // The transaction that was sent to the network to deploy the Contract
        console.log("[Emitter] Transaction Hash: ", contract.deployTransaction.hash);

        // The contract is NOT deployed yet; we must wait until it is mined
        await contract.deployed();
        // Done! The contract is deployed.

        let EmitterInstance = new ethers.Contract(contract.address, CentrifugeContract.abi, provider);
    } finally {
        
    }
};

// Deployment is asynchronous, so we use an async IIFE
async function deployEmitterTest() {

    // Create an instance of a Contract Factory
    let factory = new ethers.ContractFactory(TestEmitterContract.abi, TestEmitterContract.bytecode, wallet);

    // Deploy
    let contract = await factory.deploy();

    // The address the Contract WILL have once mined
    console.log("[Test Emitter] Contract address: ", contract.address);

    // The transaction that was sent to the network to deploy the Contract
    console.log("[Test Emitter] Transaction Hash: ", contract.deployTransaction.hash);

    // The contract is NOT deployed yet; we must wait until it is mined
    await contract.deployed()
    // Done! The contract is deployed.

    let EmitterInstance = new ethers.Contract(contract.address, CentrifugeContract.abi, provider);
};

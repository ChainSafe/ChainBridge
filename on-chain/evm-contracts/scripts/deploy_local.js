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
const deployerAddress = "0x0c6CD6Dc5258EF556eA7c6dab2abE302fB60e0b6";
const deployerPrivKey = "0x000000000000000000000000000000000000000000000000000000416c696365";


// TODO Remove these and cycle through mnemonic
const validatorAddress = [
    "0x0c6CD6Dc5258EF556eA7c6dab2abE302fB60e0b6",
    "0x0E17A926c6525b59921846c85E1efD7a5396a47B",
    "0x0f05849291a309EC001bbd2dAd7DC6F989c40c80",
    "0x3003d03276434dd23429D4D33090DA5948b7b510",
    "0x251e6F841549D519dE6De1e99241695bc1000A26",
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
        const validators = validatorAddress.slice(0, numValidators);

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

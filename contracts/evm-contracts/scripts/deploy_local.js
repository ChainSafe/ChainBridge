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
    .option('-d, --deposit-threshold <value>', 'Value of deposit threshold', 2)
    .option('-p, --port <value>', 'Port of RPC instance', 8545)
cli.parse(process.argv);

// Connect to the network
let url = `http://localhost:${cli.port}`;
let provider = new ethers.providers.JsonRpcProvider(url);

// Config
let numValidators = 2;
let validatorThreshold = 2;
let depositThreshold = 2;

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
    "0x34c59fBf82C9e31BA9CBB5faF4fe6df05de18Ad4",
    "0x0a4c3620AF8f3F182e203609f90f7133e018Bf5D",
    "0x75A8e662d70904e6f5f5cABed8f874f3cE5714f1",
    "0x7FdAFcCa9A3e69C483D88BF23126b5B4021f7E24",
    "0xd695b3C074aD73EbEE55B71A28497dA6b2894307",
    "0x57dd6d89bBb594911645C19149D2f979CF09ebf1",
    "0x2F73cdf831FedfcFD257c795cb6aDD83BfEC722d",
    "0xEC33e3945A50eD81d69D2454Bcd21f87989ACf83",
    "0x16025F9d271283ba7cEAFe87d5031e17E427dA84",
    "0x87c2e6be98f1B88D8d3748220eA7bCb51393B69E"
]
const validatorPrivKeys = [
    "0x39a9ea0dce63086c64a80ce045b796bebed2006554e3992d92601515c7b19807",
    "0x5de3b6992e5ad40dc346cfd3b00595f58bd16ea38b43511e7a00a2a60d380225",
    "0x342254b80c2f78806850cd5d6a7e91b1f004271ea08714654aef28bb0c0b81f2",
    "0x4bb140bee3a48b88ebd2ebe5903e2ccfc3afadf4be373930875b93576b0c48f7",
    "0xec638dcf102e0b0213b498451c6311e5af0532bb2d8deff5ce526c38fb74519a",
    "0x34548c03095760bd48a027de381a48db562e1358df15d176dd1fd07063604841",
    "0xc1380998e9cb034c8c92d1d4bf7cfce647da0568de72775d00b8c9f3528875f8",
    "0x1bbf053a7b2ed73ad549437a175f5310373154f63e0993bd9b5a5aa013f10794",
    "0x6d77d1d80f38de6e6ec67f65c52b3f9f301571ef6e178aab3d732c72eae8b8c1",
    "0x9fd8879319af178930822a55e88ee9177958094c760b1a56b1362b4098574c5d"
] 

// Load the wallet to deploy the contract with
let wallet = new ethers.Wallet(deployerPrivKey, provider);

// Deployment is asynchronous, so we use an async IIFE
(async function () {

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

    // The address the Contract WILL have once mined
    console.log("[Receiver] Contract address: ", contract.address);

    // The transaction that was sent to the network to deploy the Contract
    console.log("[Receiver] Transaction Hash: ", contract.deployTransaction.hash);
    // The contract is NOT deployed yet; we must wait until it is mined
    await contract.deployed()
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
    await deployEmitterTest();
})();

// Deployment is asynchronous, so we use an async IIFE
async function deployCentrifuge () {

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
    await contract.deployed()
    // Done! The contract is deployed.

    let CentrifugeInstance = new ethers.Contract(contract.address, CentrifugeContract.abi, provider);
};

// Deployment is asynchronous, so we use an async IIFE
async function deployEmitter () {

    // Create an instance of a Contract Factory
    let factory = new ethers.ContractFactory(EmitterContract.abi, EmitterContract.bytecode, wallet);

    // Deploy
    let contract = await factory.deploy();

    // The address the Contract WILL have once mined
    console.log("[Emitter] Contract address: ", contract.address);

    // The transaction that was sent to the network to deploy the Contract
    console.log("[Emitter] Transaction Hash: ", contract.deployTransaction.hash);

    // The contract is NOT deployed yet; we must wait until it is mined
    await contract.deployed()
    // Done! The contract is deployed.

    let EmitterInstance = new ethers.Contract(contract.address, CentrifugeContract.abi, provider);
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

const ethers = require('ethers');
const ReceiverContract = require("../build/contracts/Receiver.json");
const cli = require('commander');

// Capture argument
cli
    .option('--validators <value>', 'Number of validators', 2)
    .option('-v, --validator-threshold <value>', 'Value of validator threshold', 2)
    .option('-d, --deposit-threshold <value>', 'Value of deposit threshold', 2);
cli.parse(process.argv);

// Connect to the network
let provider = new ethers.providers.JsonRpcProvider();

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
const deployerPubKey = "0xC5F737aE7EaBB7226f21121E335b0949d8eA6365";
const deployerPrivKey = "0xae6a8b4518e3970a0501ecf796a51dc0dab9143a66be75e948bf352582db15d5";

// TODO Remove these and cycle through mnemonic
const validatorPubkeys = [
    "0x0Cf5165Fabf32B514b89544E75CDf0dD359eCB3D",
    "0x683DD6A2A92E6a03f821d207D1CC2f815E5845cB",
    "0xa215F1d6334C8237c97b901fa7aA68dBc86c994e",
    "0x89347d4fFC69d472B37fA85A6eb4d3dB85282087",
    "0xD765355ba39b09bfAb25Ed73406Dd8C7fdc2fd18",
    "0xC8A0298eb0f3D1f3F1E8574674ffA22BBBe5F733",
    "0xb459ED56176EfacdD5a55CF07b9f94FEDa75F22A",
    "0x95A64A55BE87062674810142A57879F9fF102e38",
    "0x4A44784cBbeAb4E9F3DC4fE1c462D2967563ef4E",
];
const validatorPrivKeys = [
    "0x093ccd3d4b79ba0a4f54316ef82fb8c699ad09d589538a47f854900f83d53049",
    "0x7668bc519fc8e3155a6877239fc2b11db42e55363442b88185f1c6a04f58f64b",
    "0x249cb7f3c7d29c29dbaae32df23f23b33b74aa197713451572ef12a891fdeb2f",
    "0x9b97be90c59a6fc9a5408b1d9620d5e332783ac0391ec0bf4c681968f5a0d990",
    "0x624b171fc4aa1afc81cd54c349fa839053c2060a4bf263f38a443082023d53f7",
    "0x589b1d2c52cc8cbb1caf1abfb827a4c4303d0aff9593f66a53ad03dd54425a26",
    "0x597095465c6a64656ff7378a425741bbc5d151252f2163056c6480ecd23bccbc",
    "0x5917b2cc250c047013878d56a2df8c61bd7d382e72a2d87aa5cfd2a60c1d79b3",
    "0x9e0e1ae61566bc6717795a719b8e38acb91784e83bae3e7852ab0e72c05665a2",
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
    console.log("Contract address: ", contract.address);

    // The transaction that was sent to the network to deploy the Contract
    console.log("Transaction Hash: ", contract.deployTransaction.hash);

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
})();
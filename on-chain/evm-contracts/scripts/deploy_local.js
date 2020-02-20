const ethers = require('ethers');
const ReceiverContract = require("../build/contracts/Receiver.json");
const EmitterContract = require("../build/contracts/Emitter.json");
const TestEmitterContract = require("../build/contracts/SimpleEmitter.json");
const CentrifugeContract = require("../build/contracts/BridgeAsset.json");
const ERC20Contract = require("../build/contracts/ERC20Mintable.json");
const ERC721Contract = require("../build/contracts/ERC721Mintable.json");
const cli = require('commander');

// Capture argument
cli
    .option('--validators <value>', 'Number of validators', 2)
    .option('-v, --validator-threshold <value>', 'Value of validator threshold', 2)
    .option('-d, --deposit-threshold <value>', 'Value of deposit threshold', 1)
    .option('-p, --port <value>', 'Port of RPC instance', 8545)
    .option('--deposit-erc', "Make an ERC20 deposit", false)
    .option('--deposit-nft', "Make an ERC721 deposit", false)
    .option('--deposit-asset', "Make a test deployment", false)
    .option('--test-only', "Skip main contract depoyments, only run tests", false)
cli.parse(process.argv);

// Connect to the network
let url = `http://${process.env.BASE_URL || "localhost"}:${cli.port}`;
let provider = new ethers.providers.JsonRpcProvider(url);

// Config
let numValidators = 2;
let validatorThreshold = 2;
let depositThreshold = 1;

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

// These are deterministic
const RECEIVER_ADDRESS = "0x705D4Fa884AF2Ae59C7780A0f201109947E2Bf6D";
const CENTRIFUGE_ADDRESS = "0x290f41e61374c715C1127974bf08a3993afd0145";
const EMITTER_ADDRESS = "0x1fA38b0EfccA4228EB9e15112D4d98B0CEe3c600";
const TEST_EMITTER_ADDRESS = "0x70486404e42d17298c57b046Aa162Dc3aCc075f0";

// Deployment is asynchronous, so we use an async IIFE
(async function () {
    if (!cli.testOnly) {
        await deployReceiver();
        await deployCentrifuge();
        await deployEmitter();
        await deployEmitterTest();
    }

    if (cli.depositErc) {
        await erc20Transfer();
    } else if (cli.depositNft) {
        await erc721Transfer();
    } else if (cli.depositAsset) {
        await deployAssetTest();
    }
})();

async function deployReceiver() {
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
    } catch (e) {
        console.log({ e });
    }
}

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

        // let CentrifugeInstance = new ethers.Contract(contract.address, CentrifugeContract.abi, provider);
        // CENTRIFUGE_ADDRESS = contract.address;
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
        // EMITTER_ADDRESS = contract.address;
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

    // TEST_EMITTER_ADDRESS = contract.address;
};

async function deployAssetTest() {
    try {
        const deployerWallet = new ethers.Wallet(validatorPrivKeys[0], provider);
        let emitterInstance = new ethers.Contract(TEST_EMITTER_ADDRESS, TestEmitterContract.abi, deployerWallet);
        // Trigger fallback
        const tx = await wallet.sendTransaction({
            to: emitterInstance.address,
            value: ethers.utils.parseEther("0.0")
        });
        console.log("[Deploy Asset] Tx hash: ", tx.hash);
    } catch (e) {
        console.log({e})
    }
}

async function erc20Transfer() {
    try {
        const minterWallet = new ethers.Wallet(validatorPrivKeys[0], provider);

        // Create token
        let tokenFactory = new ethers.ContractFactory(ERC20Contract.abi, ERC20Contract.bytecode, minterWallet);
        const tokenContract = await tokenFactory.deploy();
        await tokenContract.deployed();
        console.log("[ERC20 Transfer] Deployed token!")
        
        // Mint & Approve tokens
        let erc20Instance = new ethers.Contract(tokenContract.address, ERC20Contract.abi, minterWallet);
        await erc20Instance.mint(minterWallet.address, 100);
        console.log("[ERC20 Transfer] Minted tokens!");
        
        await erc20Instance.approve(EMITTER_ADDRESS, 1);
        console.log("[ERC20 Transfer] Approved tokens!");

        // Perform deposit
        const emitterInstance = new ethers.Contract(EMITTER_ADDRESS, EmitterContract.abi, minterWallet);

        // Check the balance before the transfer
        const prebal = await emitterInstance.balances(erc20Instance.address);
        console.log("[ERC20 Transfer] Pre token balaance: ", prebal.toNumber());

        // Make the deposit
        await emitterInstance.depositGenericErc(0, 1, validatorPubkeys[1], erc20Instance.address);
        console.log("[ERC20 Transfer] Created deposit!");

        // Check the balance after the deposit
        const postbal = await emitterInstance.balances(erc20Instance.address);
        console.log("[ERC20 Transfer] Post token balaance: ", postbal.toNumber());

        console.log("[ERC20 Transfer] has the balance increased?", postbal.toNumber() > prebal.toNumber());
    } catch (e) {
        console.log({ e });
    }
}

async function erc721Transfer() {
    try {
        console.log("[ERC721 Transfer] EMITTER_ADDRESS:", EMITTER_ADDRESS);
        const minterWallet = new ethers.Wallet(validatorPrivKeys[0], provider);
       
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
        await erc721Instance.approve(EMITTER_ADDRESS, 1);
        console.log("[ERC721 Transfer] Approved tokens!");

        // Create emitter instance
        const emitterInstance = new ethers.Contract(EMITTER_ADDRESS, EmitterContract.abi, minterWallet);

        // Check pre balance
        const prebal = await erc721Instance.balanceOf(EMITTER_ADDRESS);
        console.log("[ERC721 Transfer] Pre balance:", prebal.toNumber());
        
        // Check the owner
        let owner = await erc721Instance.ownerOf(1);
        console.log("[ERC721 Transfer] Owner of token 1:", owner);
        
        // // Perform deposit
        await emitterInstance.depositNFT(0, validatorPubkeys[1], erc721Instance.address, 1, "0x");
        console.log("[ERC721 Transfer] Created deposit!");

        // Check post balance
        const postbal = await erc721Instance.balanceOf(EMITTER_ADDRESS);
        console.log("[ERC721 Transfer] Pre balance:", postbal.toNumber());
        
        console.log("[ERC20 Transfer] has the balance increased?", postbal.toNumber() > prebal.toNumber());

        // Check the owner
        owner = await erc721Instance.ownerOf(1);
        console.log("[ERC721 Transfer] Owner of token 1:", owner);
    } catch (e) {
        console.log({ e });
    }
}
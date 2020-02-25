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
    .option('--dest <value>', "destination chain", 1)
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

// These are deterministic
const RECEIVER_ADDRESS = "0x705D4Fa884AF2Ae59C7780A0f201109947E2Bf6D";
const CENTRIFUGE_ADDRESS = "0x290f41e61374c715C1127974bf08a3993afd0145";
const EMITTER_ADDRESS = "0x3c747684333605408F9A4907DA043ee4c1A72D9c";
const TEST_EMITTER_ADDRESS = "0x8090062239c909eB9b0433F1184c7DEf6124cc78";

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
        console.log(cli.dest)
        await erc721Transfer(Number(cli.dest));
    } else if (cli.depositAsset) {
        await deployAssetTest();
    }
})();

async function deployReceiver() {
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
        await emitterInstance.depositGenericErc(0, 1, validatorAddress[1], erc20Instance.address);
        console.log("[ERC20 Transfer] Created deposit!");

        // Check the balance after the deposit
        const postbal = await emitterInstance.balances(erc20Instance.address);
        console.log("[ERC20 Transfer] Post token balaance: ", postbal.toNumber());

        console.log("[ERC20 Transfer] has the balance increased?", postbal.toNumber() > prebal.toNumber());
    } catch (e) {
        console.log({ e });
    }
}

async function erc721Transfer(chain) {
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
        const d = await emitterInstance.depositNFT(chain, validatorAddress[1], erc721Instance.address, 1, "0x");
        console.log("[ERC721 Transfer] Created deposit!")
        console.log("[ERC721 Transfer] Deposit Hash", d.hash);

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
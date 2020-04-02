package e2e

import (
	"context"
	"math/big"
	"testing"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func deployTestContracts(t *testing.T, id msg.ChainId) *ethereum.DeployedContracts {
	//deployPK string, chainID *big.Int, url string, relayers int, initialRelayerThreshold *big.Int, minCount uint8
	contracts, err := ethereum.DeployContracts(
		hexutil.Encode(AliceEthKp.Encode())[2:],
		big.NewInt(int64(id)),
		TestEthEndpoint,
		2, // Num relayers
		big.NewInt(2), // Relayer threshold
		0, // unused?
	)
	if err != nil {
		t.Fatal(err)
	}

	return contracts
}

func createEthClient(t *testing.T) (*ethclient.Client, *bind.TransactOpts) {
	ctx := context.Background()
	rpcClient, err := rpc.DialWebsocket(ctx, TestEthEndpoint, "/ws")
	if err != nil {
		t.Fatal(err)
	}
	client := ethclient.NewClient(rpcClient)

	nonce, err := client.PendingNonceAt(ctx, AliceAddr)
	opts := bind.NewKeyedTransactor(AliceEthKp.PrivateKey())
	opts.Nonce = big.NewInt(int64(nonce - 1)) // -1 since we always increment before calling
	opts.Value = big.NewInt(0)               // in wei
	opts.GasLimit = uint64(ethereum.DefaultGasLimit) // in units
	opts.GasPrice = big.NewInt(ethereum.DefaultGasPrice)
	opts.Context = ctx

	return client, opts
}

func deployMintApproveErc20(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address) common.Address {
	// Deploy
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Addr, _, erc20Instance, err := erc20Mintable.DeployERC20Mintable(opts, client)
	if err != nil {
		t.Fatal(err)
	}

	// Mint
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Mint(opts, opts.From, big.NewInt(99))
	if err != nil {
		t.Fatal(err)
	}

	// Approve
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Approve(opts, erc20Handler, big.NewInt(99))
	if err != nil {
		t.Fatal(err)
	}

	return erc20Addr
}

func constructErc20Data(erc20Address, destRecipient common.Address, amount *big.Int) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(erc20Address.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(destRecipient.Bytes(), 32)...)
	data = append(data, math.PaddedBigBytes(amount, 32)...)

	return data
}

func createErc20Deposit(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, contracts *ethereum.DeployedContracts, erc20Contract common.Address) {
	data := constructErc20Data(erc20Contract, BobSubAddr, big.NewInt(10))

	bridgeInstance, err := bridge.NewBridge(contracts.BridgeAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	// Incrememnt Nonce by one
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	if _, err := bridgeInstance.Deposit(
		opts,
		big.NewInt(int64(SubChainId)),
		contracts.ERC20HandlerAddress,
		data,
	); err != nil {
		t.Fatal(err)
	}
}

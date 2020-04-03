// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	centrifugeHandler "github.com/ChainSafe/ChainBridge/bindings/CentrifugeAssetHandler"
	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	msg "github.com/ChainSafe/ChainBridge/message"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func deployTestContracts(t *testing.T, id msg.ChainId) *ethereum.DeployedContracts {
	//deployPK string, chainID *big.Int, url string, relayers int, initialRelayerThreshold *big.Int, minCount uint8
	contracts, err := ethereum.DeployContracts(
		hexutil.Encode(AliceEthKp.Encode())[2:],
		big.NewInt(int64(id)),
		TestEthEndpoint,
		3,             // Num relayers
		big.NewInt(2), // Relayer threshold
		0,             // unused?
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("====== Contracts =======")
	fmt.Printf("Bridge: %s\n", contracts.BridgeAddress.Hex())
	fmt.Printf("Erc20Handler: %s\n", contracts.ERC20HandlerAddress.Hex())
	fmt.Printf("(Generic) Centrifuge Handler: %s\n", contracts.CentrifugeHandlerAddress.Hex())
	fmt.Println("========================")
	return contracts
}

func createEthClient(t *testing.T) (*ethclient.Client, *bind.TransactOpts) {
	ctx := context.Background()
	rpcClient, err := rpc.DialWebsocket(ctx, TestEthEndpoint, "/ws")
	if err != nil {
		t.Fatal(err)
	}
	client := ethclient.NewClient(rpcClient)

	nonce, err := client.PendingNonceAt(ctx, CharlieEthAddr)
	if err != nil {
		t.Fatal(err)
	}
	opts := bind.NewKeyedTransactor(CharlieEthKp.PrivateKey())
	opts.Nonce = big.NewInt(int64(nonce - 1))        // -1 since we always increment before calling
	opts.Value = big.NewInt(0)                       // in wei
	opts.GasLimit = uint64(ethereum.DefaultGasLimit) // in units
	opts.GasPrice = big.NewInt(ethereum.DefaultGasPrice)
	opts.Context = ctx

	return client, opts
}

// deployMintApproveErc20 funds the account with a newly created erc20, to prepare for initiating a transfer
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
	fmt.Printf("Approving: %s\n", erc20Handler.Hex())
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Approve(opts, erc20Handler, big.NewInt(99))
	if err != nil {
		t.Fatal(err)
	}

	return erc20Addr
}

// deployAndFundEr20 sets up a new erc20 contract and funds the bridge/handler
func deployAndFundErc20(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address) common.Address { //nolint:unused,deadcode
	// Deploy
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Addr, _, erc20Instance, err := erc20Mintable.DeployERC20Mintable(opts, client)
	if err != nil {
		t.Fatal(err)
	}

	// Mint to handler
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Mint(opts, erc20Addr, big.NewInt(99))
	if err != nil {
		t.Fatal(err)
	}

	return erc20Addr
}

// constructErc20Data constructs the data field to be passed into a deposit call
func constructErc20DepositData(erc20Address common.Address, destRecipient []byte, amount *big.Int) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(erc20Address.Bytes(), 32)...)
	data = append(data, math.PaddedBigBytes(amount, 32)...)
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(destRecipient))), 32)...)
	data = append(data, destRecipient...)
	return data
}

func createErc20Deposit(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, contracts *ethereum.DeployedContracts, erc20Contract common.Address) {
	data := constructErc20DepositData(erc20Contract, BobSubKp.AsKeyringPair().PublicKey, big.NewInt(10))

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

func waitForEvent(t *testing.T, client *ethclient.Client, contract common.Address, subStr ethereum.EventSig) {
	query := eth.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{subStr.GetTopic()},
		},
	}

	ch := make(chan ethtypes.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case evt := <-ch:
			fmt.Printf("%s: %#v\n", subStr, evt.Topics)
			sub.Unsubscribe()
			close(ch)
			return
		case err := <-sub.Err():
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func getHash(t *testing.T, client *ethclient.Client, hash [32]byte, contract common.Address) {
	instance, err := centrifugeHandler.NewCentrifugeAssetHandler(contract, client)
	if err != nil {
		t.Fatal(err)
	}

	exists, err := instance.GetHash(&bind.CallOpts{}, hash)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("Hash does not exist")
	}
}

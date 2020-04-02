package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	"github.com/ChainSafe/ChainBridge/chains/substrate"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
)

const TestTimeout = time.Second * 30
const TestEthEndpoint = "ws://localhost:8545"
const TestSubEndpoint = "ws://localhost:9944"
const EthChainId = msg.ChainId(0)
const SubChainId = msg.ChainId(1)

var AliceEthKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var BobEthKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]
var AliceSubKp = keystore.TestKeyRing.SubstrateKeys[keystore.AliceKey]
var BobSubKp = keystore.TestKeyRing.SubstrateKeys[keystore.BobKey]
var AliceAddr = common.HexToAddress(AliceEthKp.Address())
var BobAddr = common.HexToAddress(BobEthKp.Address())
var BobSubAddr = BobSubKp.AsKeyringPair().Address

func setLogger(lvl log.Lvl) {
	logger := log.Root()
	handler := logger.GetHandler()
	log.Root().SetHandler(log.LvlFilterHandler(lvl, handler))
}

func createEthConfig(key, bridgeAddress, erc20HandlerAddress string) *core.ChainConfig {
	return &core.ChainConfig{
		Name:         fmt.Sprintf("ethereum(%s)", key),
		Id:           EthChainId,
		Endpoint:     TestEthEndpoint,
		From:         "",
		KeystorePath: key,
		Insecure:     true,
		Opts:         map[string]string{
			"contract": bridgeAddress,
			"erc20Handler": erc20HandlerAddress,
		},
	}
}

func createSubConfig(key, bridgeAddress, erc20HandlerAddress string) *core.ChainConfig {
	return &core.ChainConfig{
		Name:         fmt.Sprintf("substrate(%s)", key),
		Id:           SubChainId,
		Endpoint:     TestSubEndpoint,
		From:         "",
		KeystorePath: key,
		Insecure:     true,
		Opts:         map[string]string{
			"contract": bridgeAddress,
			"erc20Handler": erc20HandlerAddress,
		},
	}
}

func createBridge(t *testing.T, name, bridgeAddress, erc20HandlerAddres string) *core.Core {
	eth, err := ethereum.InitializeChain(createEthConfig(name, bridgeAddress, erc20HandlerAddres))
	if err != nil {
		t.Fatal(err)
	}
	sub, err := substrate.InitializeChain(createSubConfig(name, bridgeAddress, erc20HandlerAddres))
	if err != nil {
		t.Fatal(err)
	}

	bridge := core.NewCore()
	bridge.AddChain(eth)
	bridge.AddChain(sub)

	return bridge
}

func TestErc20ToSubstrate(t *testing.T) {
	setLogger(log.LvlInfo)
	// Deploy contracts, mint, approve

	// Creates an alice client we can use to initiate everything
	contracts := deployTestContracts(t, EthChainId)
	ethClient, opts := createEthClient(t)
	erc20Contract := deployMintApproveErc20(t, ethClient, opts, contracts.ERC20HandlerAddress)

	// Create and start two bridges with both chains
	bridgeA := createBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex())
	bridgeB := createBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex())

	go bridgeA.Start()
	go bridgeB.Start()
	// Initiate transfer
	createErc20Deposit(t, ethClient, opts, contracts, erc20Contract)

	// Check for success event
	subClient := createSubClient(t)
	success := make(chan bool)
	fail := make(chan bool)
	go watchForProposalSucceeded(t, subClient, success)
	go watchForProposalFailed(t, subClient, fail)


	select {
	case <-success:
		return
	case <- fail:
		t.Fatal("Proposal failed")
	case <-time.After(TestTimeout):
		t.Fatalf("test timed out")
	}
}
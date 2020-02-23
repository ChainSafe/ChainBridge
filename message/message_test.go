package msg

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	ethcmn "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

var TestAddress = ethcmn.HexToAddress("34c59fBf82C9e31BA9CBB5faF4fe6df05de18Ad4")

func TestCreateDepositProposalData(t *testing.T) {
	hash, err := hex.DecodeString("b6e25575ab25a1938070eeb64ac4cd6df49af423327877522bec719815dc5e27")
	if err != nil {
		t.Fatal(err)
	}

	hashBytes := [32]byte{}
	copy(hashBytes[:], hash)
	keccakHash := ethcrypto.Keccak256Hash(hashBytes[:]).Bytes()

	depositId := big.NewInt(0)
	originChain := ChainId(1)

	m := &Message{
		Type: CreateDepositProposalType,
		Data: hashBytes[:],
	}

	m.EncodeCreateDepositProposalData(depositId, originChain)

	hashRes, depositIdRes, originChainRes, err := m.DecodeCreateDepositProposalData()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(keccakHash, hashRes[:]) {
		t.Fatalf("Fail: got %x x %x", hashRes, hashBytes)
	}

	if depositId.Cmp(depositIdRes) != 0 {
		t.Fatalf("Fail: got %d expected %d", depositIdRes, depositId)
	}

	if originChain == ChainId(originChainRes.Int64()) {
		t.Fatalf("Fail: got %d expected %d", originChainRes, originChain)
	}
}

func TestVoteDepositProposalData(t *testing.T) {
	depositId := big.NewInt(0)
	originChain := big.NewInt(1)
	vote := byte(1)

	m := &Message{
		Type: VoteDepositProposalType,
	}

	m.EncodeVoteDepositProposalData(depositId, originChain, vote)

	depositIdRes, originChainRes, voteRes, err := m.DecodeVoteDepositProposalData()
	if err != nil {
		t.Fatal(err)
	}

	if depositId.Cmp(depositIdRes) != 0 {
		t.Fatalf("Fail: got %d expected %d", depositIdRes, depositId)
	}

	if originChain.Cmp(originChainRes) != 0 {
		t.Fatalf("Fail: got %d expected %d", originChainRes, originChain)
	}

	if vote != voteRes {
		t.Fatalf("Fail: got %d expected %d", voteRes, vote)
	}
}

func TestExecuteDepositData(t *testing.T) {
	depositId := big.NewInt(0)
	originChain := big.NewInt(1)
	address := TestAddress
	data := []byte("noot")

	m := &Message{
		Type: VoteDepositProposalType,
	}

	m.EncodeExecuteDepositData(depositId, originChain, address, data)

	depositIdRes, originChainRes, addressRes, dataRes, err := m.DecodeExecuteDepositData()
	if err != nil {
		t.Fatal(err)
	}

	if depositId.Cmp(depositIdRes) != 0 {
		t.Fatalf("Fail: got %d expected %d", depositIdRes, depositId)
	}

	if originChain.Cmp(originChainRes) != 0 {
		t.Fatalf("Fail: got %d expected %d", originChainRes, originChain)
	}

	if !bytes.Equal(data, dataRes) {
		t.Fatalf("Fail: got %x expected %x", dataRes, data)
	}

	if !bytes.Equal(address.Bytes(), addressRes.Bytes()) {
		t.Fatalf("Fail: got %x expected %x", addressRes, address)
	}
}

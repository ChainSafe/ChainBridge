package msg

import (
	"errors"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

type MessageType string

var DepositAssetType MessageType = "centrifuge_deposit_asset"
var CreateDepositProposalType MessageType = "create_deposit_proposal"
var VoteDepositProposalType MessageType = "vote_deposit_proposal"
var ExecuteDepositType MessageType = "execute_deposit"

// Message is used as a generic format to communicate between chains
type Message struct {
	Source      ChainId // Source where message was initiated
	Destination ChainId // Destination chain of message
	// TODO: Add data fields
	Type MessageType // type of bridge transfer
	Data []byte      // data associated with event sequence
}

func (m *Message) EncodeCreateDepositProposalData(depositId *big.Int, originChain ChainId) {
	hash := ethcrypto.Keccak256Hash(m.Data).Bytes()
	depositIdBytes := byteSliceTo32Bytes(depositId.Bytes())
	originChainBytes := byteSliceTo32Bytes(make([]byte, originChain))
	m.Data = append(append(hash[:], depositIdBytes...), originChainBytes[:]...)
}

func (m *Message) DecodeCreateDepositProposalData() ([32]byte, *big.Int, *big.Int, error) {
	if len(m.Data) < 96 {
		return [32]byte{}, nil, nil, errors.New("cannot decode: data is <96 bytes")
	}

	hashSlice := m.Data[:32]
	hash := [32]byte{}
	copy(hash[:], hashSlice)

	depositIdBytes := m.Data[32:64]
	originChainBytes := m.Data[64:96]

	depositId := big.NewInt(0).SetBytes(depositIdBytes)
	originChain := big.NewInt(0).SetBytes(originChainBytes)
	return hash, depositId, originChain, nil
}

func (m *Message) EncodeVoteDepositProposalData(depositId, originChain *big.Int, vote byte) {
	depositIdBytes := byteSliceTo32Bytes(depositId.Bytes())
	originChainBytes := byteSliceTo32Bytes(originChain.Bytes())
	m.Data = append(append(depositIdBytes[:], originChainBytes[:]...), vote)
}

func (m *Message) DecodeVoteDepositProposalData() (*big.Int, *big.Int, byte, error) {
	if len(m.Data) < 65 {
		return nil, nil, 0, errors.New("cannot decode: data is <65 bytes")
	}

	depositIdBytes := m.Data[:32]
	originChainBytes := m.Data[32:64]

	depositId := big.NewInt(0).SetBytes(depositIdBytes)
	originChain := big.NewInt(0).SetBytes(originChainBytes)
	return depositId, originChain, m.Data[64], nil
}

func (m *Message) EncodeExecuteDepositData(depositId, originChain *big.Int, address ethcommon.Address, data []byte) {
	depositIdBytes := byteSliceTo32Bytes(depositId.Bytes())
	originChainBytes := byteSliceTo32Bytes(originChain.Bytes())
	m.Data = append(append(append(depositIdBytes[:], originChainBytes[:]...), address[:]...), data...)
}

func (m *Message) DecodeExecuteDepositData() (*big.Int, *big.Int, ethcommon.Address, []byte, error) {
	if len(m.Data) < 84 {
		return nil, nil, ethcommon.Address{}, nil, errors.New("cannot decode: data is <84 bytes")
	}

	depositIdBytes := m.Data[:32]
	originChainBytes := m.Data[32:64]

	depositId := big.NewInt(0).SetBytes(depositIdBytes)
	originChain := big.NewInt(0).SetBytes(originChainBytes)

	address := m.Data[64:84]

	return depositId, originChain, ethcommon.BytesToAddress(address), m.Data[84:], nil
}

func byteSliceTo32Bytes(in []byte) []byte {
	out := [32]byte{}
	copy(out[32-len(in):], in)
	return out[:]
}

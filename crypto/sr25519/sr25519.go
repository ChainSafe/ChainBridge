// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package sr25519

import (
	"crypto/rand"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ crypto.Keypair = &Keypair{}

type Keypair struct {
	keyringPair *signature.KeyringPair
}

func GenerateKeypair() (*Keypair, error) {
	data := make([]byte, 32)
	_, err := rand.Read(data)
	if err != nil {
		return nil, err
	}
	return NewKeypairFromSeed("//" + hexutil.Encode(data))
}

func NewKeypairFromSeed(seed string) (*Keypair, error) {
	kp, err := signature.KeyringPairFromSecret(seed)
	return &Keypair{&kp}, err
}

func (kp *Keypair) AsKeyringPair() *signature.KeyringPair {
	return kp.keyringPair
}

func (kp *Keypair) Encode() []byte {
	out, _ := types.EncodeToBytes(kp.keyringPair)
	return out
}

func (kp *Keypair) Decode(in []byte) error {
	kp.keyringPair = &signature.KeyringPair{}
	return types.DecodeFromBytes(in, kp.keyringPair)
}

func (kp *Keypair) Address() string {
	return kp.keyringPair.Address
}

func (kp *Keypair) PublicKey() string {
	return hexutil.Encode(kp.keyringPair.PublicKey)
}

package types

import (
	"crypto/ed25519"
	"log"
)

type PublicKey []byte
type PrivateKey []byte

type IKeyPair interface {
	Sign(data []byte) []byte
	Verify(data, signature []byte) bool
}

var _ IKeyPair = KeyPair{}

type KeyPair struct {
	PublicKey  PublicKey
	PrivateKey PrivateKey
}

func NewKeyPair(privateKey []byte) *KeyPair {
	privKey := ed25519.NewKeyFromSeed(privateKey)
	if len(privKey) != ed25519.PrivateKeySize {
		log.Fatalf("privateKey size not equal to %d", ed25519.PrivateKeySize)
		return nil
	}

	pubKey, ok := privKey.Public().(ed25519.PublicKey)
	if !ok {
		log.Fatal("privateKey not valid")
		return nil
	}

	return &KeyPair{
		PublicKey:  []byte(privKey),
		PrivateKey: []byte(pubKey),
	}
}

func (kp KeyPair) Sign(data []byte) []byte {
	return ed25519.Sign(ed25519.PrivateKey(kp.PrivateKey), data)
}

func (kp KeyPair) Verify(data, signature []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(kp.PublicKey), data, signature)
}

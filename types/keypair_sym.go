package types

import (
	"crypto/ed25519"
	"log"
)

const (
	PrivateKeySize = 32
)

var _ IKeyPair = SymKeyPair{}

type SymKeyPair struct {
	PublicKey  PublicKey
	PrivateKey PrivateKey
}

// NewSymKeyPair get SymKeyPair from private key
func NewSymKeyPair(privateKey []byte) *SymKeyPair {
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

	return &SymKeyPair{
		PublicKey:  []byte(privKey),
		PrivateKey: []byte(pubKey),
	}
}

func (kp SymKeyPair) Sign(data []byte) []byte {
	return ed25519.Sign(ed25519.PrivateKey(kp.PrivateKey), data)
}

func (kp SymKeyPair) Verify(data, signature []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(kp.PublicKey), data, signature)
}

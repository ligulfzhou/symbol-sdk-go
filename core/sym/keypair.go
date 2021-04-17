package sym

import (
	"crypto/ed25519"
	"log"
	"strconv"
	"symbol-sdk-go/core/utils"
)

type VerifierBase interface {
	Verify() bool
}

type KeyPair struct {
	PrivKey []byte
	PubKey  []byte
}

func NewKeyPair(privateKey string) *KeyPair {
	privKey := ed25519.NewKeyFromSeed(utils.Unhexify(privateKey))
	if len(privKey) != ed25519.SeedSize {
		log.Fatal("privateKey size not equal to " + strconv.Itoa(ed25519.SeedSize))
		return nil
	}
	PubKey, ok := privKey.Public().(ed25519.PublicKey)
	if !ok {
		log.Fatal("privateKey not valid")
		return nil
	}

	return &KeyPair{
		PrivKey: privKey,
		PubKey:  PubKey,
	}
}

func (kp KeyPair) PrivateKey() string {
	return utils.Hexify(kp.PrivKey)
}

func (kp KeyPair) PublicKey() string {
	return utils.Hexify(kp.PubKey)
}

type Verifier struct {
	PubKey []byte
}

func NewVerifier(publicKey string) *Verifier {

}

/// todo,
/*
class Verifier:
    """Verifies signatures signed by a single key pair."""

    def __init__(self, public_key):
        """Creates a verifier from a public key."""
        self._pk = ed25519.Ed25519PublicKey.from_public_bytes(public_key.bytes)

    def verify(self, message, signature):
        """Verifies a message signature."""
        try:
            self._pk.verify(signature.bytes, message)
            return True
        except InvalidSignature:
            return False
*/

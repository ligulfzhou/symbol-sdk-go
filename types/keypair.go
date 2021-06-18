package types

type PublicKey []byte
type PrivateKey []byte

type IKeyPair interface {
	Sign(data []byte) []byte
	Verify(data, signature []byte) bool
}

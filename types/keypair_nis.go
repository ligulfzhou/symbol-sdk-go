package types

var _ IKeyPair = NisKeyPair{}

type NisKeyPair struct {
	PublicKey  PublicKey
	PrivateKey PrivateKey
}

// NewNisKeyPair todo
func NewNisKeyPair(privateKey []byte) *NisKeyPair {
	return &NisKeyPair{}
}

// Sign todo
func (kp NisKeyPair) Sign(data []byte) []byte {
	return []byte{}
}

// Verify todo
func (kp NisKeyPair) Verify(data, signature []byte) bool {
	return true
}

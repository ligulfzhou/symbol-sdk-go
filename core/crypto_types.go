package core

var (
	Hash256Zero   *Hash256
	SignatureZero *Signature
)

const (
	Hash256Size   = 32
	SignatureSize = 64
)

type Hash256 struct {
	ByteArray
}

func NewHash256(hash256 interface{}) *Hash256 {
	return &Hash256{
		*NewByteArray(Hash256Size, hash256, ""),
	}
}

type Signature struct {
	*ByteArray
}

func NewSignature(signature interface{}) *Signature {
	return &Signature{
		NewByteArray(SignatureSize, signature, "signature"),
	}
}

func init() {
	Hash256Zero = NewHash256(GenByteArray(0, Hash256Size))
	SignatureZero = NewSignature(GenByteArray(0, SignatureSize))
}

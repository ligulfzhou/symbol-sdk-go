package core

import "symbol-sdk-go/core/utils"

const (
	Hash256Size = 32
)

type Hash256 ByteArray

func NewHash256(hash256 []byte) *Hash256 {
	return (*Hash256)(NewByteArray(
		Hash256Size, hash256, "",
	))
}

func Hash256Zero() *Hash256 {
	hash256 := utils.GenByteArray(0, Hash256Size)
	return NewHash256(hash256)
}

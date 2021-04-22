package core

import (
	"bytes"
	"strings"
)

type ByteArray struct {
	Tag   string
	Size  int
	Bytes []byte
}

func NewByteArray(size int, array interface{}, tag string) *ByteArray {
	ba := &ByteArray{
		Size: size,
		Tag:  tag,
	}

	arrayString, ok := array.(string)
	if ok {
		ba.Bytes = Unhexify(arrayString)
	}

	arrayBytes, ok := array.([]byte)
	if ok {
		ba.Bytes = arrayBytes
	}

	return ba
}

func (ba ByteArray) Equal(other ByteArray) bool {
	return bytes.Equal(ba.Bytes, other.Bytes) && ba.Tag == other.Tag
}

func (ba ByteArray) String() string {
	return strings.ToUpper(Hexify(ba.Bytes))
}

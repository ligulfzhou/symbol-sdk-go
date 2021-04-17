package core

import (
	"bytes"
	"strings"
	"symbol-sdk-go/core/utils"
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
		ba.Bytes = utils.Unhexify(arrayString)
	}

	arrayBytes, ok := array.([]byte)
	if ok {
		ba.Bytes = arrayBytes
	}

	return ba
}

func (ba ByteArray) Equal(other ByteArray) bool {
	return bytes.Compare(ba.Bytes, other.Bytes) == 0 && ba.Tag == other.Tag
}

func (ba ByteArray) String() string {
	return strings.ToUpper(utils.Hexify(ba.Bytes))
}

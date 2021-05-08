package utils

import (
	"encoding/binary"
	"encoding/hex"
)

const (
	Hash256Size = 32
)

var (
	Hash256Zero []byte = GenByteArray(0, Hash256Size)
)

func GenByteArray(item, count int) []byte {
	res := make([]byte, count)
	idx := 0
	for idx < count {
		res[idx] = byte(item)
		idx = idx + 1
	}
	return res
}

func Unhexify(s string) []byte {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return nil
	}
	return decoded
}

func Hexify(bs []byte) string {
	return hex.EncodeToString(bs)
}

func Int2BigEndianBytes(len, num int) []byte {
	bs := make([]byte, len)
	binary.BigEndian.PutUint32(bs, uint32(num))
	return bs
}
func Int2LittleEndianBytes(len, num int) []byte {
	bs := make([]byte, len)
	binary.LittleEndian.PutUint32(bs, uint32(num))
	return bs
}
func Uint642LittleEndianBytes(num uint64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, num)
	return bs
}

func BigEndianBytes2Int(bs []byte) int {
	data := binary.BigEndian.Uint32(bs)
	return int(data)
}

func LittleEndianBytes2Int(bs []byte) int {
	data := binary.LittleEndian.Uint32(bs)
	return int(data)
}

func LittleEndianBytes2Uint64(bs []byte) uint64 {
	data := binary.LittleEndian.Uint64(bs)
	return data
}

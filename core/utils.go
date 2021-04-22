package core

import "encoding/hex"

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

func GenByteArray(item, count int) []byte {
	res := make([]byte, count)
	idx := 0
	for idx < count {
		res[idx] = byte(item)
		idx = idx + 1
	}

	return res
}

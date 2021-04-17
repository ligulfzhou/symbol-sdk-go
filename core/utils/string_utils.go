package utils

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

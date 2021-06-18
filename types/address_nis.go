package types

import "encoding/base32"

type NisAddress string

func DecodeNisAddress(address string) ([]byte, error) {
	bytes, err := base32.StdEncoding.DecodeString(address)
	return bytes, err
}

func (n NisAddress) DecodeAddress() ([]byte, error) {
	return base32.StdEncoding.DecodeString(string(n))
}

func EncodeNisAddress(address []byte) NisAddress {
	return NisAddress(base32.StdEncoding.EncodeToString(address))
}

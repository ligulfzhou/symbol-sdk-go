package types

import "encoding/base32"

const (
	AddressSize = 24
)

type SymAddress string

func (s SymAddress) DecodeAddress() ([]byte, error) {
	bytes, err := base32.StdEncoding.DecodeString(string(s) + "A")
	if err != nil {
		return nil, err
	}

	return bytes[0 : len(bytes)-1], nil
}

func EncodeSymAddress(address []byte) SymAddress {
	bytes := append(address, byte(0))
	encodedAddress := base32.StdEncoding.EncodeToString(bytes)
	return SymAddress(encodedAddress[:len(encodedAddress)-1])
}

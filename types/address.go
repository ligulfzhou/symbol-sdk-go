package types

import "encoding/base32"

const (
	AddressSize = 24
)

//type Address struct {
//	Address string
//	Network string
//}

func DecodeAddress(address string) ([]byte, error) {
	bytes, err := base32.StdEncoding.DecodeString(address + "A")
	if err != nil {
		return nil, err
	}

	return bytes[0 : len(bytes)-1], nil
}

func EncodeAddress(address []byte) string {
	bytes := append(address, byte(0))
	encodedAddress := base32.StdEncoding.EncodeToString(bytes)
	return encodedAddress[:len(encodedAddress)-1]
}

package types

import (
	"bytes"
	"strings"
	"symbol-sdk-go/utils"
	"testing"
)

var (
	EncodedAddress = SymAddress("TBLYH55IHPS5QCCMNWR3GZWKV6WMCKPTNI7KSDA")
	DecodedAddress = utils.Unhexify("985783F7A83BE5D8084C6DA3B366CAAFACC129F36A3EA90C")
)

func TestDecodeAddress(t *testing.T) {
	if len(DecodedAddress) != AddressSize {
		t.Fail()
	}

	decodedAddress, err := EncodedAddress.DecodeAddress()
	if err != nil {
		t.Fail()
	}

	if bytes.Compare(decodedAddress, DecodedAddress) != 0 {
		t.Fail()
	}
}

func TestEncodeAddress(t *testing.T) {
	encodedAddress := EncodeSymAddress(DecodedAddress)
	if strings.Compare(string(encodedAddress), string(EncodedAddress)) != 0 {
		t.Fail()
	}
}

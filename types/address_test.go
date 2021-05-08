package types

import (
	"bytes"
	"strings"
	"symbol-sdk-go/utils"
	"testing"
)

var (
	EncodedAddress = "TBLYH55IHPS5QCCMNWR3GZWKV6WMCKPTNI7KSDA"
	DecodedAddress = utils.Unhexify("985783F7A83BE5D8084C6DA3B366CAAFACC129F36A3EA90C")
)

func TestDecodeAddress(t *testing.T) {

	if len(DecodedAddress) != AddressSize {
		t.Fail()
	}

	decodedAddress, err := DecodeAddress(EncodedAddress)
	t.Log(decodedAddress)
	t.Log(DecodedAddress)
	if err != nil {
		t.Fail()
	}

	if bytes.Compare(decodedAddress, DecodedAddress) != 0 {
		t.Fail()
	}
}

func TestEncodeAddress(t *testing.T) {
	encodedAddress := EncodeAddress(DecodedAddress)
	t.Log(encodedAddress)
	t.Log(EncodedAddress)
	if strings.Compare(encodedAddress, EncodedAddress) != 0 {
		t.Fail()
	}
}

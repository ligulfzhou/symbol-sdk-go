package types

import (
	"bytes"
	"symbol-sdk-go/utils"
	"testing"
)

var (
	EncodedNisAddress = NisAddress("TCFGSLITSWMRROU2GO7FPMIUUDELUPSZUNUEZF33")
	DecodedNisAddress = utils.Unhexify("988A692D13959918BA9A33BE57B114A0C8BA3E59A3684C977B")
)

func TestNisAddress_DecodeAddress(t *testing.T) {
	decodedNisAddress, err := EncodedNisAddress.DecodeAddress()
	if err != nil {
		t.Fail()
	}

	if bytes.Compare(decodedNisAddress, DecodedNisAddress) != 0 {
		t.Fail()
	}
}

func TestEncodeNisAddress(t *testing.T) {

	nisAddress := EncodeNisAddress(DecodedNisAddress)
	if string(nisAddress) != string(EncodedNisAddress) {
		t.Fail()
	}
}

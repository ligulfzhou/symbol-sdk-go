package types

import (
	"strings"
	"symbol-sdk-go/utils"
	"testing"
)

func TestNetworks(t *testing.T) {
	networkNames := []string{"public", "private", "public_test", "private_test"}

	var names []string
	for _, network := range SymNetworks {
		names = append(names, network.Name)
	}
	if !utils.StringListEqual(networkNames, names) {
		t.Fail()
	}
}

func TestNetwork_PublicKeyToAdress(t *testing.T) {
	publicKey := utils.Unhexify("C5FB65CB902623D93DF2E682FFB13F99D50FAC24D5FF2A42F68C7CA1772FE8A0")

	if strings.Compare(SymPublicNetwork.PublicKeyToAdress(publicKey), "NBLYH55IHPS5QCCMNWR3GZWKV6WMCKPTNKZIBEY") != 0 {
		t.Fail()
	}

	if strings.Compare(SymPublicTest.PublicKeyToAdress(publicKey), "TBLYH55IHPS5QCCMNWR3GZWKV6WMCKPTNI7KSDA") != 0 {
		t.Fail()
	}
}

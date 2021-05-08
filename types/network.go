package types

import (
	"symbol-sdk-go/utils"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var (
	PublicNetwork  = NewNetwork("public", 0x68, utils.Unhexify("57F7DA205008026C776CB6AED843393F04CD458E0AA2D9F1D5F31A402072B2D6"))
	PrivateNetwork = NewNetwork("private", 0x78, nil)
	PublicTest     = NewNetwork("public_test", 0x98, utils.Unhexify("45FBCF2F0EA36EFA7923C9BC923D6503169651F7FA4EFC46A8EAF5AE09057EBD"))
	PrivateTest    = NewNetwork("private_test", 0xab, nil)
	Networks       = []*Network{PublicNetwork, PrivateNetwork, PublicTest, PrivateTest}
)

type Network struct {
	Name               string
	Identifier         int
	GenerationHashSeed []byte
}

func NewNetwork(name string, identifier int, seed []byte) *Network {
	return &Network{
		Name:               name,
		Identifier:         identifier,
		GenerationHashSeed: seed,
	}
}

func (n Network) PublicKeyToAdress(publicKey PublicKey) string {
	partOneBuilder := sha3.New256()
	partOneBuilder.Write(publicKey)
	partOneHash := partOneBuilder.Sum(nil)

	partTwoBuilder := ripemd160.New()
	partTwoBuilder.Write(partOneHash)
	partTwoHash := partTwoBuilder.Sum(nil)

	var addressWithoutCheckSum []byte
	addressWithoutCheckSum = append(addressWithoutCheckSum, byte(n.Identifier))
	addressWithoutCheckSum = append(addressWithoutCheckSum, partTwoHash...)

	partThreeBuilder := sha3.New256()
	partThreeBuilder.Write(addressWithoutCheckSum)
	checkSum := partThreeBuilder.Sum(nil)[0:3]

	address := append(addressWithoutCheckSum, checkSum...)
	return EncodeAddress(address)
}

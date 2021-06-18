package types

import (
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var (
	NisMainNet  = NewNisNetwork("mainnet", 0x67, nil)
	NisTestNet  = NewNisNetwork("testnet", 0x98, nil)
	NisNetworks = []*NisNetwork{NisMainNet, NisTestNet}
)

type NisNetwork Network

func NewNisNetwork(name string, identifier int, seed []byte) *NisNetwork {
	return &NisNetwork{
		Name:               name,
		Identifier:         identifier,
		GenerationHashSeed: seed,
	}
}

func (n Network) PublicKeyToAdress(publicKey PublicKey) string {
	partOneBuilder := sha3.NewLegacyKeccak256()
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
	checkSum := partThreeBuilder.Sum(nil)

	address := append(addressWithoutCheckSum, checkSum...)
	return string(EncodeNisAddress(address))
}

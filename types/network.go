package types

type INetwork interface {
	PublicKeyToAdress(publicKey PublicKey) string
}

type Network struct {
	Name               string
	Identifier         int
	GenerationHashSeed []byte
}

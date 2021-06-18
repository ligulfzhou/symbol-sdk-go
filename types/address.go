package types

type IAddress interface {
	DecodeAddress() ([]byte, error)
}

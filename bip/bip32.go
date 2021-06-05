package bip

import (
	"crypto/hmac"
	"crypto/sha512"
	"strings"
	"symbol-sdk-go/types"
	"symbol-sdk-go/utils"

	"github.com/brianium/mnemonic"
)

const (
	DEFAULT_LANGUAGE   = string(mnemonic.English)
	DEFAULT_CURVE_NAME = "ed25519"
)

type Bip32Node struct {
	Key, Data  []byte
	PrivateKey []byte
	ChainCode  []byte
}

func NewBip32Node(key, data []byte) *Bip32Node {
	h := hmac.New(sha512.New, key)
	h.Write(data)
	hash := h.Sum(nil)

	return &Bip32Node{
		Key:        key,
		Data:       data,
		PrivateKey: hash[0:types.PrivateKeySize],
		ChainCode:  hash[types.PrivateKeySize:],
	}
}

func (node *Bip32Node) DeriveOne(identifier int) *Bip32Node {
	data := make([]byte, 0)
	data = append(data, []byte{0}...)
	data = append(data, node.PrivateKey...)
	data = append(data, utils.Int2BigEndianBytes(4, 0x80000000|identifier)...)
	return NewBip32Node(node.ChainCode, data)
}

func (node *Bip32Node) DerivePath(path []int) *Bip32Node {
	nextNode := node
	for i := 0; i < len(path); i++ {
		nextNode = nextNode.DeriveOne(path[i])
	}
	return nextNode
}

type Bip32 struct {
	Key      []byte
	Language string
}

func NewBip32(curveName, language string) *Bip32 {
	key := []byte(curveName + " seed")
	return &Bip32{
		Key:      key,
		Language: language,
	}
}

func NewDefaultBip32() *Bip32 {
	return NewBip32(DEFAULT_CURVE_NAME, DEFAULT_LANGUAGE)
}

// NodeFromSeed create bip32 root node from seed
func (bip32 *Bip32) NodeFromSeed(seed []byte) *Bip32Node {
	return NewBip32Node(bip32.Key, seed)
}

// NodeFromMnemonic create bip32 root node from BIP39 mnemonic and password
func (bip32 *Bip32) NodeFromMnemonic(mnemonicList []string, password string) *Bip32Node {
	seed := mnemonic.NewSeed(strings.Join(mnemonicList, " "), password)
	return bip32.NodeFromSeed(seed.Bytes)
}

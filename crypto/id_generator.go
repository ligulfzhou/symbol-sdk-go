package crypto

import (
	"errors"
	"strings"
	"symbol-sdk-go/types"
	"symbol-sdk-go/utils"

	"golang.org/x/crypto/sha3"
)

var NamespaceFlag uint64 = 1 << 63

// GenerateMosaicId Generates a mosaic id from an owner address and a nonce
func GenerateMosaicId(address string, nonce int) (uint64, error) {
	hasher := sha3.New256()
	hasher.Write(utils.Int2LittleEndianBytes(4, nonce))
	decodedAddr, err := types.DecodeAddress(address)
	if err != nil {
		return 0, err
	}
	hasher.Write(decodedAddr)
	digest := hasher.Sum(nil)

	result := utils.LittleEndianBytes2Uint64(digest[:8])
	if result&NamespaceFlag > 0 {
		result -= NamespaceFlag
	}
	return result, nil
}

// GenerateNamespaceId Generates a namespace id from a name and an optional parent namespace id
func GenerateNamespaceId(name string, parentNamespaceId uint64) uint64 {
	hasher := sha3.New256()
	hasher.Write(utils.Uint642LittleEndianBytes(parentNamespaceId))
	hasher.Write([]byte(name))
	digest := hasher.Sum(nil)

	result := utils.LittleEndianBytes2Uint64(digest[:8])
	return result | NamespaceFlag
}

// GenerateMosaicAliasId Generates a mosaic id from a fully qualified mosaic alias name.
func GenerateMosaicAliasId(fullyQualifiedName string) (uint64, error) {
	path, err := GenerateNamespacePath(fullyQualifiedName)
	if err != nil {
		return 0, err
	}
	return path[len(path)-1], nil
}

// IsValidNamespaceName Returns true if a name is a valid namespace name.
func IsValidNamespaceName(name string) bool {
	for _, char := range name {
		if !isAlphanum(char) && char != '-' && char != '_' {
			return false
		}
	}
	return true
}

// GenerateNamespacePath Parses a fully qualified namespace name into a path.
func GenerateNamespacePath(fullyQualifiedName string) ([]uint64, error) {
	var path []uint64
	var parentNamespaceId uint64 = 0
	for _, name := range strings.Split(fullyQualifiedName, ".") {
		if !IsValidNamespaceName(name) {
			return nil, errors.New("fully qualified name is invalid")
		}
		path = append(path, GenerateNamespaceId(name, parentNamespaceId))
		parentNamespaceId = path[len(path)-1]
	}
	return path, nil
}

func isAlphanum(char int32) bool {
	if 'a' <= char && char <= 'z' || '0' <= char && char <= '9' {
		return true
	}
	return false
}

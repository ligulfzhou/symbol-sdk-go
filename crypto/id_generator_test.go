package crypto

import (
	"symbol-sdk-go/utils"
	"testing"
)

func TestGenerateMosaicId(t *testing.T) {
	mosaicId, err := GenerateMosaicId("TATNE7Q5BITMUTRRN6IB4I7FLSDRDWZA37JGO5Q", 812613930)
	if err != nil {
		t.Fail()
	}
	if 0x570FB3ED9379624C != mosaicId {
		t.Fail()
	}
}

func TestGenerateNamespaceId(t *testing.T) {
	// correct case
	namespaceId := GenerateNamespaceId("symbol", 0)
	if namespaceId != 0xA95F1F8A96159516 {
		t.Fail()
	}

	// different name
	if GenerateNamespaceId("symbol", 0) == GenerateNamespaceId("Symbol", 0) {
		t.Fail()
	}

	// different parent
	if GenerateNamespaceId("symbol", 0xA95F1F8A96159516) == GenerateNamespaceId("symbol", 0xA95F1F8A96159517) {
		t.Fail()
	}

	for i := uint64(1); i < 10; i++ {
		namespaceId = GenerateNamespaceId("symbol", i)
		if namespaceId>>63 != 1 {
			t.Fail()
		}
	}
}

func TestGenerateMosaicAliasId(t *testing.T) {
	aliasId, err := GenerateMosaicAliasId("cat.token")
	if err != nil {
		t.Fail()
	}
	if aliasId != 0xA029E100621B2E33 {
		t.Fail()
	}
}

func TestGenerateNamespacePath(t *testing.T) {
	path, err := GenerateNamespacePath("cat")
	if err != nil {
		t.Fail()
	}
	if !utils.ListEqual(path, []uint64{0xB1497F5FBA651B4F}) {
		t.Fail()
	}

	path, err = GenerateNamespacePath("cat.token")
	if err != nil {
		t.Fail()
	}
	if !utils.ListEqual(path, []uint64{0xB1497F5FBA651B4F, 0xA029E100621B2E33}) {
		t.Fail()
	}
}

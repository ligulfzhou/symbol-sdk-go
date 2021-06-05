package bip

import (
	"bytes"
	"strings"
	"symbol-sdk-go/utils"
	"testing"

	"github.com/brianium/mnemonic"
)

const DETERIMINISTIC_MNEMONIC = "cat swing flag economy stadium alone churn speed unique patch report train"

var DETERIMINISTIC_SEED = utils.Unhexify("000102030405060708090A0B0C0D0E0F")

func TestBip32_NodeFromSeed(t *testing.T) {

	bip32 := NewDefaultBip32()
	node := bip32.NodeFromSeed(DETERIMINISTIC_SEED)
	checkBip32Node(t, node,
		utils.Unhexify("90046A93DE5380A72B5E45010748567D5EA02BBF6522F979E05C0D8D8CA9FFFB"),
		utils.Unhexify("2B4BE7F19EE27BBF30C667B642D5F4AA69FD169872F8FC3059C08EBAE2EB19E7"))
}

func TestBip32_NodeFromSeed2(t *testing.T) {
	bip32 := NewBip32("ed25519-keccak", string(mnemonic.English))
	node := bip32.NodeFromSeed(DETERIMINISTIC_SEED)
	checkBip32Node(t, node,
		utils.Unhexify("9CFCA256458AAC0A0550A30DC7639D87364E4323BA61ED41454818E3317BAED0"),
		utils.Unhexify("A3D76D92ACF784D68F4EA2F6DE5507A3520385237A80277132B6C8F3685601B2"))
}

func TestBip32Node_DeriveOne(t *testing.T) {

	bip32 := NewDefaultBip32()
	node := bip32.NodeFromSeed(DETERIMINISTIC_SEED)
	derivedNode := node.DeriveOne(0)
	checkBip32Node(t, derivedNode,
		utils.Unhexify("8B59AA11380B624E81507A27FEDDA59FEA6D0B779A778918A2FD3590E16E9C69"),
		utils.Unhexify("68E0FE46DFB67E368C75379ACEC591DAD19DF3CDE26E63B93A8E704F1DADE7A3"),
	)
}

func TestBip32Node_DerivePath(t *testing.T) {
	bip32 := NewDefaultBip32()
	rootNode := bip32.NodeFromSeed(DETERIMINISTIC_SEED)
	node0 := rootNode.DeriveOne(44).DeriveOne(4343).DeriveOne(0).DeriveOne(0).DeriveOne(0)
	node1 := rootNode.DeriveOne(44).DeriveOne(4343).DeriveOne(1).DeriveOne(0).DeriveOne(0)

	checkBip32Node(t, node0,
		utils.Unhexify("B8E16D407C8837B46A9445C6417310F3C7A4DCD9B8FF2679C383E6DEF721AC11"),
		utils.Unhexify("BB2724A538CFD64E4366FEB36BB982B954D58EA78F7163451B3B514EDD692159"),
	)

	checkBip32Node(t, node1,
		utils.Unhexify("68CA2A058611AAC20CAFB4E1CCD70961E67D8C567390B3CBFC63D0E58BAE7153"),
		utils.Unhexify("8C91D9F5D214A2E80A275E75A165F7022712F7AD52B7ECD45B3B6CC76154B571"),
	)
}

func TestBip32Node_DerivePath2(t *testing.T) {
	bip32 := NewDefaultBip32()
	rootNode := bip32.NodeFromSeed(DETERIMINISTIC_SEED)
	node0 := rootNode.DerivePath([]int{44, 4343, 0, 0, 0})
	node1 := rootNode.DerivePath([]int{44, 4343, 1, 0, 0})

	checkBip32Node(t, node0,
		utils.Unhexify("B8E16D407C8837B46A9445C6417310F3C7A4DCD9B8FF2679C383E6DEF721AC11"),
		utils.Unhexify("BB2724A538CFD64E4366FEB36BB982B954D58EA78F7163451B3B514EDD692159"),
	)

	checkBip32Node(t, node1,
		utils.Unhexify("68CA2A058611AAC20CAFB4E1CCD70961E67D8C567390B3CBFC63D0E58BAE7153"),
		utils.Unhexify("8C91D9F5D214A2E80A275E75A165F7022712F7AD52B7ECD45B3B6CC76154B571"),
	)
}

func TestBip32_NodeFromMnemonic(t *testing.T) {
	bip32 := NewDefaultBip32()
	rootNode := bip32.NodeFromMnemonic(strings.Split(DETERIMINISTIC_MNEMONIC, " "), "TREZOR")
	node0 := rootNode.DerivePath([]int{44, 4343, 0, 0, 0})
	node1 := rootNode.DerivePath([]int{44, 4343, 1, 0, 0})
	node2 := rootNode.DerivePath([]int{44, 4343, 2, 0, 0})

	if !bytes.Equal(node0.PrivateKey, utils.Unhexify("1455FB18AB105444763EED593B7CA1C53EF6DDF8BDA1AB7004276FEAB1FCF222")) {
		t.Fail()
	}
	if !bytes.Equal(node1.PrivateKey, utils.Unhexify("913967B3DFE1E94C50D5C92789DA194644D2A699E5BB75B171A3B68993B82A21")) {
		t.Fail()
	}
	if !bytes.Equal(node2.PrivateKey, utils.Unhexify("AEC7C0143FC11F26FF5DB020492DACA7C8CF2640D2377AD3C721286472571602")) {
		t.Fail()
	}
}
func checkBip32Node(t *testing.T, node *Bip32Node, expectedChainCode []byte, expectedPrivateKey []byte) {
	if !bytes.Equal(node.ChainCode, expectedChainCode) {
		t.Fail()
	}

	if !bytes.Equal(node.PrivateKey, expectedPrivateKey) {
		t.Fail()
	}
}

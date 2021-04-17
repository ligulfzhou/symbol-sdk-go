package sym

import (
	"symbol-sdk-go/core/utils"
)

type MerkleHashBuilder struct {
	Hashes [][]byte
}

func (mhb MerkleHashBuilder) Update(componentHash string) {
	mhb.Hashes = append(mhb.Hashes, utils.Unhexify(componentHash))
}

func (mhb MerkleHashBuilder) Final() {

	if len(mhb.Hashes) == 0 {
		return
	}
}

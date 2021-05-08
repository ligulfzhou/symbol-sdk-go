package crypto

import (
	"math"
	"symbol-sdk-go/utils"

	"golang.org/x/crypto/sha3"
)

type MerkleHashBuilder struct {
	Hashes [][]byte
}

func NewMerkleHashBuilder() *MerkleHashBuilder {
	return &MerkleHashBuilder{
		Hashes: [][]byte{},
	}
}

func (mhb *MerkleHashBuilder) Update(componentHash []byte) {
	mhb.Hashes = append(mhb.Hashes, componentHash)
}

func (mhb *MerkleHashBuilder) Final() []byte {
	if len(mhb.Hashes) == 0 {
		return utils.Hash256Zero
	}

	numRemainingHashes := len(mhb.Hashes)
	for numRemainingHashes > 1 {
		idx := 0

		for idx < numRemainingHashes {
			hasher := sha3.New256()
			hasher.Write(mhb.Hashes[idx])

			if idx+1 < numRemainingHashes {
				hasher.Write(mhb.Hashes[idx+1])
			} else {
				hasher.Write(mhb.Hashes[idx])
				numRemainingHashes += 1
			}
			mhb.Hashes[int(math.Floor(float64(idx)/2))] = hasher.Sum(nil)
			idx += 2
		}
		numRemainingHashes /= 2
	}
	return mhb.Hashes[0]
}

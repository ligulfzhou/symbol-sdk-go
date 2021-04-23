package sym

import (
	"fmt"
	"math"
	"symbol-sdk-go/core"

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

func (mhb *MerkleHashBuilder) Update(componentHash interface{}) {
	switch componentHash := componentHash.(type) {
	case string:
		mhb.Hashes = append(mhb.Hashes, core.Unhexify(componentHash))
	case []byte:
		mhb.Hashes = append(mhb.Hashes, componentHash)
	default:
	}
	fmt.Println(len(mhb.Hashes))
}

func (mhb *MerkleHashBuilder) Final() *core.Hash256 {
	if len(mhb.Hashes) == 0 {
		return core.Hash256Zero
	}

	numRemainingHashes := len(mhb.Hashes)
	for numRemainingHashes > 1 {
		idx := 0

		for idx < numRemainingHashes {
			hasher := sha3.New256()
			fmt.Println(idx)
			hasher.Write(mhb.Hashes[idx])

			if idx+1 < numRemainingHashes {
				fmt.Println(idx + 1)
				hasher.Write(mhb.Hashes[idx+1])
			} else {
				fmt.Println(idx)
				hasher.Write(mhb.Hashes[idx])
				numRemainingHashes += 1
			}
			mhb.Hashes[int(math.Floor(float64(idx)/2))] = hasher.Sum(nil)
			idx += 2

			// mhb.Print()
		}
		numRemainingHashes /= 2
	}
	return core.NewHash256(core.Hexify(mhb.Hashes[0]))
}

// func (mhb MerkleHashBuilder) Print() {
// 	fmt.Println("==")
// 	for _, hash := range mhb.Hashes {
// 		fmt.Println(core.Hexify(hash))
// 	}
// }

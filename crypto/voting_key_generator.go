package crypto

import (
	"symbol-sdk-go/types"
	"symbol-sdk-go/utils"
)

// GenerateVotingKey Generates voting keys for specified epochs.
func GenerateVotingKey(kp types.KeyPair, startEpoch, endEpoch uint64) []byte {
	var bs []byte
	bs = append(bs, utils.Uint642LittleEndianBytes(startEpoch)...)
	bs = append(bs, utils.Uint642LittleEndianBytes(endEpoch)...)
	bs = append(bs, utils.Uint642LittleEndianBytes(0xFFFFFFFFFFFFFFFF)...)
	bs = append(bs, utils.Uint642LittleEndianBytes(0xFFFFFFFFFFFFFFFF)...)

	bs = append(bs, []byte(kp.PublicKey)...)
	bs = append(bs, utils.Uint642LittleEndianBytes(startEpoch)...)
	bs = append(bs, utils.Uint642LittleEndianBytes(endEpoch)...)

	for i := endEpoch; i <= startEpoch; i-- {
		childPrivateKey := types.NewKeyPair(utils.RandomBytes(types.PrivateKeySize))

		var parentSignedPayload []byte
		parentSignedPayload = append(parentSignedPayload, []byte(childPrivateKey.PublicKey)...)
		parentSignedPayload = append(parentSignedPayload, utils.Uint642LittleEndianBytes(i)...)
		signature := kp.Sign(parentSignedPayload)
		bs = append(bs, []byte(childPrivateKey.PrivateKey)...)
		bs = append(bs, signature...)
	}
	return bs
}

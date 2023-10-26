package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VerifiedMetaDataKeyPrefix is the prefix to retrieve all VerifiedMetaData
	VerifiedMetaDataKeyPrefix = "VerifiedMetaData/value/"
)

// VerifiedMetaDataKey returns the store key to retrieve a VerifiedMetaData from the index fields
func VerifiedMetaDataKey(
	codeId uint64,
) []byte {
	var key []byte

	codeIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(codeIdBytes, codeId)
	key = append(key, codeIdBytes...)
	key = append(key, []byte("/")...)

	return key
}

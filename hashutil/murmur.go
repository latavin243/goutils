package hashutil

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/twmb/murmur3"
)

func Murmur128Hex(s string) string {
	h1, h2 := murmur3.StringSum128(s)
	byteSlice := convertUint64ToBytes(h1)
	byteSlice = append(byteSlice, convertUint64ToBytes(h2)...)
	return hex.EncodeToString(byteSlice)
}

func convertUint64ToBytes(i uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, i)
	return buf
}

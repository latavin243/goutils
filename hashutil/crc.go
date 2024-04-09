package hashutil

import (
	"hash/crc32"
	"hash/crc64"
)

const ISO = 0xD8000000

func Crc32(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

func Crc64(s string) uint64 {
	return crc64.Checksum([]byte(s), crc64.MakeTable(ISO))
}

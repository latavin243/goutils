package hashutil

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Sha1(s string) string {
	binHash := sha1.Sum([]byte(s))
	return hex.EncodeToString(binHash[:])
}

func Sha256(s string) string {
	binHash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(binHash[:])
}

func Sha512(s string) string {
	binHash := sha512.Sum512([]byte(s))
	return hex.EncodeToString(binHash[:])
}

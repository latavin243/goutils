package hashutil

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(in string) string {
	binHash := md5.Sum([]byte(in))
	return hex.EncodeToString(binHash[:])
}

package hashutil

import (
	"strconv"

	"github.com/spaolacci/murmur3"
)

func Murmur3(s string) string {
	return strconv.FormatUint(murmur3.Sum64([]byte(s)), 16)
}

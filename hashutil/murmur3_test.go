package hashutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/hashutil"
)

func TestMurmur128Hex(t *testing.T) {
	var (
		rawKey   = "1_2_3_hello_world"
		expected = "4d75b5c36372c97a"
	)
	res := hashutil.Murmur3(rawKey)
	fmt.Printf("Murmur3: %s\n", res)
	assert.Equal(t, expected, res)
}

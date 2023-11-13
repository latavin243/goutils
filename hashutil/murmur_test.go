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
		expected = "7ac97263c3b5754d7c99b31b980eb88d"
	)
	res := hashutil.Murmur128Hex(rawKey)
	fmt.Printf("Murmur128Hex: %s\n", res)
	assert.Equal(t, expected, res)
}

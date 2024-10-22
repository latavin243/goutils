package hashutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/hashutil"
)

func TestCrc(t *testing.T) {
	rawText := "1_2_3_hello_world"
	{
		expected := uint32(0x4c530aed)
		res := hashutil.Crc32(rawText)
		fmt.Printf("Crc32: %d\n", res)
		assert.Equal(t, expected, res)
	}
	{
		expected := uint64(0xffffffff5e2405e6)
		res := hashutil.Crc64(rawText)
		fmt.Printf("Crc64: %d\n", res)
		assert.Equal(t, expected, res)
	}
}

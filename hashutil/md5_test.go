package hashutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/hashutil"
)

func TestMd5(t *testing.T) {
	var (
		rawText  = "1_2_3_hello_world"
		expected = "53a3ca2683cd32de63528122ef7b5b2e"
	)
	res := hashutil.Md5(rawText)
	fmt.Printf("Md5: %s\n", res)
	assert.Equal(t, expected, res)
}

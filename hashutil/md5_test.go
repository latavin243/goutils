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
		expected = "22829b82f32491a6cdd894f16198a9f4"
	)
	res := hashutil.Md5(rawText)
	fmt.Printf("Md5: %s\n", res)
	assert.Equal(t, expected, res)
}

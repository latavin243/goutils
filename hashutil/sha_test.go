package hashutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/hashutil"
)

func TestSha(t *testing.T) {
	rawText := "1_2_3_hello_world"
	{
		expected := "4ac73721de3022c50c954910faa2b2826737c204"
		res := hashutil.Sha1(rawText)
		fmt.Printf("Sha1: %s\n", res)
		assert.Equal(t, expected, res)
	}
	{
		expected := "6ce901f6df48826d4a4431215f2fa18c15d0a98179ad68561d96da1694eb1267"
		res := hashutil.Sha256(rawText)
		fmt.Printf("Sha256: %s\n", res)
		assert.Equal(t, expected, res)
	}
	{
		expected := "0d8b5aa167fd9b041b3a067f2755a19849fc340a772d5b627354fa06eb27c625c3901c1267a1e53f3d5e1fab667aea1e1d406afae1b8781d5da7f44cce8c08cf"
		res := hashutil.Sha512(rawText)
		fmt.Printf("Sha512: %s\n", res)
		assert.Equal(t, expected, res)
	}
}

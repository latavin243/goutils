package respcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/respcode"
)

func TestRespCodeError(t *testing.T) {
	errDetail := "invalid userid"
	resp := respcode.CodeBusinessErr.ToError(errDetail)
	fmt.Printf("resp detail: %s\n", resp)
	assert.Equal(t, resp.Error(), errDetail)
}

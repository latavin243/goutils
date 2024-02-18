package number_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/number"
)

func TestNumToAccountingString(t *testing.T) {
	for _, tc := range []struct {
		input    float64
		decimals int
		expected string
	}{
		{1234.567, 2, "1,234.57"},
		{1234, 2, "1,234.00"},
	} {
		assert.Equal(t, tc.expected, number.NumToAccountingString(tc.input, tc.decimals))
	}
}

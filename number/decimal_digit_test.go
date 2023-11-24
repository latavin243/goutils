package number_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/number"
)

func TestDecimalDigits(t *testing.T) {
	testCases := []struct{ input, expectDigits int }{
		{12345, 5},
		{-12345, 5},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expectDigits, number.DecimalDigits(tc.input))
	}
}

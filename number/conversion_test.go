package number_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/number"
)

func TestConversion(t *testing.T) {
	assert.Equal(t, "123", number.IntToString(123))
	assert.Equal(t, "123.46", number.FloatToString(123.456, 2))
	assert.Equal(t, "1,234.57", number.NumToAccountingString(1234.567, 2))
	assert.Equal(t, "1,234.00", number.NumToAccountingString(1234, 2))
	{
		res, err := number.StringToInt[int64]("123")
		assert.NoError(t, err)
		assert.Equal(t, int64(123), res)
	}
	{
		res, err := number.StringToFloat[float64]("123.45")
		assert.NoError(t, err)
		assert.Equal(t, 123.45, res)
	}
}

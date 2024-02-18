package number_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/number"
)

func TestIntToString(t *testing.T) {
	for _, tc := range []struct {
		input    int
		expected string
	}{
		{123, "123"},
		{-123, "-123"},
	} {
		assert.Equal(t, tc.expected, number.IntToString(tc.input))
	}
}

func TestFloatToString(t *testing.T) {
	for _, tc := range []struct {
		input    float64
		decimals int
		expected string
	}{
		{123.456, 2, "123.46"},
		{-123.456, 1, "-123.5"},
	} {
		assert.Equal(t, tc.expected, number.FloatToString(tc.input, tc.decimals))
	}
}

func TestStringToInt(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected int64
	}{
		{"123", 123},
		{"-123", -123},
	} {
		res, err := number.StringToInt[int64](tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}

func TestStringToFloat(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected float64
	}{
		{"123.45", 123.45},
		{"-123.45", -123.45},
	} {
		res, err := number.StringToFloat[float64](tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, res)
	}
}

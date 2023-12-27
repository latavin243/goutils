package strfill_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/strfill"
)

func TestStrFill(t *testing.T) {
	type Record struct {
		Date int32 // yyyyMMdd, e.g. 20060102
	}
	tmpl := "my_table_{{yyyyMM .Date}}"
	refRecord := Record{20060102}
	res, err := strfill.StrFill(tmpl, refRecord)
	assert.NoError(t, err)
	assert.Equal(t, "my_table_200601", res)
}

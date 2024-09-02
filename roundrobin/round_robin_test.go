package roundrobin_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/latavin243/goutils/roundrobin"
)

type resource struct {
	id   int
	name string
}

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		resources []*resource
		iserr     bool
		expected  []string
		want      []*resource
	}{
		{
			resources: []*resource{
				{1, "resource-1"},
				{2, "resource-2"},
				{3, "resource-3"},
				{4, "resource-4"},
				{5, "resource-5"},
				{6, "resource-6"},
				{7, "resource-7"},
			},
			iserr: false,
			want: []*resource{
				{1, "resource-1"},
				{2, "resource-2"},
				{3, "resource-3"},
				{4, "resource-4"},
				{5, "resource-5"},
				{6, "resource-6"},
				{7, "resource-7"},
				{1, "resource-1"},
			},
		},
		{
			resources: []*resource{},
			iserr:     true,
			want:      []*resource{},
		},
	}

	for i, test := range tests {
		rr, err := roundrobin.New(test.resources...)
		if test.iserr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		gots := make([]*resource, 0, len(test.want))
		for j := 0; j < len(test.want); j++ {
			gots = append(gots, rr.Next())
		}

		if got, want := gots, test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("tests[%d] - RoundRobin is wrong. want: %v, got: %v", i, want, got)
		}
	}
}

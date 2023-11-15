package reflectutil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/latavin243/goutils/reflectutil"
)

type sampleStruct struct {
	Alice   int    `form:"f1"`
	Bob     string `form:"f2"`
	Charlie bool   `form:"f3"`
}

var sample = &sampleStruct{
	Alice:   1,
	Bob:     "hello",
	Charlie: true,
}

func TestCollect(t *testing.T) {
	type args struct {
		inputStruct interface{}
		tagNames    []string
	}
	tests := []struct {
		name                string
		args                args
		wantFieldTagEntries []*reflectutil.FieldTagEntry
		wantErr             bool
	}{
		{
			"test CollectTags()",
			args{sample, []string{"form"}},
			[]*reflectutil.FieldTagEntry{
				{"Alice", 1, map[string]string{"form": "f1"}},
				{"Bob", "hello", map[string]string{"form": "f2"}},
				{"Charlie", true, map[string]string{"form": "f3"}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFieldTagEntries, err := reflectutil.CollectTags(tt.args.inputStruct, tt.args.tagNames)
			if (err != nil) != tt.wantErr {
				t.Errorf("CollectTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFieldTagEntries, tt.wantFieldTagEntries) {
				t.Errorf("CollectTags() = %v, want %v", gotFieldTagEntries, tt.wantFieldTagEntries)
			}
			for _, entry := range gotFieldTagEntries {
				fmt.Println(entry)
			}
		})
	}
}

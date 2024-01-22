package filereader

import (
	"fmt"
	"path/filepath"
	"reflect"
)

func ReadFile(filePath string, target interface{}) (err error) {
	if err := checkTarget(target); err != nil {
		return err
	}

	ext := filepath.Ext(filePath)
	ft, supported := fileExtTypeMap[ext]
	if !supported {
		return fmt.Errorf("unsupported file type, filePath=%s", filePath)
	}

	parser, ok := parserMap[ft]
	if !ok {
		return fmt.Errorf("file parser not found, filePath=%s", filePath)
	}
	return parser(filePath, target)
}

func checkTarget(input interface{}) error {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target is not a pointer")
	}
	if v.IsNil() {
		return fmt.Errorf("target is nil")
	}
	return nil
}

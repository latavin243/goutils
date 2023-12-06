package filereader

import (
	"fmt"
	"path/filepath"
)

func ReadFile(filePath string, target interface{}) (err error) {
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

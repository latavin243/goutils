package filereader

import (
	"fmt"
	"path/filepath"
)

func ReadFile(filePath string, target interface{}) (err error) {
	ft, isValid := getFileType(filePath)
	if !isValid {
		return fmt.Errorf("unsupported file type, filePath=%s", filePath)
	}

	parser, ok := parserMap[ft]
	if !ok {
		return fmt.Errorf("file parser not found, filePath=%s", filePath)
	}
	return parser(filePath, target)
}

func getFileType(filePath string) (ft fileType, isValid bool) {
	ext := filepath.Ext(filePath)
	ft, ok := fileExtTypeMap[ext]
	if !ok {
		return fileTypeInvalid, false
	}
	return ft, true
}

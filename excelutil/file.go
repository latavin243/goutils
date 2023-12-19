package excelutil

import (
	"bytes"
	"errors"

	"github.com/tealeg/xlsx"
)

type File interface {
	AddSheet(name string, sheetDir SheetDir, titleStyle *xlsx.Style) (Sheet, error)
	ToBytes() ([]byte, error)
}

type XlsxFile struct {
	File     *xlsx.File
	FileName string
}

func NewExcelFile(
	fileName string,
) (File, error) {
	return &XlsxFile{
		File:     xlsx.NewFile(),
		FileName: fileName,
	}, nil
}

func (f *XlsxFile) ToBytes() ([]byte, error) {
	if f == nil || f.File == nil {
		return nil, errors.New("file is nil pointer")
	}
	buffer := &bytes.Buffer{}
	err := f.File.Write(buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (f *XlsxFile) AddSheet(
	name string, sheetDir SheetDir, titleStyle *xlsx.Style,
) (Sheet, error) {
	sheetName := name
	if len(sheetName) > sheetNameTruncateLength+3 {
		sheetName = sheetName[0:sheetNameTruncateLength] + "..."
	}
	sheet, err := f.File.AddSheet(sheetName)
	if err != nil {
		return nil, err
	}
	excelSheet := newXlsxSheet(sheet, sheetName, sheetDir, titleStyle)
	return excelSheet, nil
}

package excelutil_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/latavin243/goutils/excelutil"
)

func DullTranslate(key string) string {
	return fmt.Sprintf("%s_translated", key)
}

// LR: left-right
func TestLRExcelFileExport(t *testing.T) {
	fileName := "test_lr.xlsx"
	filePath := filepath.Join("/tmp", fileName)

	file, err := NewExcelFile(fileName)
	assert.NoError(t, err)
	sheet, err := file.AddSheet("sheet_one", SheetDirLR, nil)
	assert.NoError(t, err)
	sheet.SetWidths([]uint32{12, 12, 12, 12})

	mainTitles := []*Title{
		NewCustomTitle(""),
		NewCustomTitle("main_title_1").MergeCell(2),
		NewCustomTitle("main_title_2"),
	}
	subTitles := []*Title{
		NewCustomTitle("sub_title_1"),
		NewCustomTitle("sub_title_2").Bold(false),
		NewCustomTitle("sub_title_3"),
		NewCustomTitle("sub_title_4").Bold(false),
	}
	dataEntry1 := []*Cell{
		NewCell("1234.56"),
		NewCell("12.23%"),
		NewCell("123"),
		NewCell("46.64%"),
	}

	sheet.AddTitle(mainTitles)
	sheet.AddTitle(subTitles)
	sheet.AddEntry(dataEntry1, nil)

	bytes, err := file.ToBytes()
	assert.NoError(t, err)
	err = os.WriteFile(filePath, bytes, 0o644)
	assert.NoError(t, err)
}

// TB: top-bottom
func TestTBExcelFileExport(t *testing.T) {
	fileName := "test_tb.xlsx"
	filePath := filepath.Join("/tmp", fileName)

	file, err := NewExcelFile(fileName)
	assert.NoError(t, err)
	sheet, err := file.AddSheet("sheet_one", SheetDirTB, nil)
	assert.NoError(t, err)
	sheet.SetWidths([]uint32{12, 12, 12, 12})

	mainTitles := []*Title{
		NewCustomTitle(""),
		NewCustomTitle("main_title_1").MergeCell(2),
		NewCustomTitle("main_title_2"),
	}
	subTitles := []*Title{
		NewCustomTitle("sub_title_1"),
		NewCustomTitle("sub_title_2").Bold(false),
		NewCustomTitle("sub_title_3"),
		NewCustomTitle("sub_title_4").Bold(false),
	}
	dataEntry1 := []*Cell{
		NewCell("1234.56"),
		NewCell("12.23%"),
		NewCell("123"),
		NewCell("46.64%"),
	}
	dataEntry2 := []*Cell{
		NewCell("1235.56"),
		NewCell("13.23%"),
		NewCell("124"),
		NewCell("47.64%"),
	}

	sheet.AddTitle(mainTitles)
	sheet.AddTitle(subTitles)
	sheet.AddEntry(dataEntry1, nil)
	sheet.AddEntry(dataEntry2, nil)

	bytes, err := file.ToBytes()
	assert.NoError(t, err)
	err = os.WriteFile(filePath, bytes, 0o644)
	assert.NoError(t, err)
}

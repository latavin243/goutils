package excelutil

import (
	"sync/atomic"

	"github.com/tealeg/xlsx"
)

type Sheet interface {
	SetWidths(widths []uint32)
	AddTitle(titles []*Title)
	AddEntry(cells []*Cell, customStyle *xlsx.Style)
}

type XlsxSheet struct {
	Sheet     *xlsx.Sheet
	SheetName string
	Direction SheetDir

	TitleStyle *xlsx.Style

	ColCount uint32
	RowCount uint32
}

func newXlsxSheet(
	sheet *xlsx.Sheet, sheetName string, sheetDir SheetDir,
	titleStyle *xlsx.Style,
) Sheet {
	if sheetDir != SheetDirTB && sheetDir != SheetDirLR {
		sheetDir = SheetDirLR
	}
	if titleStyle == nil {
		titleStyle = GetDefaultTitleStyle(true, xlsx.Alignment{
			WrapText:   true,
			Horizontal: "center",
			Vertical:   "center",
		})
	}
	return &XlsxSheet{
		Sheet:      sheet,
		SheetName:  sheetName,
		Direction:  sheetDir,
		TitleStyle: titleStyle,
		ColCount:   0,
		RowCount:   0,
	}
}

func (s *XlsxSheet) SetWidths(widths []uint32) {
	for index, width := range widths {
		_ = s.Sheet.SetColWidth(index, index+1, float64(width))
	}
}

func (s *XlsxSheet) AddTitle(titles []*Title) {
	switch s.Direction {
	case SheetDirLR:
		s.addTitleRow(titles)
	case SheetDirTB:
		s.addTitleCol(titles)
	}
}

func (s *XlsxSheet) AddEntry(cells []*Cell, customStyle *xlsx.Style) {
	switch s.Direction {
	case SheetDirLR:
		s.addDataRow(cells, customStyle)
	case SheetDirTB:
		s.addDataCol(cells, customStyle)
	}
}

func (s *XlsxSheet) addTitleRow(titles []*Title) {
	rowCount := atomic.AddUint32(&s.RowCount, 1) - 1

	colCount := 0
	for _, title := range titles {
		if title == nil {
			continue
		}
		cell := s.Sheet.Cell(int(rowCount), colCount)
		// merge cell
		cellMergeNum := max(1, int(title.mergeCellNumber)) // at least 1
		if cellMergeNum == 1 {
			colCount++
		} else {
			colCount += cellMergeNum
			cell.Merge(cellMergeNum-1, 0)
		}
		s.fillTitleCell(cell, title)
	}
}

func (s *XlsxSheet) addDataRow(cells []*Cell, customStyle *xlsx.Style) {
	atomic.AddUint32(&s.RowCount, 1)

	row := s.Sheet.AddRow()
	for _, cell := range cells {
		if cell == nil {
			continue
		}
		currentCell := row.AddCell()
		if customStyle != nil {
			customStyle.Alignment = cell.alignment
			currentCell.SetStyle(customStyle)
		} else {
			currentCell.SetStyle(GetDefaultDataStyle(cell.isBold, cell.alignment))
		}
		currentCell.Value = cell.content
	}
}

func (s *XlsxSheet) addTitleCol(titles []*Title) {
	colCount := atomic.AddUint32(&s.ColCount, 1) - 1

	rowCount := 0
	for _, title := range titles {
		if title == nil {
			continue
		}
		cell := s.Sheet.Cell(rowCount, int(colCount))
		// merge cell
		cellMergeNum := max(1, int(title.mergeCellNumber)) // at least 1
		if cellMergeNum == 1 {
			rowCount++
		} else {
			rowCount += cellMergeNum
			cell.Merge(0, cellMergeNum-1)
		}
		s.fillTitleCell(cell, title)
	}
}

func (s *XlsxSheet) addDataCol(cells []*Cell, customStyle *xlsx.Style) {
	colCount := atomic.AddUint32(&s.ColCount, 1) - 1

	for rowCount, cell := range cells {
		if cell == nil {
			continue
		}
		// start from row 0 in tealeg/xlsx package
		currentCell := s.Sheet.Cell(rowCount, int(colCount))
		if customStyle != nil {
			customStyle.Alignment = cell.alignment
			currentCell.SetStyle(customStyle)
		} else {
			currentCell.SetStyle(GetDefaultDataStyle(cell.isBold, cell.alignment))
		}
		currentCell.Value = cell.content
	}
}

func (s *XlsxSheet) fillTitleCell(cell *xlsx.Cell, title *Title) {
	titleStr := title.title
	cell.SetStyle(s.TitleStyle)
	cell.Value = titleStr
}

package excelutil

import "github.com/tealeg/xlsx"

type Title struct {
	title           string
	isBold          bool
	mergeCellNumber uint
	fill            *xlsx.Fill
	alignment       xlsx.Alignment
}

func NewCustomTitle(title string) *Title {
	return &Title{
		title:           title,
		isBold:          true, // default bold
		mergeCellNumber: 1,
		fill:            nil,
		alignment: xlsx.Alignment{
			WrapText:   true,
			Horizontal: "center",
			Vertical:   "center",
		},
	}
}

func (t *Title) Bold(isBold bool) *Title {
	t.isBold = isBold
	return t
}

func (t *Title) MergeCell(mergeCellNum uint) *Title {
	t.mergeCellNumber = mergeCellNum
	return t
}

func (t *Title) WithBackgroundFill(fill *xlsx.Fill) *Title {
	t.fill = fill
	return t
}

func (t *Title) WrapText(wrapText bool) *Title {
	t.alignment.WrapText = wrapText
	return t
}

func (t *Title) HorizontalAlignment(alignment Aliment) *Title {
	switch alignment {
	case AlimentLeft, AlimentRight, AimentCenter:
		t.alignment.Horizontal = alignment.String()
	default:
		// do nothing
	}
	return t
}

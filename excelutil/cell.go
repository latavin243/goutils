package excelutil

import "github.com/tealeg/xlsx"

type Cell struct {
	content   string
	isBold    bool
	alignment xlsx.Alignment
}

func NewCell(content string) *Cell {
	return &Cell{
		content: content,
		isBold:  false,
		alignment: xlsx.Alignment{
			WrapText:   true,
			Horizontal: "right",
			Vertical:   "center",
		},
	}
}

func (c *Cell) Bold(isBold bool) *Cell {
	c.isBold = isBold
	return c
}

func (c *Cell) HorizontalAlignment(alignment Aliment) *Cell {
	switch alignment {
	case AlimentLeft, AlimentRight, AimentCenter:
		c.alignment.Horizontal = alignment.String()
	default:
		// do nothing
	}
	return c
}

func (c *Cell) WrapText(wrapText bool) *Cell {
	c.alignment.WrapText = wrapText
	return c
}

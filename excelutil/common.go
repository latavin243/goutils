package excelutil

import "github.com/tealeg/xlsx"

const (
	exportTagName           = "export"
	sheetNameTruncateLength = 27 // turn long name into "xxx...", threshold 30
)

type SheetDir int

const (
	SheetDirTB SheetDir = iota + 1 // top-bottom direction
	SheetDirLR                     // left-right direction
)

type Aliment int

const (
	AlimentLeft Aliment = iota + 1
	AlimentRight
	AimentCenter
)

func (a Aliment) String() string {
	switch a {
	case AlimentLeft:
		return "left"
	case AlimentRight:
		return "right"
	case AimentCenter:
		return "center"
	default:
		return ""
	}
}

func GetDefaultTitleStyle(isBold bool, alignment xlsx.Alignment) *xlsx.Style {
	return &xlsx.Style{
		Border: xlsx.Border{
			Left: "thin", LeftColor: "808080",
			Right: "thin", RightColor: "808080",
			Top: "thin", TopColor: "808080",
			Bottom: "thin", BottomColor: "808080",
		},
		// Fill: xlsx.Fill{PatternType: "solid", BgColor: "160160160", FgColor: "160160160"},
		Font:      xlsx.Font{Bold: isBold, Size: 10, Name: "Arial"},
		Alignment: alignment,
	}
}

func GetDefaultDataStyle(isBold bool, alignment xlsx.Alignment) *xlsx.Style {
	return &xlsx.Style{
		Border: xlsx.Border{
			Left: "thin", LeftColor: "808080",
			Right: "thin", RightColor: "808080",
			Top: "thin", TopColor: "808080",
			Bottom: "thin", BottomColor: "808080",
		},
		// Fill: xlsx.Fill{PatternType: "solid", BgColor: "160160160", FgColor: "160160160"},
		Font:      xlsx.Font{Bold: isBold, Size: 10, Name: "Arial"},
		Alignment: alignment,
	}
}

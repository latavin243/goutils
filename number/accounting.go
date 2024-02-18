package number

import "github.com/leekchan/accounting"

// NumToAccountingString converts a number to an accounting string.
// Caution: This is an example, in practical some countries may have different accounting format.
func NumToAccountingString[T Num](f T, precision int) string {
	ac := &accounting.Accounting{
		Symbol:    "",
		Precision: precision,
		Thousand:  ",",
		Decimal:   ".",
	}
	return ac.FormatMoney(f)
}

package number

import "golang.org/x/exp/constraints"

func DecimalDigits[T constraints.Integer](num T) int {
	if num == 0 {
		return 1
	}
	if num < 0 {
		num = -num
	}
	var count int
	for num > 0 {
		count++
		num /= 10
	}
	return count
}

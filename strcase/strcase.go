package strcase

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z])([0-9A-Z])")
)

// CamelToSnake converts a camelCase string to snake case
func CamelToSnake(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// SnakeToTitle	converts a snake_case string to title case
func SnakeToTitle(str string) string {
	// strings.Title(strings.Replace(str, "_", " ", -1))
	return cases.Title(language.AmericanEnglish).String(strings.Replace(str, "_", " ", -1))
}

func ToUpperCaseSlice(strs []string) []string {
	var upperCaseStrs []string
	for _, str := range strs {
		upperCaseStrs = append(upperCaseStrs, strings.ToUpper(str))
	}
	return upperCaseStrs
}

func ToLowerCaseSlice(strs []string) []string {
	var lowerCaseStrs []string
	for _, str := range strs {
		lowerCaseStrs = append(lowerCaseStrs, strings.ToLower(str))
	}
	return lowerCaseStrs
}

// ReverseString reverses the input string
func ReverseString(input string) string {
	res := ""
	for _, ch := range input {
		res = string(ch) + res
	}
	return res
}

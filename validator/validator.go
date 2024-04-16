package validator

import (
	"encoding/json"
	"regexp"
	"sync"
)

var (
	patternCompileOnce sync.Once
	emailRegexp        *regexp.Regexp
)

func IsEmail(input string) bool {
	patternCompileOnce.Do(func() {
		emailRegexp, _ = regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	})
	return emailRegexp.MatchString(input)
}

func IsJSON(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}

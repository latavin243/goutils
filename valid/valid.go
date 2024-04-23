package valid

import (
	"encoding/json"
	"net"
	"regexp"
	"sync"
)

var (
	patternCompileOnce      sync.Once
	emailRegexp, httpRegexp *regexp.Regexp
)

func init() {
	emailRegexp, _ = regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// refer to https://stackoverflow.com/questions/3809401/what-is-a-good-regular-expression-to-match-a-url
	httpRegexp, _ = regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
}

func IsEmail(input string) bool {
	return emailRegexp.MatchString(input)
}

func IsJSON(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}

func IsIPv4(input string) bool {
	ip := net.ParseIP(input)
	return len(ip) == net.IPv4len
}

func IsIPv6(input string) bool {
	ip := net.ParseIP(input)
	return len(ip) == net.IPv6len
}

func IsHttpUrl(input string) bool {
	return httpRegexp.MatchString(input)
}

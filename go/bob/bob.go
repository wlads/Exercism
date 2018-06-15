package bob

import (
	"strings"
	"regexp"
)

func Hey(remark string) string {
	r := strings.TrimSpace(remark)

	switch {
	case isSilent(r):
		return "Fine. Be that way!"
	case isQuestion(r) && isYelling(r) && hasAlpha(r):
		return "Calm down, I know what I'm doing!"
	case isYelling(r) && hasAlpha(r):
		return "Whoa, chill out!"
	case isQuestion(r):
		return "Sure."
	default:
		return "Whatever."
	}
}

func isYelling(remark string) bool {
	return strings.ToUpper(remark) == remark
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isSilent(remark string) bool {
	return remark == ""
}

var isAlpha = regexp.MustCompile(`[[:alpha:]]+`)
func hasAlpha(s string) bool {
	return isAlpha.MatchString(s)
}

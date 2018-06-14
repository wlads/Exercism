package bob

import (
	"strings"
	"regexp"
)

func Hey(remark string) string {
	var allCaps = strings.ToUpper(remark) == remark
	var alpha = regexp.MustCompile(`[a-zA-Z]+`).MatchString

	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(remark, "?") {
		if allCaps && alpha(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	}

	if !alpha(remark) {
		return "Whatever."
	}

	if allCaps {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

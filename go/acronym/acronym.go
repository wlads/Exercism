package acronym

import (
	"strings"
	"regexp"
)

func Abbreviate(s string) string {
	words := regexp.MustCompile(`[\s-]`).Split(s, -1)
	abrev := ""
	for _, word := range words {
		abrev += word[:1]
	}
	return strings.ToUpper(abrev)
}

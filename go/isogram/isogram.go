// Package isogram from exercism.io
package isogram

import "strings"

// IsIsogram receives a string and check each rune if it is unique in the string
func IsIsogram(str string) bool {
	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "-", "", -1)
	for pos, char := range str {
		otherRunes := str[pos+1:]
		for _, otherChar := range otherRunes {
			if char == otherChar {
				return false
			}
		}
	}
	return true
}

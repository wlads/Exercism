// Package raindrops convert numbers to the sound of rain
package raindrops

import (
	"strings"
	"strconv"
)

// Convert receive a number and return a string following this pattern:
// If number is a factor of 3 returns "Pling"
// If number is a factor of 5 returns "Plang"
// If number is a factor of 7 returns "Plong"
func Convert(n int) string {
	var conv_n strings.Builder
	if n % 3 == 0 {
		conv_n.WriteString("Pling");
	}
	if n % 5 == 0 {
		conv_n.WriteString("Plang");
	}
	if n % 7 == 0 {
		conv_n.WriteString("Plong");
	}
	if conv_n.Len() == 0 {
		return strconv.Itoa(n)
	} else {
		return conv_n.String()
	}
}

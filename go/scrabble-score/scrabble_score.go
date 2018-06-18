// Package scrabble compute the scrabble score for words
package scrabble

import "unicode"

// Score receives a word and returns the sum of the value of each letter
func Score(str string) int {
	runes := []rune(str)
	sum := 0
	for i := 0; i < len(runes); i++ {
		sum += convertScore(runes[i])
	}
  return sum
}

// convertScore is case insensitive
// Letter                           Value
// A, E, I, O, U, L, N, R, S, T       1
// D, G                               2
// B, C, M, P                         3
// F, H, V, W, Y                      4
// K                                  5
// J, X                               8
// Q, Z                               10
func convertScore(r rune) int {
	switch(unicode.ToUpper(r)) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

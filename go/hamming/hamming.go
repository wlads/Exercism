// Package hamming from exercism.io
package hamming

import "errors"

// Calculate the hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("disallow strand different length")
	}

	distance := 0
	for i := 0; i<len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}

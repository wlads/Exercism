package grains

import (
	"errors"
	"math"
)

// Total number of wheat grains in a chessboard: 1, 2, 4, 8, ...
func Total() uint64 {
	return 18446744073709551615
}

// Square calculates the number of wheat grains in a position (pos)
func Square(pos int) (uint64, error) {
	if pos < 1 || pos > 64 {
		return 0, errors.New("position must be between 1 and 64")
	}
	grains := uint64(0)
	for i := 0; i < pos; i++ {
		grains = uint64(math.Pow(2, float64(i)))
	}
	return grains, nil
}

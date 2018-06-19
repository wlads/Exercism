package diffsquares

// SumOfSquares (1*1, 2*2, 3*3, ...)
func SumOfSquares(n int) int {
	sum := 0
	for i := range createArrayToNumber(n) {
		sum += i * i
	}
	return sum
}

// SquareOfSums (1 + 2 + ... + 10)2
func SquareOfSums(n int) int {
	sum := 0
	for i := range createArrayToNumber(n) {
		sum += i
	}
	return sum * sum
}

// Difference between the square of the sum and the sum of the squares of the first (n) natural numbers
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func createArrayToNumber(n int) []int {
	ary := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		ary[i] = i
	}
	return ary
}

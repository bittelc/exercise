package diffsquares

import (
	"math"
)

const testVersion = 1

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func SquareOfSums(n int) int {
	base := (((n + 1) * n) / 2)
	return base * base
}

func Difference(n int) int {
	diff := float64(SquareOfSums(n) - SumOfSquares(n))
	return int(math.Abs(diff))
}

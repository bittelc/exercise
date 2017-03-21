// Package diffsquares can calculate the sum of the square, the square of the sums, and the difference between the two
package diffsquares

import (
	"math"
)

const testVersion = 1

// SumOfSquares calculates Σ((n!)^2) for a provided n
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSums calculates (Σ(n!))^2 for a provdied n
func SquareOfSums(n int) int {
	base := (((n + 1) * n) / 2)
	return base * base
}

// Difference calculates the difference between the SumOfSquares and SquareOfSums for any n
func Difference(n int) int {
	diff := float64(SquareOfSums(n) - SumOfSquares(n))
	return int(math.Abs(diff))
}

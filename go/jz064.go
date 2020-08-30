package program

import "math"

func sumNums(n int) int {
	_ = n > 0 && plus(&n, sumNums(n-1))
	return n
}

func plus(a *int, b int) bool {
	*a += b
	return true
}

func sumNumsBetter(n int) int {
	return (int(math.Pow(float64(n), float64(2))) + n) >> 1
}

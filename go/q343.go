package program

import "math"

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	x := n / 3
	y := n % 3
	if y == 0 {
		return int(math.Pow(3, float64(x)))
	}
	if y == 1 {
		return int(math.Pow(3, float64(x-1)) * 4)
	}
	return int(math.Pow(3, float64(x))) * 2

}

package program

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var left, right = 1, x / 2

	for left < right {
		mid := left + (right-left)/2 + 1
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}

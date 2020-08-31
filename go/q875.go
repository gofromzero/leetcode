package program

import "math"

func minEatingSpeed(piles []int, H int) int {
	var left = 1
	var right = 0
	for _, v := range piles {
		if v > right {
			right = v
		}
	}
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if canEat(piles, mid, H) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func canEat(piles []int, speed int, H int) bool {
	var sum float64
	for _, v := range piles {
		sum += math.Ceil(float64(v) / float64(speed))
	}
	return sum > float64(H)
}

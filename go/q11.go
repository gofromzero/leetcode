package program

func maxArea(height []int) int {
	result := 0
	left, right := 0, len(height)-1
	for left < right {
		min := min(height[left], height[right])
		result = max(result, (right-left)*min)
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}
	return result
}

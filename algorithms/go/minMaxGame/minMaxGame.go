package minMaxGame

func minMaxGame(nums []int) int {
	l := len(nums)
	for {
		if l == 1 {
			break
		}
		l /= 2
		for i := 0; i < l; i++ {
			if i%2 == 0 {
				nums[i] = min(nums[2*i], nums[2*i+1])
			} else {
				nums[i] = max(nums[2*i], nums[2*i+1])
			}

		}
	}
	return nums[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

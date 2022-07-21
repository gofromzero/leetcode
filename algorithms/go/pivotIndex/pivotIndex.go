package pivotIndex

func pivotIndex(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	left := 0
	for i, val := range nums {
		if 2*left == total-val {
			return i
		}
		left += val
	}

	return -1
}

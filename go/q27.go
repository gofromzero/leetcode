package program

func removeElement(nums []int, val int) int {
	var flag int
	for _, v := range nums {
		if v != val {
			nums[flag] = v
			flag++
		}
	}
	return flag
}

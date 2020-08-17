package program

func removeDuplicates(nums []int) int {
	var old int
	if len(nums) < 1 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[old] {
			old++
			nums[old] = nums[i]
		}
	}
	return old + 1
}

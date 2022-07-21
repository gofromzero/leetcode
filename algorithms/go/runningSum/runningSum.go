package runningSum

func runningSum(nums []int) []int {
	var temp = nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += temp
		temp = nums[i]
	}
	return nums
}

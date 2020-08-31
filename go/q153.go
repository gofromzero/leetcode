package program

func findMin(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		var mid = left + (right-left)/2 + 1
		if nums[left] < nums[mid] {
			left = mid
		} else if nums[left] > nums[mid] {
			right = mid - 1
		}
	}
	return nums[(right+1)%len(nums)]
}

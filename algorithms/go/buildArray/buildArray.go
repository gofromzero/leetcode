// Package buildArray
// @Author Spume
// @Date 2025-05-06 11:11

package buildArray

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

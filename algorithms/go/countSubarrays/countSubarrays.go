// Package countSubarrays
// @Author Spume
// @Date 2025-04-27 09:35

package countSubarrays

func countSubarrays(nums []int) int {
	num := 0
	for i := 0; i+2 < len(nums); i++ {
		if (nums[i]+nums[i+2])*2 == nums[i+1] {
			num++
		}
	}
	return num
}

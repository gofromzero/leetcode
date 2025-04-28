// Package countSubarrays2302
// @Author Spume
// @Date 2025-04-28 09:13

package countSubarrays2302

func countSubarrays(nums []int, k int64) int64 {

	var num int64
	var total int64
	for l, r := 0, 0; r < len(nums); r++ {
		total += int64(nums[r])
		for l <= r && total*int64(r-l+1) >= k {
			total -= int64(nums[l])
			l++
		}

		num += int64(r - l + 1)
	}

	return num
}

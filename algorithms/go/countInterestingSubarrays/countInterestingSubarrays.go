// Package countInterestingSubarrays
// @Author Spume
// @Date 2025-04-25 09:25

package countInterestingSubarrays

import "fmt"

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {

	num := int64(0)
	for l := 0; l < len(nums); l++ {
		cnt := 0
		for r := l + 1; r <= len(nums); r++ {
			fmt.Println(nums[l:r])
			if nums[r]%modulo == k {
				cnt++
			}
		}
		if cnt%modulo == k {
			num++
		}
	}

	return num
}

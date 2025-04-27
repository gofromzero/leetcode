// Package countCompleteSubarrays
// @Author Spume
// @Date 2025-04-24 09:35

package countCompleteSubarrays

func countCompleteSubarrays(nums []int) int {
	var num int
	m := make([]int, 2001)

	for l := 0; l < len(nums); l++ {
		for _, key := range m {
			m[key] = 0
		}
		for r := l + 1; r <= len(nums); r++ {

			m[nums[r-1]]++
			if CheckMap(m) {
				num++
			}
		}
		m[nums[l]]--
	}

	return num
}

func CheckMap(m []int) bool {
	for _, v := range m {
		if v == 0 {
			return false
		}
	}
	return true
}

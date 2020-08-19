package program

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		target := 0 - nums[i]
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			for l < r {
				if nums[l]+nums[r] == target {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[l]+nums[r] < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}

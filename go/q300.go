package program

import "sort"

func lengthOfLIS2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	dp := make([]int, 1)
	dp[0] = nums[0]
	for i := range nums {
		if dp[len(dp)-1] < nums[i] {
			dp = append(dp, nums[i])
			continue
		}
		k := sort.Search(len(dp), func(j int) bool {
			return dp[j] >= nums[i]
		})
		dp[k] = nums[i]
	}
	return len(dp)
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))

	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(dp[i], result)
	}

	return result
}

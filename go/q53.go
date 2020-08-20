package program

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		result = max(result, dp[i])
	}

	return result
}

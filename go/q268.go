package program

func missingNumber(nums []int) int {
	var result int
	for i, v := range nums {
		result ^= i ^ v
	}
	return result ^ len(nums)
}

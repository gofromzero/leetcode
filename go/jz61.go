package program

func isStraight(nums []int) bool {
	arr := make([]int, 14)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]]++
	}
	for i := 1; i < 14; i++ {
		if arr[i] > 1 {
			return false
		}
	}
	minn := 1
	maxn := 13
	for minn < 14 && arr[minn] == 0 {
		minn++
	}
	for maxn >= 0 && arr[maxn] == 0 {
		maxn--
	}
	return maxn-minn <= 4
}

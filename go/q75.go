package program

func sortColors(nums []int) {
	a := 0
	b := len(nums) - 1
	for c := 0; c <= b; c++ {
		if nums[c] == 0 {
			nums[c], nums[a] = nums[a], nums[c]
			c++
			a++
		}
		if nums[c] == 2 {
			nums[c], nums[b] = nums[b], nums[c]
			c--
			b--
		}
	}
}

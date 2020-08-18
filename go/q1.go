package program

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		if val, ok := m[target-v]; ok {
			return []int{val, k}
		}
		m[v] = k
	}
	return nil
}

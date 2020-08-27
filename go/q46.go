package program

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	dfs(nums, []int{})

	return result
}

var result [][]int

func dfs(nums []int, temp []int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	if len(tmp) == len(nums) {
		result = append(result, tmp)
		return
	}

	for _, v := range nums {
		if !contain(tmp, v) {
			tmp = append(tmp, v)
			dfs(nums, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func contain(num []int, val int) bool {
	for _, v := range num {
		if v == val {
			return true
		}
	}
	return false
}

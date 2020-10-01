package program

import "math"

func findMode(root *TreeNode) []int {
	var arr []int
	getArr(root, &arr)
	var result []int
	var maxCount = math.MinInt64
	var count int
	var val int
	for _, v := range arr {
		if v != val {
			val = v
			count = 0
		}
		count++
		if maxCount < count {
			result = []int{val}
			maxCount = count
		} else if maxCount == count {
			result = append(result, val)
		}
	}
	return result
}

func getArr(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	getArr(root.Left, arr)
	*arr = append(*arr, root.Val)
	getArr(root.Right, arr)
}

package program

import (
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	var arr []int
	getSort(root, &arr)
	var minus = math.MaxInt64
	for i := 0; i < len(arr)-1; i++ {
		temp := arr[i+1] - arr[i]
		if minus > temp {
			minus = temp
		}
	}
	return minus
}

func getSort(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	getSort(root.Left, result)
	*result = append(*result, root.Val)
	getSort(root.Right, result)
	return
}

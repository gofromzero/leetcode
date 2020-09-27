package program

import "math"

func minDiffInBST(root *TreeNode) int {
	var list []int
	getList(root, &list)
	var minDiff = math.MaxInt64
	for i := 1; i < len(list); i++ {
		diff := list[i] - list[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func getList(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	getList(root.Left, list)
	*list = append(*list, root.Val)
	getList(root.Right, list)
}

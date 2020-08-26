package program

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return checkValidBST(root, math.MinInt64, math.MaxInt64)
}

func checkValidBST(root *TreeNode, l, r int) bool {
	if root == nil {
		return true
	}
	if root.Val <= l || root.Val >= r {
		return false
	}

	return checkValidBST(root.Left, l, root.Val) && checkValidBST(root.Right, root.Val, r)
}

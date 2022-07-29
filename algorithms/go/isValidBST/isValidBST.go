package isValidBST

import "math"

func isValidBST(root *TreeNode) bool {

	return isValidBST2(root, math.MinInt64, math.MaxInt64)
}

func isValidBST2(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValidBST2(root.Left, min, root.Val) && isValidBST2(root.Right, root.Val, max)
}

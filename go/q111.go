package program

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	m := math.MaxInt64
	if root.Left != nil {
		m = min(minDepth(root.Left), m)
	}
	if root.Right != nil {
		m = min(minDepth(root.Right), m)
	}
	return m + 1
}

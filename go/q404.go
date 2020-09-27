package program

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + right
	}

	return sumOfLeftLeaves(root.Left) + right
}

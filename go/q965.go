package program

func isUnivalTree(root *TreeNode) bool {
	return isUnivalTreeCheck(root, root.Val)
}

func isUnivalTreeCheck(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}

	return isUnivalTreeCheck(root.Left, val) && isUnivalTreeCheck(root.Right, val)
}

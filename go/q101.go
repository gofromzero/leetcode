package program

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root, root)
}

func checkSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && checkSymmetric(left.Left, right.Right) && checkSymmetric(left.Right, right.Left)
}

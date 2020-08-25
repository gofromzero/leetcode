package program

func isBalanced(root *TreeNode) bool {
	if depth(root) == -1 {
		return false
	}

	return true
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left == -1 || right == -1 || left-right < -1 || right-left < -1 {
		return -1
	}
	if left < right {
		return right + 1
	}

	return left + 1
}

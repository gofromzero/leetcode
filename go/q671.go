package program

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return help(root, root.Val)
}

func help(root *TreeNode, val int) int {
	if root == nil {
		return -1
	}
	if root.Val > val {
		return root.Val
	}
	left := help(root.Left, val)
	right := help(root.Right, val)
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}

	return min(left, right)
}

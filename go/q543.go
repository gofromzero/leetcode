package program

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	dfsDiameter(root, &ans)
	return ans
}

func dfsDiameter(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := dfsDiameter(root.Left, ans)
	right := dfsDiameter(root.Right, ans)
	*ans = max(*ans, left+right)
	return max(left, right) + 1
}

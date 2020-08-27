package program

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)

	return left + right + 1
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if left == right {
		return countNodes2(root.Right) + 1<<left
	}
	return countNodes2(root.Left) + 1<<right
}

func countLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		level++
		root = root.Left
	}
	return level
}

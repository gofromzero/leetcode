package program

func isCousins(root *TreeNode, x int, y int) bool {
	left, lDepth := findValue(root, root.Val, x, 0)
	right, rDepth := findValue(root, root.Val, y, 0)
	return left != right && lDepth == rDepth
}

func findValue(root *TreeNode, father, v int, depth int) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Val == v {
		return father, depth
	}
	left, lDepth := findValue(root.Left, root.Val, v, depth+1)
	right, rDepth := findValue(root.Right, root.Val, v, depth+1)

	return left | right, lDepth | rDepth
}

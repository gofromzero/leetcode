package program

func findTilt(root *TreeNode) int {
	_, titl := findTiltInternal(root)
	return titl
}

func findTiltInternal(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lSum, lTitl := findTiltInternal(root.Left)
	rSum, rTitl := findTiltInternal(root.Right)
	titl := lSum - rSum
	sum := lSum + rSum + root.Val
	if titl < 0 {
		return sum, lTitl + rTitl - titl
	}

	return sum, lTitl + rTitl + titl
}

package program

func sumEvenGrandparent(root *TreeNode) int {
	return sumEvenDfs(root, 0)
}

func sumEvenDfs(root *TreeNode, check int) int {
	if root == nil {
		return 0
	}
	check <<= 1
	if root.Val%2 == 0 {
		check |= 1
	}

	var result int
	if check&4 > 0 {
		result += root.Val
	}
	return result + sumEvenDfs(root.Left, check) + sumEvenDfs(root.Right, check)
}

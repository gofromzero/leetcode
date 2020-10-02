package program

func longestUnivaluePath(root *TreeNode) int {
	var maxNum int
	longestUnivaluePathInternal(root, &maxNum)
	return maxNum
}

func longestUnivaluePathInternal(root *TreeNode, val *int) int {
	if root == nil {
		return 0
	}
	left := longestUnivaluePathInternal(root.Left, val)
	right := longestUnivaluePathInternal(root.Right, val)
	var arrowLeft, arrowRight int
	if root.Left != nil && root.Val == root.Left.Val {
		arrowLeft += left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		arrowRight += right + 1
	}
	*val = max(*val, arrowLeft+arrowRight)
	return max(arrowLeft, arrowRight)
}

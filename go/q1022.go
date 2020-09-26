package program

func sumRootToLeaf(root *TreeNode) int {
	return sumRootToLeafInternal(root, 0, 0)
}

func sumRootToLeafInternal(root *TreeNode, val int, depth int) int {
	if root == nil {
		return val
	}
	val <<= 1
	val += root.Val
	var sum int
	if root.Left != nil {
		sum += sumRootToLeafInternal(root.Left, val, depth+1)
	}
	if root.Right != nil {
		sum += sumRootToLeafInternal(root.Right, val, depth+1)
	}
	if sum == 0 {
		return val
	}
	return sum
}

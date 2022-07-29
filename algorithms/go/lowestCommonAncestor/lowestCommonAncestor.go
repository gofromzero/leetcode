package lowestCommonAncestor

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parr := getNode(root, p)
	qarr := getNode(root, q)
	var result *TreeNode
	for i := 0; i < len(parr) && i < len(qarr) && parr[i].Val == qarr[i].Val; i++ {
		result = qarr[i]
	}
	return result
}

func getNode(root, val *TreeNode) []*TreeNode {
	var result = []*TreeNode{}
	for {
		result = append(result, root)
		if root.Val > val.Val {
			root = root.Left
		} else if val.Val > root.Val {
			root = root.Right
		} else {
			return result
		}
	}
}

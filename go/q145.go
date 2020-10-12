package program

func postorderTraversal(root *TreeNode) (res []int) {
	nodes := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(nodes) > 0 {
		for root != nil {
			nodes = append(nodes, root)
			root = root.Left
		}
		root = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			nodes = append(nodes, root)
			root = root.Right
		}
	}
	return res
}

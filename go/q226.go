package program

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(temp)
	return root
}

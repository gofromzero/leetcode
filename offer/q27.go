package offer

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(temp)
	return root
}

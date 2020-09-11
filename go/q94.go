package program

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inorderTraversal(root.Left)

	right := inorderTraversal(root.Right)
	result := append(left, root.Val)
	return append(result, right...)
}

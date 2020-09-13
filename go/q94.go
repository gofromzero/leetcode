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

func inorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	arr := make([]*TreeNode, 0)
	var head = root
	for head != nil || len(arr) != 0 {
		for head != nil {
			arr = append(arr, head)
			head = head.Left
		}
		head = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		result = append(result, head.Val)
		head = head.Right
	}

	return result
}

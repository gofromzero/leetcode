package program

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	for _, v := range root.Children {
		result = append(result, postorder(v)...)
	}

	result = append(result, root.Val)
	return result
}

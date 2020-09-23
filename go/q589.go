package program

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, root.Val)
	for _, v := range root.Children {
		result = append(result, preorder(v)...)
	}
	return result
}

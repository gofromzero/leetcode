package preorder

//func preorder(root *Node) []int {
//	if root == nil {
//		return nil
//	}
//	result := []int{root.Val}
//	for _, val := range root.Children {
//		result = append(result, preorder(val)...)
//	}
//	return result
//}

func preorder(root *Node) (ans []int) {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return
}

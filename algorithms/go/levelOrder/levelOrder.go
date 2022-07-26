package levelOrder

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var stack = []*TreeNode{root}
	l, r := 0, 1
	var result [][]int
	for len(stack) != l {
		var ans []int
		for ; l < r; l++ {
			ans = append(ans, stack[l].Val)
			if stack[l].Left != nil {
				stack = append(stack, stack[l].Left)
			}
			if stack[l].Right != nil {
				stack = append(stack, stack[l].Right)
			}
		}
		r = len(stack)
		result = append(result, ans)
	}

	return result
}

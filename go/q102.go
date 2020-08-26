package program

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	list := []*TreeNode{root}
	for len(list) > 0 {
		l := len(list)
		temp := make([]int, 0)
		for i := 0; i < l; i++ {
			head := list[0]
			list = list[1:]
			if head != nil {
				temp = append(temp, head.Val)
				list = append(list, head.Left, head.Right)
			}
		}
		if len(temp) > 0 {
			result = append(result, temp)
		}
	}
	return result
}

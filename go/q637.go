package program

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	nodes := []*TreeNode{root}
	result := make([]float64, 0)
	for len(nodes) != 0 {
		l := len(nodes)
		var sum int
		for i := 0; i < l; i++ {
			temp := nodes[i]
			sum += temp.Val
			if temp.Left != nil {
				nodes = append(nodes, temp.Left)
			}
			if temp.Right != nil {
				nodes = append(nodes, temp.Right)
			}
		}
		nodes = nodes[l:]
		result = append(result, float64(sum)/float64(l))
	}
	return result
}

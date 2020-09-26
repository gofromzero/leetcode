package program

func maxNDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var depth int
	for _, v := range root.Children {
		temp := maxNDepth(v)
		if depth < temp {
			depth = temp
		}
	}
	return depth + 1
}

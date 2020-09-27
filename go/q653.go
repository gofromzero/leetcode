package program

//func findTarget(root *TreeNode, k int) bool {
//	return findTargetM(root, k, map[int]bool{})
//}
//
//func findTargetM(root *TreeNode, k int, m map[int]bool)bool {
//	if root == nil {
//		return false
//	}
//	if _, ok := m[root.Val]; ok {
//		return true
//	}
//	m[k - root.Val]= true
//	return findTargetM(root.Left, k, m) || findTargetM(root.Right, k, m)
//}

// Better get value list
func findTarget(root *TreeNode, k int) bool {
	var list []int
	getTargetList(root, &list)
	m := make(map[int]bool)
	for _, v := range list {
		if _, ok := m[v]; ok {
			return true
		}
		m[k-v] = true
	}

	return false
}

func getTargetList(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	getTargetList(root.Left, arr)
	getTargetList(root.Right, arr)
}

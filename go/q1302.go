package program

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	arr := []*TreeNode{root}
	for len(arr) != 0 {
		l := len(arr)
		for i := 0; i < l; i++ {
			if arr[i].Left != nil {
				arr = append(arr, arr[i].Left)
			}
			if arr[i].Right != nil {
				arr = append(arr, arr[i].Right)
			}
		}
		if l == len(arr) {
			break
		}
		arr = arr[l:]
	}
	var sum int
	for _, v := range arr {
		sum += v.Val
	}
	return sum
}

var maxLevel, sum int

func deepestLeavesSum2(root *TreeNode) int {
	maxLevel, sum = 0, 0
	deepestDfs(root, 0)
	return sum
}

func deepestDfs(root *TreeNode, level int) {
	if root != nil {
		if level > maxLevel {
			maxLevel = level
			sum = root.Val
		} else if level == maxLevel {
			sum = sum + root.Val
		}
		deepestDfs(root.Left, level+1)
		deepestDfs(root.Right, level+1)
	}
}

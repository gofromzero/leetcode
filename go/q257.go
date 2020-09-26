package program

import "strconv"

//func binaryTreePaths(root *TreeNode) []string {
//	return binaryTreePathsInternal(root, "")
//}
//
//func binaryTreePathsInternal(root *TreeNode, str string) []string {
//	if root == nil {
//		return []string{}
//	}
//	if str == "" {
//		str = strconv.Itoa(root.Val)
//	} else {
//		str += "->"+ strconv.Itoa(root.Val)
//	}
//
//	var result []string
//	result = append(result, binaryTreePathsInternal(root.Left, str)...)
//	result = append(result, binaryTreePathsInternal(root.Right, str)...)
//
//	if len(result) == 0 {
//		result = append(result, str)
//	}
//	return result
//}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	path := make([]string, 0)
	binaryTreePathsdfs(root, "", &path)
	return path
}
func binaryTreePathsdfs(root *TreeNode, str string, path *[]string) {
	if root == nil {
		return
	}
	str += strconv.Itoa(root.Val) + "->"
	if root.Left == nil && root.Right == nil {
		// 判断 并删除->
		*path = append(*path, str[:len(str)-2])
	}

	binaryTreePathsdfs(root.Left, str, path)
	binaryTreePathsdfs(root.Right, str, path)
}

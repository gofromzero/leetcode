package program

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	var index int
	for i := 0; i < l; i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}

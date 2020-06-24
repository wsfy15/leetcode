/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	var res int
	tilt(root, &res)
	return res
}

func tilt(root *TreeNode, total *int) int {
	if root == nil {
		return 0
	}

	left, right := tilt(root.Left, total), tilt(root.Right, total)
	*total += abs(left - right)

	return left + right + root.Val
}

func abs(num int) int {
	y := num >> 31
	return num ^ y - y
}
// @lc code=end


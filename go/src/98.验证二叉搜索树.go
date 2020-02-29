/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
 import "math"
func isValidBST(root *TreeNode) bool {
	return util(root, math.MaxInt64, math.MinInt64)
}

func util(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return util(root.Left, root.Val, min) && util(root.Right, max, root.Val)
}
// @lc code=end


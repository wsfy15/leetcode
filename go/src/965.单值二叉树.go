/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Val) && dfs(root.Right, root.Val)
}

func dfs(root *TreeNode, val int) bool {
	if root == nil {
		return true
	} 
	if root.Val == val {
		return dfs(root.Left, val) && dfs(root.Right, val)
	} else {
		return false
	}
}
// @lc code=end


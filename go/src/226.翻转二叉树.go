/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
	util(root)
	return root
}

func util(root *TreeNode) {
	if root == nil {
		return 
	}

	util(root.Left)
	util(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
// @lc code=end


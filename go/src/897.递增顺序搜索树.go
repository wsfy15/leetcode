/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序搜索树
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
func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	inorder(root, dummy)
	return dummy.Right
}

func inorder(root, prev *TreeNode) *TreeNode {
	if root == nil {
		return prev
	}

	tmp := inorder(root.Left, prev)
	tmp.Right = root
	root.Left = nil
	return inorder(root.Right, root)
}

// @lc code=end


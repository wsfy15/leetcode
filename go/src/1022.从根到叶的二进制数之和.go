/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
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
 var res int
func sumRootToLeaf(root *TreeNode) int {
	res = 0
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, cur int) {
	if root == nil {
		return
	}

	cur = (cur << 1) + root.Val
	if root.Left == nil && root.Right == nil {
		res = (res + cur) % 1_000_000_007
	} else {
		if root.Left != nil {
			dfs(root.Left, cur)
		}
		if root.Right != nil {
			dfs(root.Right, cur)
		}
	}
}
// @lc code=end


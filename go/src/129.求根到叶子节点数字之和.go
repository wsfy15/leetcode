/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func sumNumbers(root *TreeNode) int {
	res := 0
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, cur int, res *int) {
	if root == nil {
		return
	}

	cur = cur * 10 + root.Val

	if root.Left == nil && root.Right == nil {
		*res += cur
		return
	}

	dfs(root.Left, cur, res)
	dfs(root.Right, cur, res)
}
// @lc code=end


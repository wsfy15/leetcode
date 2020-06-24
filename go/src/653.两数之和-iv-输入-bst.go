/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
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
func findTarget(root *TreeNode, k int) bool {
	counter := map[int]bool{}
	return dfs(root, k, counter)
}

func dfs(root *TreeNode, k int, counter map[int]bool) bool {
	if root == nil {
		return false
	}

	if _, ok := counter[k - root.Val]; ok {
		return true
	}
	counter[root.Val] = true

	return dfs(root.Left, k, counter) || dfs(root.Right, k, counter)
}
// @lc code=end


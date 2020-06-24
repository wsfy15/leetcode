/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}
	dfs(root, "", &res)
	return res
}

func dfs(root *TreeNode, cur string, res *[]string) {
	cur = cur + strconv.Itoa(root.Val) + "->"
	if root.Left == nil && root.Right == nil {
		*res = append(*res, cur[:len(cur) - 2])
	}

	if root.Left != nil {
		dfs(root.Left, cur, res)
	}
	if root.Right != nil {
		dfs(root.Right, cur, res)
	}
}

// @lc code=end


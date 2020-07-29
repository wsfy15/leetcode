/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
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
func rangeSumBST(root *TreeNode, L int, R int) int {
	res := 0
	dfs(root, L, R, &res)

	return res
}

func dfs(root *TreeNode, L int, R int, res *int) {
	if root == nil {
		return
	}

	if L <= root.Val && root.Val <= R {
		*res += root.Val
		dfs(root.Left, L, R, res)
		dfs(root.Right, L, R, res)
	} else if root.Val < L {
		dfs(root.Right, L, R, res)
	} else if root.Val > R {
		dfs(root.Left, L, R, res)
	}
}
// @lc code=end


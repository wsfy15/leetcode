/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	var diff, prev int = math.MaxInt32, -1
	inOrder(root, &diff, &prev)
	return diff
}

func inOrder(root *TreeNode, diff, prev *int) {
	if root == nil {
		return
	}

	inOrder(root.Left, diff, prev)
	if *prev != -1 && root.Val - *prev < *diff {
		*diff = root.Val - *prev
	}
	*prev = root.Val
	inOrder(root.Right, diff, prev)
}
// @lc code=end


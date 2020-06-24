/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
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
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	dfs(t1, t2)
	return t1
}

func dfs(t1 *TreeNode, t2 *TreeNode) {
	t1.Val += t2.Val
	if t1.Left != nil && t2.Left != nil {
		dfs(t1.Left, t2.Left)
	} else if t1.Left == nil {
		t1.Left = t2.Left
	}

	if t1.Right != nil && t2.Right != nil {
		dfs(t1.Right, t2.Right)
	} else if t1.Right == nil {
		t1.Right = t2.Right
	}
}
// @lc code=end


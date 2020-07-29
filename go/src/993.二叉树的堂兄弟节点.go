/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
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
func isCousins(root *TreeNode, x int, y int) bool {
	xdp := dfs(root, x, root.Val, 0) // xdp[0]: depth xdp[1]: parent
	ydp := dfs(root, y, root.Val, 0)
	return xdp[0] == ydp[0] && xdp[1] != ydp[1]
}

func dfs(root *TreeNode, val, parent, depth int) []int {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return []int{depth + 1, parent}
	}

	res := dfs(root.Left, val, root.Val, depth + 1)
	if res != nil {
		return res
	}
	return dfs(root.Right, val, root.Val, depth + 1)
}
// @lc code=end


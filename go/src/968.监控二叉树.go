/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [968] 监控二叉树
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

func minCameraCover(root *TreeNode) int {
	res = 0
	if root == nil {
		return res
	}
	if dfs(root) == 0 {
		res++
	}
	return res
}

// 返回该节点的状态：
// 0: 未安装监视器，不可被子女观测
// 1: 已安装监视器
// 2: 未安装，子女可观测
func dfs(root *TreeNode) int {
	l, r := 2, 2
	if root.Left != nil {
		l = dfs(root.Left)
	}
	if root.Right != nil {
		r = dfs(root.Right)
	}

	if l == 0 || r == 0 {
		res++
		return 1
	} else if l == 1 || r == 1 {
		return 2
	}
	return 0
}

// @lc code=end


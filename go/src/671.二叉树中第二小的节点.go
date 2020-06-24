/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}
	res := getBiggerValue(root, root.Val)
	if res == root.Val {
		return -1
	}
	return res
}

func getBiggerValue(root *TreeNode, val int) int {
	if root == nil {
		return val
	}

	if root.Val > val || root.Left == nil {
		return root.Val
	}

	left := getBiggerValue(root.Left, root.Val)
	right := getBiggerValue(root.Right, root.Val)
	if left == root.Val && right == root.Val {
		return root.Val
	}
	if left == root.Val || right == root.Val {
		return max(left, right)
	}
	return min(left, right)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end


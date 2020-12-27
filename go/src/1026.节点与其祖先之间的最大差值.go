/*
 * @lc app=leetcode.cn id=1026 lang=golang
 *
 * [1026] 节点与其祖先之间的最大差值
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
func maxAncestorDiff(root *TreeNode) int {
	return inorder(root, math.MinInt32, math.MaxInt32)
}

func inorder(root *TreeNode, maxVal, minVal int) int {
	if root == nil {
		return 0
	}

	maxVal = max(maxVal, root.Val)
	minVal = min(minVal, root.Val)

	return max(maxVal-root.Val, root.Val-minVal, inorder(root.Left, maxVal, minVal), inorder(root.Right, maxVal, minVal))
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end


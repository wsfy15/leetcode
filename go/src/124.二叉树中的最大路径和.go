/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	get(root, &res)
	return res
}

//  [-10,9,20,null,null,15,7]
// 返回最长路径的最长部分，左子树或右子树
func get(root *TreeNode, res *int) int {
	if root == nil {
		return math.MinInt32
	}
	left := get(root.Left, res)
	right := get(root.Right, res)
	*res = max(*res, left + right + root.Val, left + root.Val, right + root.Val, root.Val)
	return max(left + root.Val, right + root.Val, root.Val)
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
// @lc code=end


/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
//  中序遍历
func findMode(root *TreeNode) []int {
	res := []int{}
	var pre *TreeNode
	var max, count int
	inOrder(root, &count, &max, &pre, &res)
	return res
}

func inOrder(root *TreeNode, count, max *int, pre **TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, count, max, pre, res)
	if *pre != nil && (*pre).Val == root.Val {
		(*count)++
	} else {
		*count = 1
	}

	if *count > *max {
		*max = *count
		*res = []int{root.Val}
	} else if *count == *max {
		*res = append(*res, root.Val)
	}

	*pre = root
	inOrder(root.Right, count, max, pre, res)
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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
//  路径有可能是遍布在左右子树上的，而不是单方向一直下去
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	getPathLen(root, &res)
	return res
}

func getPathLen(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	llen, rlen := 0, 0
	var l, r int
	l = getPathLen(root.Left, res)
	r = getPathLen(root.Right, res)
	if root.Left != nil && root.Left.Val == root.Val {
		llen = l + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		rlen = r + 1
	}

	*res = max(*res, llen + rlen)
	return max(llen, rlen)
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


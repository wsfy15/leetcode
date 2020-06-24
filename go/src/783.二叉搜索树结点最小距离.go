/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树结点最小距离
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
func minDiffInBST(root *TreeNode) int {
	var lastNum, res int = math.MaxInt32, math.MaxInt32
	validLastNum := false
	inOrder(root, &lastNum, &res, &validLastNum)

	return res
}

func inOrder(root *TreeNode, lastNum, res *int, validLastNum *bool) {
	if root == nil {
		return
	}

	inOrder(root.Left, lastNum, res, validLastNum)
	if *validLastNum {
		*res = min(*res, root.Val - *lastNum)
	}

	*validLastNum = true
	*lastNum = root.Val
	inOrder(root.Right, lastNum, res, validLastNum)
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


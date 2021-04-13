/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
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
	if root == nil {
		return 0
	}

	var prev *TreeNode
	var res int = math.MaxInt32
	inorder(root, &prev, &res)
	return res
}

func inorder(root *TreeNode, prev **TreeNode, res *int) {
	if root == nil {
		return
	}

	inorder(root.Left, prev, res)
	if *prev != nil {
		*res = min(*res, root.Val-(*prev).Val)
	}
	*prev = root
	inorder(root.Right, prev, res)
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


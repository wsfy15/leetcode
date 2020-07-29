/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	res, _ := inorder(root, k)
	return res
}

func inorder(root *TreeNode, k int) (int, int) {
	if root == nil {
		return 0, k
	}

	res, k := inorder(root.Left, k)
	if k == 0 {
		return res, 0
	} else if k == 1 {
		return root.Val, 0
	}
	return inorder(root.Right, k - 1)
}
// @lc code=end


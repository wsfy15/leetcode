/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序查找树
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
 var res *TreeNode
func increasingBST(root *TreeNode) *TreeNode {
	res = nil
	var last *TreeNode = nil
	inorder(root, &last)

	return res
}

func inorder(root *TreeNode, last **TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left, last)

	if *last == nil {
		res = root
	} else {
		(*last).Right = root
	}

	*last = root
	root.Left = nil
	
	inorder(root.Right, last)
}
// @lc code=end


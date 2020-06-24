/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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

var pre, node1, node2 *TreeNode
func recoverTree(root *TreeNode)  {
	pre, node1, node2 = nil, nil, nil
	inorder(root)
	node1.Val, node2.Val = node2.Val, node1.Val
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	if pre != nil {
		if pre.Val > root.Val {
			if node1 == nil {
				node1 = pre
			}
			node2 = root
		}
	}
	pre = root
	inorder(root.Right)
}
// @lc code=end


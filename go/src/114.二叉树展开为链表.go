/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode)  {
	for root != nil {
		if root.Left != nil {
			tmp := root.Left
			// 把当前节点的右子树接到左子树最右节点
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			tmp.Right = root.Right
			// 然后把左子树移到右子树
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}
// @lc code=end


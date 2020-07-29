/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 得到左右子树高度
	// 如果左子树更高，说明右子树已经是高度为right的满二叉树了
	// 如果左右子树高度相同，则说明左子树是高度为left的满二叉树，而右子树不一定是满二叉树
	left, right := getHeight(root.Left), getHeight(root.Right)
	if left == right {
		return 1 + (1 << left) - 1 + countNodes(root.Right)
	} else {
		return 1 + (1 << right) - 1 + countNodes(root.Left)
	}
}

func getHeight(root *TreeNode) int {
	count := 0
	for root != nil {
		count++
		root = root.Left
	}
	return count
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
 var max int 
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = 0
	dfs(root)

	return max
}


func dfs(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	left := dfs(cur.Left) // 左子树的高度
	right := dfs(cur.Right) // 右子树高度
	if left + right > max { // 两个高度相加就是最大边数
		max = left + right
	}
	if left > right { // +1 加上从当前节点到父节点这条边
		return left + 1
	} 
	return right + 1
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	queue := []*TreeNode{root}
	backupQueue := []*TreeNode{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Left != nil {
			backupQueue = append(backupQueue, cur.Left)
		}
		if cur.Right != nil {
			backupQueue = append(backupQueue, cur.Right)
		}

		if len(queue) == 0 {
			res = append(res, cur.Val)
			queue = backupQueue
			backupQueue = []*TreeNode{}
		}
	}
	return res
}
// @lc code=end


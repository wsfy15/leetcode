/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	res := []float64{}
	queue, backupQueue := []*TreeNode{root}, []*TreeNode{}
	sum, count := 0, 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		sum += cur.Val

		if cur.Left != nil {
			backupQueue = append(backupQueue, cur.Left)
		}
		if cur.Right != nil {
			backupQueue = append(backupQueue, cur.Right)
		}

		if len(queue) == 0 {
			res = append(res, float64(sum) / float64(count))
			queue, backupQueue = backupQueue, []*TreeNode{}
			sum = 0
			count = 0
		}
	}

	return res
}
// @lc code=end


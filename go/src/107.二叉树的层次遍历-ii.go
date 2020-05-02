/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue, backupQueue := []*TreeNode{root}, []*TreeNode{}

	var level []int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		level = append(level, cur.Val)

		if cur.Left != nil {
			backupQueue = append(backupQueue, cur.Left)
		}
		if cur.Right != nil {
			backupQueue = append(backupQueue, cur.Right)
		}

		if len(queue) == 0 {
			queue, backupQueue = backupQueue, []*TreeNode{}
			res = append(res, level)
			level = []int{}
		}
	}

	i, j := 0, len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
// @lc code=end


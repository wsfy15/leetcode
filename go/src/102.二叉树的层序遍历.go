/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{[]int{}}
	index := 0
	queue := []*TreeNode{root}
	backupQueue := []*TreeNode{} 

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res[index] = append(res[index], cur.Val)
		if cur.Left != nil {
			backupQueue = append(backupQueue, cur.Left)
		}
		if cur.Right != nil {
			backupQueue = append(backupQueue, cur.Right)
		}

		if len(queue) == 0 {
			queue, backupQueue = backupQueue, []*TreeNode{}
			index++
			if len(queue) > 0 {
				res = append(res, []int{})
			}
		}
	}

	return res
}
// @lc code=end


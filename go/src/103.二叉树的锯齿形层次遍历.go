/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	flag := 0 // 奇数则逆序输出
	for len(queue) > 0 {
		n := len(queue)
		cur := make([]int, n)
		for i := 0; i < n; i++ {
			if flag & 1 == 0 {
				cur[i] = queue[i].Val
			} else {
				cur[n - i - 1] = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		res = append(res, cur)
		flag++
		queue = queue[n:]
	}
	return res
}
// @lc code=end


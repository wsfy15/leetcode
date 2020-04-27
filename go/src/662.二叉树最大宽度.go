/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
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

type node struct {
	pos int
	level int
	t *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max, curLevel, left := 0, 0, 0
	queue := []*node{&node{0, 0, root}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.level != curLevel {
			curLevel++
			left = cur.pos // 新的一层最左节点的下标
		}

		if cur.pos - left > max {
			max = cur.pos - left 
		}
		if cur.t.Left != nil {
			queue = append(queue, &node{cur.pos * 2, curLevel + 1, cur.t.Left})
		}
		if cur.t.Right != nil {
			queue = append(queue, &node{cur.pos * 2 + 1, curLevel + 1, cur.t.Right})
		}
	}

	return max + 1
}
// @lc code=end


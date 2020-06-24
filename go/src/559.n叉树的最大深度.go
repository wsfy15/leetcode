/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
  if root == nil {
		return 0
	}

	max := 0
	for _, child := range root.Children {
		depth := maxDepth(child)
		if depth > max {
			max = depth
		}
	}
	return max + 1
}
// @lc code=end


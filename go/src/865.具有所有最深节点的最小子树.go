/*
 * @lc app=leetcode.cn id=865 lang=golang
 *
 * [865] 具有所有最深节点的最小子树
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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left, right := dfs(root.Left), dfs(root.Right)
	if left == right {
		return root
	}

	if left > right {
		return subtreeWithAllDeepest(root.Left)
	}
	return subtreeWithAllDeepest(root.Right)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfs(root.Left), dfs(root.Right)) + 1
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end


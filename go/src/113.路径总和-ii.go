/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	dfs(root, sum, []int{}, &res)
	return res
}

func dfs(root *TreeNode, sum int, cur []int, res *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		sum -= root.Val
		if sum == 0 {
			dst := make([]int, len(cur))
			copy(dst, cur)
			dst = append(dst, root.Val)
			*res = append(*res, dst)
		}
		return
	}

	dfs(root.Left, sum - root.Val, append(cur, root.Val), res)
	dfs(root.Right, sum - root.Val, append(cur, root.Val), res)
}
// @lc code=end


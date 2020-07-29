/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
//  类似数组前n项和的思路，sum[4] == sum[1]的话，说明从 1 到 4 的路径和为0
func pathSum(root *TreeNode, sum int) int {
	nums := map[int]int{}
	nums[0] = 1
	return dfs(root, sum, 0, nums)
}

func dfs(root *TreeNode, sum, cur int, nums map[int]int) int {
	if root == nil {
		return 0
	}

	cur += root.Val
	res := nums[cur - sum]
	nums[cur]++
	res += dfs(root.Left, sum, cur, nums) + dfs(root.Right, sum, cur, nums)
	nums[cur]--
	return res
}
// @lc code=end


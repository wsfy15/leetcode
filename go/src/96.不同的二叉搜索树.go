/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	dp := make([]int, n + 1) // i个节点有多少种情况
	dp[0] = 1
	dp[1] = 1

	// 考虑1...i为节点共有多少种情况
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ { // 根节点的值从1到i
			dp[i] += dp[j - 1] * dp[i - j] // 左子树有j - 1个节点 右子树有i - j个节点
		}
	}

	return dp[n]
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	if n * m == 0 {
		return 0
	}
	
	// dp[i][j] 表示 以A[i-1]与B[j-1]结尾的公共字串的长度
	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m + 1)
	}

	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if A[i - 1] == B[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
				res = max(res, dp[i][j])
			}
		}
	}

	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end


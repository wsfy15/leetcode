/*
 * @lc app=leetcode.cn id=1269 lang=golang
 *
 * [1269] 停在原地的方案数
 */

// @lc code=start
func numWays(steps int, arrLen int) int {
	var mod int = 1e9 + 7
	dp := make([]int, arrLen)
	dp[0] = 1

	for i := 1; i <= steps; i++ {
		ndp := make([]int, arrLen)
		for j := 0; j <= min(i, arrLen-1); j++ {
			ndp[j] = dp[j]
			if j > 0 {
				ndp[j] += dp[j-1]
			}
			if j < arrLen-1 {
				ndp[j] += dp[j+1]
			}

			ndp[j] %= mod
		}
		dp = ndp
	}

	return dp[0]
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end


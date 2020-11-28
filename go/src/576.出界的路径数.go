/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */

// @lc code=start
func findPaths(m int, n int, N int, i int, j int) int {
	var mod int = 1e9 + 7

	// dp[i][j][k]: 从外界出发第k步走到(i,j)的路径总数
	dp := make([][][]int, m+2)
	for i := 0; i < m+2; i++ {
		dp[i] = make([][]int, n+2)
		for j := 0; j < n+2; j++ {
			dp[i][j] = make([]int, N+1)
		}
	}

	// 初始化四周的边界为1
	for i := 0; i < m+2; i++ {
		dp[i][0][0] = 1
		dp[i][n+1][0] = 1
	}
	for i := 1; i < n+2; i++ {
		dp[0][i][0] = 1
		dp[m+1][i][0] = 1
	}

	for k := 1; k <= N; k++ {
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				dp[i][j][k] = (dp[i-1][j][k-1] + dp[i+1][j][k-1] + dp[i][j-1][k-1] + dp[i][j+1][k-1]) % mod
			}
		}
	}

	sum := 0
	for k := 1; k <= N; k++ {
		sum = (sum + dp[i+1][j+1][k]) % mod
	}

	return sum
}

// @lc code=end


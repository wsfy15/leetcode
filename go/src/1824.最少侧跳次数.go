/*
 * @lc app=leetcode.cn id=1824 lang=golang
 *
 * [1824] 最少侧跳次数
 */

// @lc code=start
func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
		dp[i][0], dp[i][1], dp[i][2] = n+1, n+1, n+1
	}
	dp[0][0], dp[0][1], dp[0][2] = 1, 0, 1

	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			if obstacles[i] == j+1 {
				continue
			}

			dp[i][j] = dp[i-1][j]
		}

		for j := 0; j < 3; j++ {
			if obstacles[i] == j+1 {
				continue
			}

			for k := 0; k < 3; k++ {
				if k == j || k+1 == obstacles[i] {
					continue
				}

				dp[i][j] = min(dp[i][j], dp[i][k]+1)
			}
		}
	}

	return min(dp[n-1]...)
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


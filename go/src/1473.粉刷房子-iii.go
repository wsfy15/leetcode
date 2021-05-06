/*
 * @lc app=leetcode.cn id=1473 lang=golang
 *
 * [1473] 粉刷房子 III
 */

// @lc code=start
func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// dp[i][j][k] 当第i个房子涂为第j种颜色，且形成k个街区时的cost
	// i: [0, m - 1]  j:[0, n]
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, target+1)
			for k := 0; k <= target; k++ {
				dp[i][j][k] = math.MaxInt32
			}
		}
	}

	if houses[0] != 0 {
		dp[0][houses[0]][1] = 0
	} else {
		for i := 1; i <= n; i++ {
			dp[0][i][1] = cost[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		if houses[i] == 0 {
			for j := 1; j <= n; j++ { // 涂上每一种颜色
				for k := 1; k <= target; k++ {
					for j2 := 1; j2 <= n; j2++ {
						if j == j2 {
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][j][k]+cost[i][j-1])
						} else {
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][j2][k-1]+cost[i][j-1])
						}
					}
				}
			}
		} else {
			for k := 1; k <= target; k++ {
				for j := 1; j <= n; j++ {
					if houses[i] == j {
						dp[i][houses[i]][k] = min(dp[i][j][k], dp[i-1][j][k])
					} else {
						dp[i][houses[i]][k] = min(dp[i][houses[i]][k], dp[i-1][j][k-1])
					}
				}
			}
		}
	}

	res := math.MaxInt32
	for i := 1; i <= n; i++ {
		res = min(res, dp[m-1][i][target])
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
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


/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}

	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n - 1 && j == m - 1 {
				dp[i][j] = max(1, 1 - dungeon[i][j])
			} else if i == n - 1 {
				dp[i][j] = max(1, dp[i][j+1] - dungeon[i][j])
			} else if j == m - 1 {
				dp[i][j] = max(1, dp[i+1][j] - dungeon[i][j])
			} else {
				dp[i][j] = max(1, min(dp[i][j+1], dp[i+1][j]) - dungeon[i][j])
			}
		}
	}

	return dp[0][0]
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

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end


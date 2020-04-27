/*
 * @lc app=leetcode.cn id=629 lang=golang
 *
 * [629] K个逆序对数组
 */

// @lc code=start
// 用dp[i - 1][j]表示用数字[1, 2, ……, i - 1]组成j个逆序对的方案数
// 考虑第i个数时，插入到[1,……,i-1]的每个位置，会增加[i - 1, i - 2, ……, 0]种方案
// dp[i][j] = dp[i - 1][j] + dp[i - 1][j - 1] + …… + dp[i - 1][j - i + 1]
// dp[i][j-1] = dp[i-1][j-1] + dp[i-1][j-2] + …… + dp[i-1][j-i]
// dp[i][j] - dp[i][j-1] = dp[i-1][j] - dp[i-1][j-i]
// dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-i]

// dp[i][0] = 1 for all i
// dp[i][j] = 0 for j < 0
func kInversePairs(n int, k int) int {
	if k <= 0 {
		return 1
	}

	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k + 1)
		dp[i][0] = 1
	}

	var mod int = 1e9 + 7
	// dp[1][0] = 1 // 1个数只能构成0个逆序对
	for i := 2; i <= n; i++ {
		max := i * (i - 1) / 2 // i个数最多构成逆序对数量最大值 逆序时
		for j := 1; j <= k && j <= max; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			if j - i >= 0 {
				dp[i][j] -= dp[i-1][j-i]
			}
			dp[i][j] += mod // 先加再取模，防止减后的值是负数
			dp[i][j] %= mod
		}
	}

	return dp[n][k]
}
// @lc code=end


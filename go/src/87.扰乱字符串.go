/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	n := len(s1)
	dp := make([][][]bool, n) // dp[i][j][k]：s1[i:i+k]和s2[j:j+k] 是否互为扰乱字符串
	// dp[i][j][k] = (dp[i][j][w] && dp[i+w][j+w][k-w] for 1 <= w < k) || (dp[i][j+k-w][w] && dp[i+w][j][k-w] for 1 <= w < k)
	
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n + 1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	// 从长度为2的子串开始
	for k := 2; k <= n; k++ {
		for i := 0; i <= n - k; i++ {
			for j := 0; j <= n - k; j++ {
				for w := 1; w < k; w++ {
					if (dp[i][j][w] && dp[i+w][j+w][k-w]) || dp[i][j+k-w][w] && dp[i+w][j][k-w] {
						dp[i][j][k] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][n]
}
// @lc code=end


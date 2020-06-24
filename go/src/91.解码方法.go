/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	} 

	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			if s[i - 1] > '2' || s[i - 1] == '0' {
				return 0
			} else {
				dp[i + 1] = dp[i - 1]
			}
		} else if s[i - 1] == '0' {
			dp[i + 1] = dp[i]
		}	else {
			if v, _ := strconv.Atoi(s[i-1:i+1]); v < 27 {
				dp[i + 1] = dp[i] + dp[i - 1]
			} else {
				dp[i + 1] = dp[i]
			}
		}
	}
	return dp[n]
}
// @lc code=end


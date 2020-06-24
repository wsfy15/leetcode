/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// s能不能由wordDict中的一个或多个单词(不用全部)连续组合而成
	m := map[string]bool{}
	for _, word := range wordDict {
		m[word] = true
	}

	n := len(s)
	dp := make([]bool, n + 1) // dp[i]表示s[:i]可以由wordDict中的单词组合而成
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				if _, ok := m[s[j: i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n]
}
// @lc code=end


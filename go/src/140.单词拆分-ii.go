/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	dp, wordSet := wordBreak2(s, wordDict)
	if !dp[len(s)] {
		return nil
	}

	maxLength := 0
	for k, _ := range wordSet {
		maxLength = max(len(k), maxLength)
	}

	res, cur := []string{}, []string{}
	dfs(s, 0, maxLength, dp, wordSet, cur, &res)

	return res
}

func dfs(s string, index, maxLength int, dp []bool, wordSet map[string]bool, cur []string, res *[]string) {
	if index == len(s) {
		*res = append(*res, strings.Join(cur, " "))
		return
	}

	for i := index; i < min(len(s), index+maxLength); i++ {
		if dp[i+1] && wordSet[s[index:i+1]] {
			cur = append(cur, s[index:i+1])
			dfs(s, i+1, maxLength, dp, wordSet, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}

func wordBreak2(s string, wordDict []string) ([]bool, map[string]bool) {
	// s能不能由wordDict中的一个或多个单词(不用全部)连续组合而成
	m := map[string]bool{}
	for _, word := range wordDict {
		m[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1) // dp[i]表示s[:i]可以由wordDict中的单词组合而成
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] {
				if _, ok := m[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp, m
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
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


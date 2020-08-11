/*
 * @lc app=leetcode.cn id=1332 lang=golang
 *
 * [1332] 删除回文子序列
 */

// @lc code=start
func removePalindromeSub(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	i, j := 0, n-1
	for i < j {
		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}
	return 1
}

// @lc code=end


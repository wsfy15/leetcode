/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n & 1 > 0 {
		return false
	}

	// 周期为i
	for i := 1; i <= n >> 1; i++ {
		if n % i > 0 {
			continue
		}

		j := i
		for j < n && s[j] == s[j % i] {
			j++
		}
		if j == n {
			return true
		}
	}
	return false
}
// @lc code=end


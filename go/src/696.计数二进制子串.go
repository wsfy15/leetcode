/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	res := 0
	last, cur := 0, 1
	for i := 1; i < n; i++ {
		num := s[i] - '0'
		if num == s[i - 1] - '0' {
			cur++
		} else {
			last, cur = cur, 1
		}

		if cur <= last {
			res++
		}
	}

	return res
}
// @lc code=end


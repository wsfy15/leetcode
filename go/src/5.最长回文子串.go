/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	n, maxLen := len(s), 0
	if n <= 1 {
		return s
	}
	// 以每个位置为中心点考虑最长回文串
	left, right := 0, 0

	for i := 0; i < n; i++ {
		start, end := getLongestPalindrome(s, i, i)
		if end - start > maxLen {
			maxLen = end - start
			left, right = start, end
		}

		start, end = getLongestPalindrome(s, i, i + 1)
		if end - start > maxLen {
			maxLen = end - start
			left, right = start, end
		}
	}

	return s[left:right+1]
}

// 返回起始和结束下标
func getLongestPalindrome(s string, start, end int) (int, int) {
	// 对于初始start == end的情况，肯定会执行这个循环至少一次，
	//   不满足时start和end要么在边界外，或者s[start] != s[end]
	// 对于初始start == end + 1的情况，如果一开始就不相等，直接返回也无所谓
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}

	return start + 1, end - 1
}
// @lc code=end


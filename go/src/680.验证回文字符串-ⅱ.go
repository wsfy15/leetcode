/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
func validPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return isValid(s, i + 1, j) || isValid(s, i, j - 1)
		}
	}

	return true
}

func isValid(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
// @lc code=end


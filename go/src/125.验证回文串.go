/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	s = strings.ToLower(s)
	for i < j {
		for !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') ) && i < j {
			i++
		}
		for !((s[j] >= 'a' && s[j] <= 'z')|| (s[j] >= '0' && s[j] <= '9') ) && i < j {
			j--
		}
	
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
// @lc code=end


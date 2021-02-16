/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	counter := make([]int, 26)
	for _, ch := range s1 {
		counter[ch - 'a']++
	}

	length, start := len(s1), 0
	for index, ch := range s2 {
		if counter[ch - 'a'] > 0 {
			counter[ch - 'a']--
			if length == index - start + 1 {
				return true
			}
		} else {
			for start < index && s2[start] != s2[index] {
				counter[s2[start] - 'a']++
				start++
			}
			start++
		}
	}

	return false
}
// @lc code=end


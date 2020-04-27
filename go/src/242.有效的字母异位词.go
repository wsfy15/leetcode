/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if s == t {
		return true
	}
	size := len(s)
	if size != len(t) {
		return false
	}

	counter := make([]int, 26)
	for i := 0; i < size; i++ {
		counter[s[i] - 'a']++
	}

	for i := 0; i < size; i++ {
		counter[t[i] - 'a']--
	}

	for i := 0; i < 26; i++ {
		if counter[i] != 0 {
			return false
		}
	}
	return true
}
// @lc code=end


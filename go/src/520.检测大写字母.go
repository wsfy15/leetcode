/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	n := len(word)
	if n < 2 {
		return true
	}
	upper, lower := 0, 0
	for i := 0; i < n; i++ {
		if 'A' <= word[i] && word[i] <= 'Z' {
			upper++
		} else {
			lower++
		}
	}

	if upper == 1 && 'A' <= word[0] && word[0] <= 'Z' {
		return true
	}
	return lower == 0 || upper == 0
}
// @lc code=end


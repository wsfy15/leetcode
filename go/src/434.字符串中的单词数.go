/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	n := len(s)
	count := 0
	i := 0
	for i < n && s[i] == ' ' {
		i++
	}
	for i < n {
		count++
		for i < n && s[i] != ' ' {
			i++
		}
		for i < n && s[i] == ' ' {
			i++
		}
	}

	return count
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		if counter[s[i] - 'a'] == 0 {
			// 第一次出现
			counter[s[i] - 'a'] = i + 1
		} else {
			// 出现多次
			counter[s[i] - 'a'] = -1
		}
	}

	index := len(s) + 1
	for i := 0; i < 26; i++ {
		if counter[i] > 0 && counter[i] <= index {
			index = counter[i]
		}
	}
	
	if index == len(s) + 1 {
		return -1
	}
	return index - 1
}
// @lc code=end


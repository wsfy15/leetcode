/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {
	size := len(s)
	if size < 2 {
		return size
	}

	count := make(map[byte]int) // 记录每个字符出现次数
	for i := range s {
		count[s[i]]++
	}

	res := 0
	single := false
	for _, v := range count {
		if v & 1 == 0 {
			res += v
		} else {
			res += v - 1
			single = true
		}
	}

	if single {
		return res + 1
	}
	return res
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	n := len(s) - 1
	for n >= 0 && s[n] == ' ' {
		n--
	}

	i := n
	for i >= 0 && s[i] != ' ' {
		i--
	}

	return n - i
}
// @lc code=end


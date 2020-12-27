/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	lastIndex := map[byte]int{}
	stack := []byte{}
	for i, ch := range []byte(s) {
		lastIndex[ch] = i
	}

out:
	for i, ch := range []byte(s) {
		for _, val := range stack {
			if val == ch {
				continue out
			}
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch && lastIndex[stack[len(stack)-1]] > i {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
	}
	return string(stack)
}

// @lc code=end


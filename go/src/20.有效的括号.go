/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	var parentheses = map[byte]byte{
		')': '(', '}': '{', ']': '[',
	}

	var stack []byte
	var length = len(s)
	for i := 0; i < length; i++ {
		if v, ok := parentheses[s[i]]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

// @lc code=end

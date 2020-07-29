/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(S string, T string) bool {
	var stack []byte
	for i := 0; i < len(S); i++ {
		if S[i] == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			}
		} else {
			stack = append(stack, S[i])
		}
	}

	var index int = len(stack) - 1
	var del int = 0
	for i := len(T) - 1; i >= 0 ; i-- {
		if T[i] == '#' {
			del++
		} else if del > 0 {
			del--
			continue
		} else if index >= 0 && T[i] == stack[index] {
			index--
		} else {
			return false
		}
	}
	return index == -1
}
// @lc code=end


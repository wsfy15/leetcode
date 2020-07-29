/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	sb := strings.Builder{}
	n := len(S)
	outerExist := false
	leftBrack := 0
	for i := 0; i < n; i++ {
		if !outerExist {
			outerExist = true
			continue
		}
		if S[i] == '(' {
			leftBrack++
			sb.WriteByte('(')
		} else if leftBrack > 0 {
			leftBrack--
			sb.WriteByte(')')
		} else {
			outerExist = false
		}
	}

	return sb.String()
}
// @lc code=end


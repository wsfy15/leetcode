/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(S string) string {
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && S[i] == stack[len(stack) - 1] {
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, S[i]) 
		}
	}

	return string(stack)
}
// @lc code=end

